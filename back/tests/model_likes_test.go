package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)

func TestSaveALike(t *testing.T) {
	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatalf("Error refreshing user, trip and like table %v\n", err)
	}
	user, trip, err := seedOneUserAndOneTrip()
	if err != nil {
		log.Fatalf("Cannot seed user and trip %v\n", err)
	}
	newLike := models.Like{
		ID:     1,
		UserID: user.ID,
		TripID: trip.ID,
	}
	savedLike, err := newLike.SaveLike(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the like: %v\n", err)
		return
	}
	assert.Equal(t, newLike.ID, savedLike.ID)
	assert.Equal(t, newLike.UserID, savedLike.UserID)
	assert.Equal(t, newLike.TripID, savedLike.TripID)
}

func TestGetLikeInfoForATrip(t *testing.T) {

	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatalf("Error refreshing user, trip and like table %v\n", err)
	}
	trip, users, likes, err := seedUsersTripsAndLikes()
	if err != nil {
		log.Fatalf("Error seeding user, trip and like table %v\n", err)
	}
	//Where likeInstance is an instance of the post initialize in setup_test.go
	_, err = likeInstance.GetLikesInfo(server.DB, trip.ID)
	if err != nil {
		t.Errorf("this is the error getting the likes: %v\n", err)
		return
	}
	assert.Equal(t, len(likes), 2)
	assert.Equal(t, len(users), 2) //two users like the trip
}

func TestDeleteALike(t *testing.T) {

	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatalf("Error refreshing user, trip and like table %v\n", err)
	}
	_, _, likes, err := seedUsersTripsAndLikes()
	if err != nil {
		log.Fatalf("Error seeding user, trip and like table %v\n", err)
	}
	// Delete the first like
	for _, v := range likes {
		if v.ID == 2 {
			continue
		}
		likeInstance.ID = v.ID //likeInstance is defined in setup_test.go
	}
	deletedLike, err := likeInstance.DeleteLike(server.DB)
	if err != nil {
		t.Errorf("this is the error deleting the like: %v\n", err)
		return
	}
	assert.Equal(t, deletedLike.ID, likeInstance.ID)
}

// When a trip is deleted, delete its likes
func TestDeleteLikesForATrip(t *testing.T) {

	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatalf("Error refreshing user, trip and like table %v\n", err)
	}
	trip, _, _, err := seedUsersTripsAndLikes()
	if err != nil {
		log.Fatalf("Error seeding user, trip and like table %v\n", err)
	}
	numberDeleted, err := likeInstance.DeleteTripLikes(server.DB, trip.ID)
	if err != nil {
		t.Errorf("this is the error deleting the like: %v\n", err)
		return
	}
	assert.Equal(t, numberDeleted, int64(2))
}

// When a user is deleted, delete its likes
func TestDeleteLikesForAUser(t *testing.T) {
	var userID uint64
	err := refreshUserTripAndLikeTable()
	if err != nil {
		log.Fatalf("Error refreshing user, trip and like table %v\n", err)
	}
	_, users, likes, err := seedUsersTripsAndLikes()
	if err != nil {
		log.Fatalf("Error seeding user, trip and like table %v\n", err)
	}
	for _, v := range likes {
		if v.ID == 2 {
			continue
		}
		likeInstance.ID = v.ID //likeInstance is defined in setup_test.go
	}
	// get the first user, this user has one like
	for _, v := range users {
		if v.ID == 2 {
			continue
		}
		userID = v.ID
	}
	numberDeleted, err := likeInstance.DeleteUserLikes(server.DB, userID)
	if err != nil {
		t.Errorf("this is the error deleting the like: %v\n", err)
		return
	}
	assert.Equal(t, numberDeleted, int64(1))
}
