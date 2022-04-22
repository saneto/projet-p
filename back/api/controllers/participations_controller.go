package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/mailer"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

// ParticipationTrip : function to participate to a trip
func (server *Server) ParticipationTrip(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

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
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// check if the user exist
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
	// check if the trip exist
	trip := models.Trip{}
	err = server.DB.Debug().Model(models.Trip{}).Where("id = ?", pid).Take(&trip).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	participate := models.ParticipateTrip{}
	participate.UserID = user.ID
	participate.TripID = trip.ID

	participateCreated, err := participate.SaveParticipation(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	// check if the user exist
	userCreatedTrip := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", trip.AuthorID).Take(&userCreatedTrip).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	fmt.Println("SEND EMAIL OCCURED")

	//Send mail to the user participating in the trip
	responseMail, err := mailer.SendMail.SendParticipationUser(user.Email, trip.Title, os.Getenv("SENDGRID_FROM"), os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	responseMail2, err2 := mailer.SendMail.SendParticipationUserCreatedTrip(userCreatedTrip.Email, trip.Title, user.Firstname, user.Email, os.Getenv("SENDGRID_FROM"), os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err2,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":        http.StatusCreated,
		"response":      participateCreated,
		"responseMail":  responseMail.RespBody,
		"responseMail2": responseMail2.RespBody,
	})
}

// GetParticipations : funtion to get the participations
func (server *Server) GetParticipations(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	tripID := c.Param("id")

	// Is a valid trip id given to us?
	pid, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		fmt.Println("this is the error: ", err)
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
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

	participate := models.ParticipateTrip{}

	participations, err := participate.GetParticipationInfo(server.DB, pid)
	if err != nil {
		errList["No_participation"] = "No Participation found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": participations,
	})
}

// GetParticipationByID : function to get the participation by user id
func (server *Server) GetParticipationByID(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	userID := c.Param("id")

	// Is a valid user id given to us?
	pid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		fmt.Println("this is the error: ", err)
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// check if the user exist
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", pid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	participate := models.ParticipateTrip{}

	participations, err := participate.GetParticipationsByUserID(server.DB, pid)
	if err != nil {
		errList["No_participations"] = "No Participations found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": participations,
	})
}

// WithdrawParticipation : function to withdraw a participation of a trip
/*func (server *Server) WithdrawParticipation(c *gin.Context) {

	participationID := c.Param("id")
	// Is a valid apply id given to us?
	lid, err := strconv.ParseUint(participationID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// Check if the participation exist
	participation := models.ParticipateTrip{}
	err = server.DB.Debug().Model(models.ParticipateTrip{}).Where("id = ?", lid).Take(&participation).Error
	if err != nil {
		errList["No_participation"] = "No Participation Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	// Is the authenticated user, the owner of this participation?
	if uid != participation.UserID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// If all the conditions are met, delete the participation
	_, err = participation.DeleteParticipation(server.DB)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Participation deleted",
	})
}*/

// DeleteParticipation : funtion to delete a participation
func (server *Server) DeleteParticipation(c *gin.Context) {

	tripID := c.Param("id")
	// Is a valid trip id given to us?
	pid, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		fmt.Println("this is the error: ", err)
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
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

	fmt.Println("delete a participation")

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	participation := models.ParticipateTrip{}

	// If all the conditions are met, delete the participation
	_, err = participation.DeleteParticipationTripAndUser(server.DB, pid, uid)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	participations, err := participation.GetParticipationInfo(server.DB, pid)
	if err != nil {
		errList["No_participation"] = "No Participation found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":                 http.StatusOK,
		"response":               "Participation deleted",
		"responseParticipations": participations,
	})
}
