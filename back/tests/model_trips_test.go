package tests

import (
	"log"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)

func TestFindAllTrips(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error refreshing user and trip table %v\n", err)
	}
	_, _, err = seedUsersAndTrips()
	if err != nil {
		log.Fatalf("Error seeding user and trip table %v\n", err)
	}
	//Where tripInstance is an instance of the trip initialize in setup_test.go
	trips, err := tripInstance.FindAllTrips(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the trips: %v\n", err)
		return
	}
	assert.Equal(t, len(*trips), 2)
}

func TestSaveTrip(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error user and trip refreshing table %v\n", err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	newTrip := models.Trip{
		Country:     "France",
		Title:       "Voyage au sud de la France",
		Description: "Voyage en petit comités pour s'amuser, apprendre à se connaître et découvrir des choses",
		StartDate:   start,
		EndDate:     end,
		NbDays:      5,
		MiddleAge:   20,
		NbTraveler:  4,
		Program:     "On partirai de Nice et nous longerons toute la côte Est jusqu'à Menton.",
		Lodging:     "Principalement du airbnb",
		Budget:      350,
		AuthorID:    user.ID,
	}

	savedTrip, err := newTrip.SaveTrip(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the trip: %v\n", err)
		return
	}
	assert.Equal(t, newTrip.ID, savedTrip.ID)
	assert.Equal(t, newTrip.Title, savedTrip.Title)
	assert.Equal(t, newTrip.Description, savedTrip.Description)
	assert.Equal(t, newTrip.AuthorID, savedTrip.AuthorID)
}

func TestFindTripByID(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error refreshing user and trip table: %v\n", err)
	}
	_, trip, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundTrip, err := trip.FindTripByID(server.DB, trip.ID)
	if err != nil {
		t.Errorf("this is the error getting one trip: %v\n", err)
		return
	}
	assert.Equal(t, foundTrip.ID, trip.ID)
	assert.Equal(t, foundTrip.Title, trip.Title)
	assert.Equal(t, foundTrip.Description, trip.Description)
}

func TestUpdateATrip(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error refreshing user and trip table: %v\n", err)
	}
	_, trip, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}

	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	tripUpdate := models.Trip{
		ID:          1,
		Country:     "France",
		Title:       "Voyage update",
		Description: "Voyage update en petit comités pour s'amuser.",
		StartDate:   start,
		EndDate:     end,
		NbDays:      5,
		MiddleAge:   20,
		NbTraveler:  4,
		Program:     "On partirai de Nice et nous longerons toute la côte Est jusqu'à Menton.",
		Lodging:     "Principalement du airbnb",
		Budget:      350,
		AuthorID:    trip.AuthorID,
	}

	updatedTrip, err := tripUpdate.UpdateATrip(server.DB, trip.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedTrip.ID, tripUpdate.ID)
	assert.Equal(t, updatedTrip.Title, tripUpdate.Title)
	assert.Equal(t, updatedTrip.Description, tripUpdate.Description)
	assert.Equal(t, updatedTrip.AuthorID, tripUpdate.AuthorID)
}

func TestDeleteATrip(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error refreshing user and trip table: %v\n", err)
	}
	_, trip, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := trip.DeleteATrip(server.DB, trip.ID, trip.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}

func TestDeleteUserTrips(t *testing.T) {

	err := refreshUserAndTripTable()
	if err != nil {
		log.Fatalf("Error refreshing user and trip table: %v\n", err)
	}
	user, _, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}

	numberDeleted, err := tripInstance.DeleteUserTrips(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error deleting the trip: %v\n", err)
		return
	}
	assert.Equal(t, numberDeleted, int64(1))
}
