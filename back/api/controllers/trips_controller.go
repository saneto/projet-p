package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/mailer"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

//CreateTrip : function to create a trip
func (server *Server) CreateTrip(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	trip := models.Trip{}

	err = json.Unmarshal(body, &trip)
	if err != nil {
		/*	errList["Unmarshal_error"] = "Cannot unmarshal body"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return*/
	}
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// check if the user exist:
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	trip.AuthorID = uid //the authenticated user is the one creating the trip

	trip.Prepare()
	errorMessages := trip.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	tripCreated, err := trip.SaveTrip(server.DB)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	d := models.ELKNotification{
		Message : "Votre voyage a été crée",
		Type: "Trip",
		TargetName: user.Firstname+" "+user.Lastname,
		CreatedAt : time.Now(),
		TargetID: 	user.ID,
		TargetEmail: user.Email,
	}
	
	server.sendNotification(d)
	server.insertTrips(c, &trip)
	server.insertELKTrip( &trip)

	fmt.Println("SEND EMAIL OCCURED")
	//Send the mail to the user
	responseMail, err := mailer.SendMail.SendCreateTrip(user.Email, os.Getenv("SENDGRID_FROM"), os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":       http.StatusCreated,
		"response":     tripCreated,
		"responseMail": responseMail.RespBody,
	})
}

//GetTrips : function to get all trips
func (server *Server) GetTrips(c *gin.Context) {

	trip := models.Trip{}

	trips, err := trip.FindAllTrips(server.DB)
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": trips,
	})
}

//GetTrip : function to get a trip
func (server *Server) GetTrip(c *gin.Context) {

	tripID := c.Param("id")
	pid, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	trip := models.Trip{}

	tripReceived, err := trip.FindTripByID(server.DB, pid)
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": tripReceived,
	})
}

//UpdateTrip : function to update a trip
func (server *Server) UpdateTrip(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	tripID := c.Param("id")
	// Check if the trip id is valid
	pid, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	//Check if the auth token is valid and get the user id from it
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	//Check if the trip exist
	origTrip := models.Trip{}
	err = server.DB.Debug().Model(models.Trip{}).Where("id = ?", pid).Take(&origTrip).Error
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	if uid != origTrip.AuthorID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// Read the data posted
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	// Start processing the request data
	trip := models.Trip{}
	err = json.Unmarshal(body, &trip)
	if err != nil {
		/*	errList["Unmarshal_error"] = "Cannot unmarshal body"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return*/
	}
	trip.ID = origTrip.ID //this is important to tell the model, the trip id to update.
	trip.AuthorID = origTrip.AuthorID

	trip.Prepare()
	errorMessages := trip.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	tripUpdated, err := trip.UpdateATrip(server.DB, pid)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	
	server.insertELKTrip(tripUpdated)
	
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": tripUpdated,
	})
}

//DeleteTrip : function to delete a trip
func (server *Server) DeleteTrip(c *gin.Context) {

	tripID := c.Param("id")
	// Is a valid trip id given to us?
	pid, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	fmt.Println("delete a trip")

	// Is this user authenticated ?
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// Check if the trip exist
	trip := models.Trip{}
	err = server.DB.Debug().Model(models.Trip{}).Where("id = ?", pid).Take(&trip).Error
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	// Is the authenticated user, the owner of this trip ?
	if uid != trip.AuthorID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If all the conditions are met, delete the trip
	_, err = trip.DeleteATrip(server.DB, pid, uid)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	comment := models.Comment{}
	like := models.Like{}
	participation := models.ParticipateTrip{}

	// Also delete the likes, participations and the comments that this trip have.
	_, err = comment.DeleteTripComments(server.DB, pid)

	_, err = participation.DeleteTripParticipations(server.DB, pid)

	_, err = like.DeleteTripLikes(server.DB, pid)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Trip deleted",
	})
}

//GetUserTrips : function to get the trips of a user
func (server *Server) GetUserTrips(c *gin.Context) {

	userID := c.Param("id")
	// Is a valid user id given to us?
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	trip := models.Trip{}
	trips, err := trip.FindUserTrips(server.DB, uint64(uid))
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": trips,
	})
}

//CreateProposal : function to send proposal to user
func (server *Server) CreateProposal(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	requestBody := map[string]string{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	fmt.Println("BODY")
	fmt.Println(requestBody["firstname"])

	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// check if the user exist:
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	

	fmt.Println("SEND EMAIL OCCURED")
	//Send the mail to the user
	responseMail, err := mailer.SendMail.SendTravelProposal(user.Firstname, user.Email, requestBody["firstname"], requestBody["email"], requestBody["message"], os.Getenv("SENDGRID_FROM"), os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"responseMail": responseMail.RespBody,
	})
}
