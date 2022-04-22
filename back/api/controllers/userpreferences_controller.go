package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

//CreateUserPreferences : function to create a user preferences
func (server *Server) CreateUserPreferences(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	userPref := models.UserPreferences{}

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

	userPref.UserID = uid //the authenticated user is the one creating the user preference

	userPref.Prepare()

	userPrefCreated, err := userPref.SaveUserPreferences(server.DB)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userPrefCreated,
	})
}

//GetAllUserPreferences : function to get all user preferences
func (server *Server) GetAllUserPreferences(c *gin.Context) {

	userPref := models.UserPreferences{}

	userprefs, err := userPref.FindAllUserPreferences(server.DB)
	if err != nil {
		errList["No_user_preferences"] = "No User Preferences Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userprefs,
	})
}

//GetUserPreference : function to get a user preference
func (server *Server) GetUserPreference(c *gin.Context) {

	userPreferenceID := c.Param("id")
	pid, err := strconv.ParseUint(userPreferenceID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	userPref := models.UserPreferences{}

	userPrefReceived, err := userPref.FindUserPreferenceByID(server.DB, pid)
	if err != nil {
		errList["No_user_preferences"] = "No User Preferences Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userPrefReceived,
	})
}

//UpdateUserPreference : function to update a user preference
func (server *Server) UpdateUserPreference(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	userPreferenceID := c.Param("id")
	// Check if the user preference id is valid
	pid, err := strconv.ParseUint(userPreferenceID, 10, 64)
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
	//Check if the user preference exist
	origUserPreference := models.UserPreferences{}
	err = server.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", pid).Take(&origUserPreference).Error
	if err != nil {
		errList["No_user_preferences"] = "No User Preferences Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	if uid != origUserPreference.UserID {
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
	userPref := models.UserPreferences{}
	err = json.Unmarshal(body, &userPref)
	if err != nil {
		/*	errList["Unmarshal_error"] = "Cannot unmarshal body"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return*/
	}
	userPref.ID = origUserPreference.ID //this is important to tell the model, the user preference id to update.
	userPref.UserID = origUserPreference.UserID

	userPref.Prepare()

	userPrefUpdated, err := userPref.UpdateAUserPreferences(server.DB, pid)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userPrefUpdated,
	})
}

//DeleteUserPreference : function to delete a user preference
func (server *Server) DeleteUserPreference(c *gin.Context) {

	userPrefID := c.Param("id")
	// Is a valid user preference id given to us?
	pid, err := strconv.ParseUint(userPrefID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	fmt.Println("delete a user preference")

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
	// Check if the user preference exist
	userPref := models.UserPreferences{}
	err = server.DB.Debug().Model(models.UserPreferences{}).Where("id = ?", pid).Take(&userPref).Error
	if err != nil {
		errList["No_user_preferences"] = "No User Preferences Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	// Is the authenticated user, the owner of this user preference ?
	if uid != userPref.UserID {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If all the conditions are met, delete the user preference
	_, err = userPref.DeleteAUserPreferences(server.DB, pid, uid)
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
		"response": "User Preference deleted",
	})
}

//GetUserPreferenceByUser : function to get the user preference of a user
func (server *Server) GetUserPreferenceByUser(c *gin.Context) {

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

	userPref := models.UserPreferences{}
	userPreferences, err := userPref.FindUserPreferenceByUserID(server.DB, uint64(uid))
	if err != nil {
		errList["No_user_preferences"] = "No User Preferences Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userPreferences,
	})
}
