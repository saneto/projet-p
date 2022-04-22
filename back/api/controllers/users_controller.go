package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/mailer"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/security"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/fileformat"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
)

// CreateUser : function Post for User
func (server *Server) CreateUser(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}
	userData := make(map[string]interface{})

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
	errorMessages := user.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	token, errtok := auth.CreateToken(userCreated.ID)
	if errtok != nil {
		fmt.Println("this is the error creating the token: ", errtok)
		return
	}

	userData["token"] = token
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["avatar_path"] = user.AvatarPath
	userData["firstname"] = user.Firstname
	userData["lastname"] = user.Lastname
	userData["dateofbirth"] = user.Dateofbirth
	userData["sexe"] = user.Sexe
	userData["city"] = user.City
	userData["phone_number"] = user.PhoneNumber
	userData["departure_airport"] = user.DepartureAirport
	userData["description"] = user.Description
	userData["is_guide"] = user.IsGuide
	userData["country_guide"] = user.CountryGuide

	userPreferences := models.UserPreferences{}
	userPreferences.UserID = userCreated.ID //create the user preference with the ID of the new user created
	userPreferences.Prepare()

	userPreferencesCreated, err := userPreferences.SaveUserPreferences(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	fmt.Println(userCreated.ID)

	fmt.Println("SEND MAIL OCCURRED HERE")

	//Send the mail to the user
	response, err := mailer.SendMail.SendWelcome(userCreated.Email, os.Getenv("SENDGRID_FROM"), os.Getenv("SENDGRID_API_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":                  http.StatusCreated,
		"response":                userData,
		"responseMail":            response.RespBody,
		"responseUserPreferences": userPreferencesCreated,
	})
}

// GetUsers : function to get all users
func (server *Server) GetUsers(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": users,
	})
}

// GetUsersGuide : function to get all users guide of a country
func (server *Server) GetUsersGuide(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	country := c.Param("country")

	fmt.Println(country)

	user := models.User{}

	users, err := user.FindUserGuideByCountry(server.DB, country)
	if err != nil {
		errList["No_user_guide"] = "No User Guide Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": users,
	})
}

// GetUser : function to get a specific user
func (server *Server) GetUser(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	userID := c.Param("id")

	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	user := models.User{}

	userGotten, err := user.FindUserByID(server.DB, uint64(uid))
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userGotten,
	})
}

