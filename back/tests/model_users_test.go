package tests

import (
	"log"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/stretchr/testify/assert"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	ti, _ := time.Parse(layout1, input1)

	newUser := models.User{
		ID:          1,
		Email:       "test@gmail.com",
		Firstname:   "test",
		Lastname:    "Test 2",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: ti,
		Sexe:        "Homme",
	}

	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Firstname, savedUser.Firstname)
}

func TestFindUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Firstname, user.Firstname)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	ti, _ := time.Parse(layout1, input1)

	userUpdate := models.User{
		ID:          1,
		Firstname:   "modi",
		Lastname:    "Update",
		Email:       "modiupdate@gmail.com",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: ti,
		Sexe:        "Homme",
	}

	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Firstname, userUpdate.Firstname)
}

func TestDeleteAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}
	isDeleted, err := userInstance.DeleteAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}
