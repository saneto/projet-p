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

func TestCreateTrip(t *testing.T) {

	gin.SetMode(gin.TestMode)

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}
	// Note: the value of the user password before it was hashed is "password". so:
	password := "password"
	tokenInterface, err := server.SignIn(user.Email, password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	// Note that the author id is obtained from the token, so we dont pass it
	samples := []struct {
		inputJSON   string
		statusCode  int
		title       string
		description string
		country     string
		tokenGiven  string
	}{
		{
			inputJSON:   `{"title":"The title", "description": "the content", "country": "the country"}`,
			statusCode:  201,
			tokenGiven:  tokenString,
			title:       "The title",
			description: "the content",
			country:     "the country",
		},
		{
			// When no token is passed
			inputJSON:  `{"title":"When no token is passed", "description": "the content", "country": "the country"}`,
			statusCode: 401,
			tokenGiven: "",
		},
		{
			// When incorrect token is passed
			inputJSON:  `{"title":"When incorrect token is passed", "description": "the content", "country": "the country"}`,
			statusCode: 401,
			tokenGiven: "This is an incorrect token",
		},
		{
			inputJSON:  `{"title": "", "description": "the content", "country": "the country"}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
		{
			inputJSON:  `{"title": "the title", "description": "", "country": "the country"}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
	}

	for _, v := range samples {

		r := gin.Default()

		r.POST("/trips", server.CreateTrip)
		req, err := http.NewRequest(http.MethodPost, "/trips", bytes.NewBufferString(v.inputJSON))
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
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["description"], v.description)
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 {
			responseMap := responseInterface["error"].(map[string]interface{})

			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["Required_title"] != nil {
				assert.Equal(t, responseMap["Required_title"], "Required Title")
			}
			if responseMap["Required_description"] != nil {
				assert.Equal(t, responseMap["Required_description"], "Required Description")
			}
		}
	}
}

func TestGetTrips(t *testing.T) {

	gin.SetMode(gin.TestMode)

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = seedUsersAndTrips()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/trips", server.GetTrips)

	req, err := http.NewRequest(http.MethodGet, "/trips", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	tripsInterface := make(map[string]interface{})

	err = json.Unmarshal([]byte(rr.Body.String()), &tripsInterface)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	// This is so that we can get the length of the trips:
	theTrips := tripsInterface["response"].([]interface{})
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(theTrips), 2)
}