//UpdateAvatar : connect to DigitalOcean to change avatar of an user
func (server *Server) UpdateAvatar(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	userID := c.Param("id")
	// Check if the user id is valid
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// Get user id from the token for valid tokens
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If the id is not the authenticated user id
	if tokenID != 0 && tokenID != uint64(uid) {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		errList["Invalid_file"] = "Fichier invalide"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		errList["Invalid_file"] = "Fichier invalide"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	defer f.Close()

	size := file.Size
	//The image should not be more than 500KB
	if size > int64(512000) {
		errList["Too_large"] = "Désolé, veuillez insérer une image de 500KB ou moins"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	buffer := make([]byte, size)
	f.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	//if the image is valid
	if !strings.HasPrefix(fileType, "image") {
		errList["Not_Image"] = "Veuillez insérer une image valide"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	filePath := fileformat.UniqueFormat(file.Filename)
	path := "/profile-photos/" + filePath
	params := &s3.PutObjectInput{
		Bucket:        aws.String("chodapi"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	}
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("DO_SPACES_KEY"), os.Getenv("DO_SPACES_SECRET"), os.Getenv("DO_SPACES_TOKEN")),
		Endpoint: aws.String(os.Getenv("DO_SPACES_ENDPOINT")),
		Region:   aws.String(os.Getenv("DO_SPACES_REGION")),
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	_, err = s3Client.PutObject(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := models.User{}
	user.AvatarPath = filePath
	user.Prepare()
	updatedUser, err := user.UpdateAUserAvatar(server.DB, uint64(uid))
	if err != nil {
		errList["Cannot_Save"] = "Impossible de sauvegarder l'image, Veuillez réessayer plus tard"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": updatedUser,
	})
}

// UpdateUser : function to update a user
func (server *Server) UpdateUser(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	userID := c.Param("id")
	// Check if the user id is valid
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// Get user id from the token for valid tokens
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If the id is not the authenticated user id
	if tokenID != 0 && tokenID != uint64(uid) {
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
	// Check for previous details
	formerUser := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&formerUser).Error
	if err != nil {
		errList["User_invalid"] = "The user is does not exist"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	newUser := models.User{}

	//When current password has content.
	if requestBody["current_password"] == "" && requestBody["new_password"] != "" {
		errList["Empty_current"] = "Veuillez fournir le mot de passe actuel"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] == "" {
		errList["Empty_new"] = "Veuillez fournir le nouveau mot de passe"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] != "" {
		//Also check if the new password
		if len(requestBody["new_password"]) < 6 {
			errList["Invalid_password"] = "Le mot de passe doit comporter au moins 6 caractères"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//if they do, check that the former password is correct
		err = security.VerifyPassword(formerUser.Password, requestBody["current_password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			errList["Password_mismatch"] = "Le mot de passe actuel n'est pas correct"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//update all the fields entered
		newUser.Firstname = requestBody["firstname"]
		newUser.Lastname = requestBody["lastname"]
		newUser.Sexe = requestBody["sexe"]
		newUser.Email = requestBody["email"]
		newUser.Password = requestBody["new_password"]
		newUser.City = requestBody["city"]
		newUser.IsGuide = requestBody["is_guide"]
		newUser.CountryGuide = requestBody["country_guide"]
		newUser.PhoneNumber = requestBody["phone_number"]
		newUser.DepartureAirport = requestBody["departure_airport"]
		newUser.Description = requestBody["description"]
		newUser.Dateofbirth = requestBody["dateofbirth"]
		newUser.AvatarPath = requestBody["avatar_path"]
	}

	//The password fields not entered, so update only the other fields
	newUser.Firstname = requestBody["firstname"]
	newUser.Lastname = requestBody["lastname"]
	newUser.Sexe = requestBody["sexe"]
	newUser.Email = requestBody["email"]
	newUser.City = requestBody["city"]
	newUser.IsGuide = requestBody["is_guide"]
	newUser.CountryGuide = requestBody["country_guide"]
	newUser.PhoneNumber = requestBody["phone_number"]
	newUser.DepartureAirport = requestBody["departure_airport"]
	newUser.Description = requestBody["description"]
	newUser.Dateofbirth = requestBody["dateofbirth"]
	newUser.AvatarPath = requestBody["avatar_path"]
	fmt.Println(requestBody["date_of_birth"])
	newUser.Prepare()
	errorMessages := newUser.Validate("update")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	updatedUser, err := newUser.UpdateAUser(server.DB, uint64(uid))
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
		"response": updatedUser,
	})
}

// DeleteUser : function to delete a specific user
func (server *Server) DeleteUser(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}
	var tokenID uint64
	userID := c.Param("id")
	// Check if the user id is valid
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// Get user id from the token for valid tokens
	tokenID, err = auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	// If the id is not the authenticated user id
	if tokenID != 0 && tokenID != uint64(uid) {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// Also delete the trips, likes, participation and the comments that this user created if any:
	comment := models.Comment{}
	like := models.Like{}
	trip := models.Trip{}
	participation := models.ParticipateTrip{}
	userPreferences := models.UserPreferences{}

	_, err = comment.DeleteUserComments(server.DB, uint64(uid))

	_, err = like.DeleteUserLikes(server.DB, uint64(uid))

	_, err = participation.DeleteUserParticipations(server.DB, uint64(uid))

	_, err = userPreferences.DeleteUserPreferencesUserDeleted(server.DB, uint64(uid))

	_, err = trip.DeleteUserTrips(server.DB, uint64(uid))

	user := models.User{}
	_, err = user.DeleteAUser(server.DB, uint64(uid))
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
		"response": "User deleted",
	})
}
