package tests

import (
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

func TestLikeTrip(t *testing.T) {

	var firstUserEmail, secondUserEmail string
	var firstTripID uint64

	err := refreshUserTripAndLikeTable()
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

	samples := []struct {
		tripIDString string
		statusCode   int
		userID       uint64
		tripID       uint64
		tokenGiven   string
	}{
		{
			// User 1 can like his trip
			tripIDString: strconv.Itoa(int(firstTripID)), //we need the id as a string
			statusCode:   201,
			userID:       1,
			tripID:       firstTripID,
			tokenGiven:   firstUserToken,
		},
		{
			// User 2 can also like user 1 post
			tripIDString: strconv.Itoa(int(firstTripID)),
			statusCode:   201,
			userID:       2,
			tripID:       firstTripID,
			tokenGiven:   secondUserToken,
		},
		{
			// An authenticated user cannot like a post more than once
			tripIDString: strconv.Itoa(int(firstTripID)),
			statusCode:   500,
			tokenGiven:   firstUserToken,
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
	}

	for _, v := range samples {

		gin.SetMode(gin.TestMode)

		r := gin.Default()

		r.POST("/likes/:id", server.LikeTrip)
		req, err := http.NewRequest(http.MethodPost, "/likes/"+v.tripIDString, nil)
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
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 {
			responseMap := responseInterface["error"].(map[string]interface{})

			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["Double_like"] != nil {
				assert.Equal(t, responseMap["Double_like"], "Vous aimez déjà ce voyage")
			}
		}
	}
}

func TestGetLikes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatal(err)
	}
	trip, users, likes, err := seedUsersTripsAndLikes()
	if err != nil {
		log.Fatalf("Cannot seed tables %v\n", err)
	}
	likesSample := []struct {
		tripID      string
		usersLength int
		likesLength int
		statusCode  int
	}{
		{
			tripID:      strconv.Itoa(int(trip.ID)),
			statusCode:  200,
			usersLength: len(users),
			likesLength: len(likes),
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
	for _, v := range likesSample {

		r := gin.Default()
		r.GET("/likes/:id", server.GetLikes)
		req, err := http.NewRequest(http.MethodGet, "/likes/"+v.tripID, nil)
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
			assert.Equal(t, len(responseMap), v.likesLength)
			assert.Equal(t, v.usersLength, 2)
		}
		if v.statusCode == 400 || v.statusCode == 404 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["No_trip"] != nil {
				assert.Equal(t, responseMap["No_trip"], "No Trip Found")
			}
		}
	}
}

func TestDeleteLike(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var secondUserEmail, secondUserPassword string
	var secondLike uint64

	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatal(err)
	}
	_, users, likes, err := seedUsersTripsAndLikes()
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
	// Get only the second like
	for _, like := range likes {
		if like.ID == 1 {
			continue
		}
		secondLike = like.ID
	}

	//Login the user and get the authentication token
	tokenInterface, err := server.SignIn(secondUserEmail, secondUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	likesSample := []struct {
		likeID      string
		usersLength int
		tokenGiven  string
		likesLength int
		statusCode  int
	}{
		{
			likeID:     strconv.Itoa(int(secondLike)),
			statusCode: 200,
			tokenGiven: tokenString,
		},
		{
			//an id that does not exist
			likeID:     strconv.Itoa(12322),
			statusCode: 404,
			tokenGiven: tokenString,
		},
		{
			//When the user is not authenticated
			likeID:     strconv.Itoa(int(secondLike)),
			statusCode: 401,
			tokenGiven: "",
		},
		{
			//When wrong token is passed
			likeID:     strconv.Itoa(int(secondLike)),
			statusCode: 401,
			tokenGiven: "this is a wrong token",
		},
		{
			// When id passed is invalid
			likeID:     "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range likesSample {

		r := gin.Default()
		r.GET("/likes/:id", server.UnLikeTrip)
		req, err := http.NewRequest(http.MethodGet, "/likes/"+v.likeID, nil)
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
			assert.Equal(t, responseMap, "Like deleted")
		}
		if v.statusCode == 400 || v.statusCode == 401 || v.statusCode == 404 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["No_like"] != nil {
				assert.Equal(t, responseMap["No_like"], "No Like Found")
			}
		}
	}
}
