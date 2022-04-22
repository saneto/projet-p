package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCommentTrip(t *testing.T) {

	var firstUserEmail, secondUserEmail string
	var firstTripID uint64

	err := refreshUserTripAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	users, trips, err := seedUsersAndTrips()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	for _, user := range users {
		if user.ID == 1 {
			firstUserEmail = user.Email
		}
		if user.ID == 2 {
			secondUserEmail = user.Email
		}
	}
	// Get only the first trip, which belongs to first user
	for _, trip := range trips {
		if trip.ID == 2 {
			continue
		}
		firstTripID = trip.ID
	}
	// Login both users
	password := "password"

	// Login First User
	tokenInterface1, err := server.SignIn(firstUserEmail, password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token1 := tokenInterface1["token"] //get only the token
	firstUserToken := fmt.Sprintf("Bearer %v", token1)

	// Login Second User
	tokenInterface2, err := server.SignIn(secondUserEmail, password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token2 := tokenInterface2["token"] //get only the token
	secondUserToken := fmt.Sprintf("Bearer %v", token2)
	fmt.Println("this is the second user token: ", secondUserToken)

	samples := []struct {
		tripIDString string
		inputJSON    string
		statusCode   int
		userID       uint64
		tripID       uint64
		Body         string
		tokenGiven   string
	}{
		{
			// User 1 can comment on his trip
			tripIDString: strconv.Itoa(int(firstTripID)), //we need the id as a string
			inputJSON:    `{"body": "comment from user 1"}`,
			statusCode:   201,
			userID:       1,
			tripID:       firstTripID,
			Body:         "comment from user 1",
			tokenGiven:   firstUserToken,
		},
		{
			// User 2 can also comment on user 1 trip
			tripIDString: strconv.Itoa(int(firstTripID)),
			inputJSON:    `{"body":"comment from user 2"}`,
			statusCode:   201,
			userID:       2,
			tripID:       firstTripID,
			Body:         "comment from user 2",
			tokenGiven:   secondUserToken,
		},
		{
			// When no body is provided:
			tripIDString: strconv.Itoa(int(firstTripID)),
			inputJSON:    `{"body":""}`,
			statusCode:   422,
			tripID:       firstTripID,
			tokenGiven:   secondUserToken,
		},
		{
			// Not authenticated (No token provided)
			tripIDString: strconv.Itoa(int(firstTripID)),
			statusCode:   401,
			tokenGiven:   "",
		},
		{
			// Wrong Token
			tripIDString: strconv.Itoa(int(firstTripID)),
			statusCode:   401,
			tokenGiven:   "This is an incorrect token",
		},
		{
			// When invalid trip id is given
			tripIDString: "unknwon",
			statusCode:   400,
		},
	}

	for _, v := range samples {

		gin.SetMode(gin.TestMode)

		r := gin.Default()

		r.POST("/comments/:id", server.CreateComment)
		req, err := http.NewRequest(http.MethodPost, "/comments/"+v.tripIDString, bytes.NewBufferString(v.inputJSON))
		req.Header.Set("Authorization", v.tokenGiven)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 201 {
			responseMap := responseInterface["response"].(map[string]interface{})
			assert.Equal(t, responseMap["trip_id"], float64(v.tripID))
			assert.Equal(t, responseMap["user_id"], float64(v.userID))
			assert.Equal(t, responseMap["body"], v.Body)
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["Required_body"] != nil {
				assert.Equal(t, responseMap["Required_body"], "Le commentaire ne peut pas être vide")
			}
		}
	}
}

func TestGetComments(t *testing.T) {
	gin.SetMode(gin.TestMode)
	err := refreshUserTripAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	trip, users, comments, err := seedUsersTripsAndComments()
	if err != nil {
		log.Fatalf("Cannot seed tables %v\n", err)
	}
	commentsSample := []struct {
		tripID         string
		usersLength    int
		commentsLength int
		statusCode     int
	}{
		{
			tripID:         strconv.Itoa(int(trip.ID)),
			statusCode:     200,
			usersLength:    len(users),
			commentsLength: len(comments),
		},
		{
			tripID:     "unknwon",
			statusCode: 400,
		},
		{
			tripID:     strconv.Itoa(12322), //an id that does not exist
			statusCode: 404,
		},
	}
	for _, v := range commentsSample {
		r := gin.Default()
		r.GET("/comments/:id", server.GetComments)
		req, err := http.NewRequest(http.MethodGet, "/comments/"+v.tripID, nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json here: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			responseMap := responseInterface["response"].([]interface{})
			assert.Equal(t, len(responseMap), v.commentsLength)
			assert.Equal(t, v.usersLength, 2)
		}
		if v.statusCode == 400 || v.statusCode == 404 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["No_trip"] != nil {
				assert.Equal(t, responseMap["No_trip"], "No trip found")
			}
		}
	}
}

func TestUpdateComment(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var secondUserEmail, secondUserPassword string
	var secondUserID uint64
	var secondCommentID uint64

	err := refreshUserTripAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	trip, users, comments, err := seedUsersTripsAndComments()
	if err != nil {
		log.Fatalf("Cannot seed tables %v\n", err)
	}
	// Get only the second user
	for _, user := range users {
		if user.ID == 1 {
			continue
		}
		secondUserID = user.ID
		secondUserEmail = user.Email
		secondUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	// Get only the second comment
	for _, comment := range comments {
		if comment.ID == 1 {
			continue
		}
		secondCommentID = comment.ID
	}
	//Login the user and get the authentication token
	tokenInterface, err := server.SignIn(secondUserEmail, secondUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	commentsSample := []struct {
		commentID  string
		updateJSON string
		Body       string
		tokenGiven string
		statusCode int
	}{
		{
			commentID:  strconv.Itoa(int(secondCommentID)),
			updateJSON: `{"Body":"This is the update body"}`,
			statusCode: 200,
			Body:       "This is the update body",
			tokenGiven: tokenString,
		},
		{
			// When the body field is empty
			commentID:  strconv.Itoa(int(secondCommentID)),
			updateJSON: `{"Body":""}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
		{
			//an id that does not exist
			commentID:  strconv.Itoa(12322),
			statusCode: 404,
			tokenGiven: tokenString,
		},
		{
			//When the user is not authenticated
			commentID:  strconv.Itoa(int(secondCommentID)),
			statusCode: 401,
			tokenGiven: "",
		},
		{
			//When wrong token is passed
			commentID:  strconv.Itoa(int(secondCommentID)),
			statusCode: 401,
			tokenGiven: "this is a wrong token",
		},
		{
			// When id passed is invalid
			commentID:  "unknwon",
			statusCode: 400,
		},
	}

	for _, v := range commentsSample {
		r := gin.Default()
		r.PUT("/comments/:id", server.UpdateComment)
		req, err := http.NewRequest(http.MethodPut, "/comments/"+v.commentID, bytes.NewBufferString(v.updateJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		req.Header.Set("Authorization", v.tokenGiven)
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json here: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			responseMap := responseInterface["response"].(map[string]interface{})
			assert.Equal(t, responseMap["trip_id"], float64(trip.ID))
			assert.Equal(t, responseMap["user_id"], float64(secondUserID))
			assert.Equal(t, responseMap["body"], v.Body)
		}
		if v.statusCode == 400 || v.statusCode == 401 || v.statusCode == 404 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["No_comment"] != nil {
				assert.Equal(t, responseMap["No_comment"], "No Comment Found")
			}
		}
	}
}

func TestDeleteComment(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var secondUserEmail, secondUserPassword string
	var secondCommentID uint64

	err := refreshUserTripAndCommentTable()
	if err != nil {
		log.Fatal(err)
	}
	_, users, comments, err := seedUsersTripsAndComments()
	if err != nil {
		log.Fatalf("Cannot seed tables %v\n", err)
	}
	// Get only the second user
	for _, user := range users {
		if user.ID == 1 {
			continue
		}
		secondUserEmail = user.Email
		secondUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	// Get only the second comment
	for _, comment := range comments {
		if comment.ID == 1 {
			continue
		}
		secondCommentID = comment.ID
	}

	//Login the user and get the authentication token
	tokenInterface, err := server.SignIn(secondUserEmail, secondUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	commentsSample := []struct {
		commentID      string
		usersLength    int
		tokenGiven     string
		commentsLength int
		statusCode     int
	}{
		{
			commentID:  strconv.Itoa(int(secondCommentID)),
			statusCode: 200,
			tokenGiven: tokenString,
		},
		{
			//an id that does not exist
			commentID:  strconv.Itoa(12322),
			statusCode: 404,
			tokenGiven: tokenString,
		},
		{
			//When the user is not authenticated
			commentID:  strconv.Itoa(int(secondCommentID)),
			statusCode: 401,
			tokenGiven: "",
		},
		{
			//When wrong token is passed
			commentID:  strconv.Itoa(int(secondCommentID)),
			statusCode: 401,
			tokenGiven: "this is a wrong token",
		},
		{
			// When id passed is invalid
			commentID:  "unknwon",
			statusCode: 400,
		},
	}

	for _, v := range commentsSample {

		r := gin.Default()
		r.DELETE("/comments/:id", server.DeleteComment)
		req, err := http.NewRequest(http.MethodDelete, "/comments/"+v.commentID, nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		req.Header.Set("Authorization", v.tokenGiven)
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json here: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			responseMap := responseInterface["response"]
			assert.Equal(t, responseMap, "Comment deleted")
		}
		if v.statusCode == 400 || v.statusCode == 401 || v.statusCode == 404 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["No_comment"] != nil {
				assert.Equal(t, responseMap["No_comment"], "No Comment Found")
			}
		}
	}
}