func TestGetTripByID(t *testing.T) {

	gin.SetMode(gin.TestMode)

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatal(err)
	}
	_, trip, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatal(err)
	}

	tripSample := []struct {
		id          string
		statusCode  int
		title       string
		description string
		country     string
		author_id   uint64
	}{
		{
			id:          strconv.Itoa(int(trip.ID)),
			statusCode:  200,
			title:       trip.Title,
			description: trip.Description,
			country:     trip.Country,
			author_id:   trip.AuthorID,
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
		{
			id:         strconv.Itoa(12322), //an id that does not exist
			statusCode: 404,
		},
	}

	for _, v := range tripSample {
		req, _ := http.NewRequest("GET", "/trips/"+v.id, nil)
		rr := httptest.NewRecorder()

		r := gin.Default()
		r.GET("/trips/:id", server.GetTrip)
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			responseMap := responseInterface["response"].(map[string]interface{})
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["description"], v.description)
			assert.Equal(t, responseMap["country"], v.country)
			assert.Equal(t, responseMap["author_id"], float64(v.author_id))
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

func TestUpdateTrip(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var TripUserEmail, TripUserPassword string
	var AuthTripID uint64

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatal(err)
	}
	users, trips, err := seedUsersAndTrips()
	if err != nil {
		log.Fatal(err)
	}
	// Get only the first user
	for _, user := range users {
		if user.ID == 2 {
			continue
		}
		TripUserEmail = user.Email
		TripUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	// Get only the first trip
	for _, trip := range trips {
		if trip.ID == 2 {
			continue
		}
		AuthTripID = trip.ID
	}
	//Login the user and get the authentication token
	tokenInterface, err := server.SignIn(TripUserEmail, TripUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	samples := []struct {
		id          string
		updateJSON  string
		statusCode  int
		title       string
		description string
		country     string
		tokenGiven  string
	}{
		{
			// Convert int64 to int first before converting to string
			id:          strconv.Itoa(int(AuthTripID)),
			updateJSON:  `{"title":"The updated trip", "description": "This is the updated description", "country": "this is the updated country"}`,
			statusCode:  200,
			title:       "The updated trip",
			description: "This is the updated description",
			country:     "This is the updated country",
			tokenGiven:  tokenString,
		},
		{
			// When no token is provided
			id:         strconv.Itoa(int(AuthTripID)),
			updateJSON: `{"title":"This is still another title", "description": "This is the updated description", "country": "this is the updated country"}`,
			tokenGiven: "",
			statusCode: 401,
		},
		{
			// When incorrect token is provided
			id:         strconv.Itoa(int(AuthTripID)),
			updateJSON: `{"title":"This is still another title", "description": "This is the updated description", "country": "this is the updated country"}`,
			tokenGiven: "this is an incorrect token",
			statusCode: 401,
		},
		{
			// When title is not given
			id:         strconv.Itoa(int(AuthTripID)),
			updateJSON: `{"title":"", "description": "This is the updated description", "country": "this is the updated country"}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
		{
			// When description is not given
			id:         strconv.Itoa(int(AuthTripID)),
			updateJSON: `{"title":"Awesome title", "description": "", "country": "this is the updated country"}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
		{
			// When country is not given
			id:         strconv.Itoa(int(AuthTripID)),
			updateJSON: `{"title":"Awesome title", "description": "Awesome description", "country": ""}`,
			statusCode: 422,
			tokenGiven: tokenString,
		},
		{
			// When invalid post id is given
			id:         "unknwon",
			statusCode: 400,
		},
	}

	for _, v := range samples {

		r := gin.Default()

		r.PUT("/trips/:id", server.UpdateTrip)
		req, err := http.NewRequest(http.MethodPut, "/trips/"+v.id, bytes.NewBufferString(v.updateJSON))
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

		if v.statusCode == 200 {
			//casting the interface to map:
			responseMap := responseInterface["response"].(map[string]interface{})
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["description"], v.description)
			assert.Equal(t, responseMap["country"], v.country)
		}
		if v.statusCode == 400 || v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 {
			responseMap := responseInterface["error"].(map[string]interface{})
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Required_title"] != nil {
				assert.Equal(t, responseMap["Required_title"], "Required Title")
			}
			if responseMap["Required_description"] != nil {
				assert.Equal(t, responseMap["Required_description"], "Required Description")
			}
			if responseMap["Required_country"] != nil {
				assert.Equal(t, responseMap["Required_country"], "Required Country")
			}
		}
	}
}

func TestDeleteTrip(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var TripUserEmail, TripUserPassword string
	var AuthTripID uint64

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatal(err)
	}
	users, trips, err := seedUsersAndTrips()
	if err != nil {
		log.Fatal(err)
	}
	// Get only the second user
	for _, user := range users {
		if user.ID == 1 {
			continue
		}
		TripUserEmail = user.Email
		TripUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	// Get only the second trip
	for _, trip := range trips {
		if trip.ID == 1 {
			continue
		}
		AuthTripID = trip.ID
	}
	//Login the user and get the authentication token
	tokenInterface, err := server.SignIn(TripUserEmail, TripUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	token := tokenInterface["token"] //get only the token
	tokenString := fmt.Sprintf("Bearer %v", token)

	postSample := []struct {
		id           string
		tokenGiven   string
		statusCode   int
		errorMessage string
	}{
		{
			// Convert int64 to int first before converting to string
			id:         strconv.Itoa(int(AuthTripID)),
			tokenGiven: tokenString,
			statusCode: 200,
		},
		{
			// When empty token is passed
			id:         strconv.Itoa(int(AuthTripID)),
			tokenGiven: "",
			statusCode: 401,
		},
		{
			// When incorrect token is passed
			id:         strconv.Itoa(int(AuthTripID)),
			tokenGiven: "This is an incorrect token",
			statusCode: 401,
		},
		{
			id:         "unknwon",
			tokenGiven: tokenString,
			statusCode: 400,
		},
		{
			id:           strconv.Itoa(int(1)),
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
	}

	for _, v := range postSample {
		r := gin.Default()
		r.DELETE("/trips/:id", server.DeleteTrip)
		req, _ := http.NewRequest(http.MethodDelete, "/trips/"+v.id, nil)
		req.Header.Set("Authorization", v.tokenGiven)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseInterface := make(map[string]interface{})

		err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
		if err != nil {
			t.Errorf("Cannot convert to json here: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, responseInterface["response"], "Trip deleted")
		}

		if v.statusCode == 400 || v.statusCode == 401 {
			responseMap := responseInterface["error"].(map[string]interface{})

			if responseMap["Invalid_request"] != nil {
				assert.Equal(t, responseMap["Invalid_request"], "Invalid Request")
			}
			if responseMap["Unauthorized"] != nil {
				assert.Equal(t, responseMap["Unauthorized"], "Unauthorized")
			}
		}
	}
}
