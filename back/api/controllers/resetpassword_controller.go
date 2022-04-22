package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/mailer"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/security"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

// ForgotPassword : funtion forgot password
func (server *Server) ForgotPassword(c *gin.Context) {
	//remove any possible error, because the frontend dont reload
	errList = map[string]string{}

	// Start processing the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	user.Prepare()
	errorMessages := user.Validate("forgotpassword")
	if len(errorMessages) > 0 {
		//errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorMessages,
		})
		return
	}
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		errList["No_email"] = "Désolé, nous ne reconnaissons pas cet e-mail"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	resetPassword := models.ResetPassword{}
	resetPassword.Prepare()

	//generate the token
	token := security.TokenHash(user.Email)
	resetPassword.Email = user.Email
	resetPassword.Token = token

	resetDetails, err := resetPassword.SaveDetails(server.DB)
	if err != nil {
		errList = formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	fmt.Println("SEND MAIL OCCURRED HERE")
	//Send the mail to the user
	response, err := mailer.SendMail.SendResetPassword(resetDetails.Email, os.Getenv("SENDGRID_FROM"), resetDetails.Token, os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": response.RespBody,
	})
}

// ResetPassword : funtion to reset a password
func (server *Server) ResetPassword(c *gin.Context) {
	//remove any possible error, because the frontend dont reload
	errList = map[string]string{}

	// Start processing the request
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
	user := models.User{}
	resetPassword := models.ResetPassword{}

	err = server.DB.Debug().Model(models.ResetPassword{}).Where("token = ?", requestBody["token"]).Take(&resetPassword).Error
	if err != nil {
		errList["Invalid_token"] = "Lien invalide. Essayez à nouveau."
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["new_password"] == "" || requestBody["retype_password"] == "" {
		errList["Empty_passwords"] = "Veuillez vous assurer que les deux champs sont remplis"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["new_password"] != "" && requestBody["retype_password"] != "" {
		//Also check if the new password
		if len(requestBody["new_password"]) < 6 || len(requestBody["retype_password"]) < 6 {
			errList["Invalid_Passwords"] = "Le mot de passe doit comporter au moins 6 caractères"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		if requestBody["new_password"] != requestBody["retype_password"] {
			errList["Password_unequal"] = "Les mots de passe fournis ne correspondent pas"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//Note this password will be hashed before it is saved in the model
		user.Password = requestBody["new_password"]
		user.Email = resetPassword.Email

		//update the password
		user.Prepare()
		err := user.UpdatePassword(server.DB)
		if err != nil {
			fmt.Println("this is the error: ", err)
			errList["Cannot_save"] = "Cannot Save, Please try again later"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//Delete the token record so is not used again
		_, err = resetPassword.DeleteDetails(server.DB)
		if err != nil {
			errList["Cannot_delete"] = "Cannot Delete record, Please try again later"
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  errList,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": "Success",
		})
	}
}

// ResetUserPassword : funtion to reset user password
func (server *Server) ResetUserPassword(c *gin.Context) {

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

	// Start processing the request
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

	err = security.VerifyPassword(user.Password, requestBody["current_password"])
	if err != nil {
		errList["Empty_passwords"] = "Wrong current password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	if requestBody["new_password"] == "" || requestBody["retype_password"] == "" {
		errList["Empty_passwords"] = "Veuillez vous assurer que les deux champs sont remplis"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	if requestBody["new_password"] != "" && requestBody["retype_password"] != "" {
		//Also check if the new password
		if len(requestBody["new_password"]) < 6 || len(requestBody["retype_password"]) < 6 {
			errList["Invalid_Passwords"] = "Le mot de passe doit comporter au moins 6 caractères"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		if requestBody["new_password"] != requestBody["retype_password"] {
			errList["Password_unequal"] = "Les mots de passe fournis ne correspondent pas"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}

		if user.Email == requestBody["email"] {
			user.Password = requestBody["new_password"]
			err := user.UpdatePassword(server.DB)
			if err != nil {
				fmt.Println("this is the error: ", err)
				errList["Cannot_save"] = "Cannot Save, Please try again later"
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"status": http.StatusUnprocessableEntity,
					"error":  errList,
				})
				return
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Success",
	})

}
