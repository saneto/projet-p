package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"time"
	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

// LikeTrip : function to like a trip
func (server *Server) LikeTrip(c *gin.Context) {

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
	err = server.DB.Debug().Model(models.Trip{}).Where("id = ?", pid).Preload("Author").Take(&trip).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	like := models.Like{}
	like.UserID = user.ID
	like.TripID = trip.ID

	likeCreated, err := like.SaveLike(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	d := models.ELKNotification{
		Message : user.Firstname+" a aimé votre voyage",
		Type: "Like",
		TargetName: trip.Author.Firstname+" "+trip.Author.Lastname,
		CreatedAt : time.Now(),
		TargetID: 	trip.Author.ID,
		TargetEmail: trip.Author.Email,
	}
	
	server.sendNotification(d)
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": likeCreated,
	})
}

// GetLikes : funtion to get the likes
func (server *Server) GetLikes(c *gin.Context) {
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

	like := models.Like{}

	likes, err := like.GetLikesInfo(server.DB, pid)
	if err != nil {
		errList["No_likes"] = "No Likes found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": likes,
	})
}

// GetLikesByUserID : function to get the likes by user id
func (server *Server) GetLikesByUserID(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	userID := c.Param("id")

	// Is a valid user id given to us?
	uid, err := strconv.ParseUint(userID, 10, 64)
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
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	like := models.Like{}

	likes, err := like.GetLikesByUserID(server.DB, uid)
	if err != nil {
		errList["No_likes"] = "No Likes found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": likes,
	})
}

// UnLikeTrip : funtion to unlike a trip
/*func (server *Server) UnLikeTrip(c *gin.Context) {

	likeID := c.Param("id")
	// Is a valid like id given to us?
	lid, err := strconv.ParseUint(likeID, 10, 64)
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
	// Check if the like exist
	like := models.Like{}
	err = server.DB.Debug().Model(models.Like{}).Where("id = ?", lid).Take(&like).Error
	if err != nil {
		errList["No_like"] = "No Like Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	// Is the authenticated user, the owner of this like?
	if uid != like.UserID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// If all the conditions are met, delete the like
	_, err = like.DeleteLike(server.DB)
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
		"response": "Like deleted",
	})
}*/

// DeleteLike : funtion to unlike a trip
func (server *Server) DeleteLike(c *gin.Context) {

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

	fmt.Println("delete a like")

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

	like := models.Like{}

	// If all the conditions are met, delete the like
	_, err = like.DeleteLikeTripAndUser(server.DB, pid, uid)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Like deleted",
	})
}

//GetAllLikes : function to get all likes
func (server *Server) GetAllLikes(c *gin.Context) {

	like := models.Like{}

	likes, err := like.FindAllLikes(server.DB)
	if err != nil {
		errList["No_likes"] = "No Likes Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": likes,
	})
}
