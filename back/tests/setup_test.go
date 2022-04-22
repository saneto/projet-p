package tests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/controllers"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}
var tripInstance = models.Trip{}
var likeInstance = models.Like{}
var commentInstance = models.Comment{}

func TestMain(m *testing.M) {
	// UNCOMMENT THIS WHILE TESTING ON LOCAL (WITHOUT USING CIRCLE CI), BUT LEAVE IT COMMENTED IF YOU ARE USING CIRCLE CI
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	// UNTIL HERE
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	////////////////////////////////// UNCOMMENT THIS WHILE TESTING ON LOCAL(WITHOUT USING CIRCLE CI) ///////////////////////
	TestDbDriver := os.Getenv("TEST_DB_DRIVER")
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	/////////////////////////////////  END OF LOCAL TEST DATABASE SETUP ///////////////////////////////////////////////////

	////////////////////////////////// COMMENT THIS WHILE TESTING ON LOCAL (USING CIRCLE CI)  //////////////////////
	// WE HAVE TO INPUT TESTING DATA MANUALLY BECAUSE CIRCLECI, CANNOT READ THE ".env" FILE WHICH, WE WOULD HAVE ADDED THE TEST CONFIG THERE
	// SO MANUALLY ADD THE NAME OF THE DATABASE, THE USER AND THE PASSWORD, AS SEEN BELOW:
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "steven", "forum_db_test", "password")
	//server.DB, err = gorm.Open("postgres", DBURL)
	//if err != nil {
	//	fmt.Printf("Cannot connect to %s database\n", "postgres")
	//	log.Fatal("This is the error:", err)
	//} else {
	//	fmt.Printf("We are connected to the %s database\n", "postgres")
	//}
	//////////////////////////////// END OF USING CIRCLE CI ////////////////////////////////////////////////////////////////
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	user := models.User{
		Firstname:   "Pet",
		Lastname:    "Last",
		Email:       "pet@gmail.com",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: t,
		Sexe:        "Homme",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	users := []models.User{
		models.User{
			Firstname:   "Steven",
			Lastname:    "Victor",
			Email:       "steven@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
		models.User{
			Firstname:   "Kenny",
			Lastname:    "Morris",
			Email:       "kenny@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func refreshUserAndTripTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Trip{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Trip{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOneTrip() (models.User, models.Trip, error) {

	err := refreshUserAndTripTable()
	if err != nil {
		return models.User{}, models.Trip{}, err
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	user := models.User{
		Firstname:   "Sam",
		Lastname:    "Phil",
		Email:       "sam@gmail.com",
		Password:    "password",
		Accesslevel: 1,
		Dateofbirth: t,
		Sexe:        "Homme",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, models.Trip{}, err
	}

	input := "2020-01-06"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-11"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	trip := models.Trip{
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

	err = server.DB.Model(&models.Trip{}).Create(&trip).Error
	if err != nil {
		return models.User{}, models.Trip{}, err
	}
	return user, trip, nil
}

func seedUsersAndTrips() ([]models.User, []models.Trip, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Trip{}, err
	}

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	var users = []models.User{
		models.User{
			Firstname:   "Steven",
			Lastname:    "Victor",
			Email:       "steven@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
		models.User{
			Firstname:   "Kenny",
			Lastname:    "Morris",
			Email:       "kenny@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
	}

	input := "2020-09-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	var trips = []models.Trip{
		models.Trip{
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
		},
		models.Trip{
			Country:     "Espagne",
			Title:       "Road Trip détente dans le sud de l'Espagne",
			Description: "Flaner dans les plus belles plages de la côte sud d'Espagne, profiter du soleil et se reposer !",
			StartDate:   start,
			EndDate:     end,
			NbDays:      10,
			MiddleAge:   25,
			NbTraveler:  6,
			Program:     "Pas encore fait, mais si vous êtes intéressé, envoyer moi un MP et on y réfléchira ensemble.",
			Lodging:     "Principalement du airbnb ou auberge de jeunesse.",
			Budget:      500,
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		trips[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Trip{}).Create(&trips[i]).Error
		if err != nil {
			log.Fatalf("cannot seed trips table: %v", err)
		}
	}
	return users, trips, nil
}

func refreshUserTripAndLikeTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Trip{}, &models.Like{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Trip{}, &models.Like{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user, trip and like tables")
	return nil
}

func seedUsersTripsAndLikes() (models.Trip, []models.User, []models.Like, error) {
	// The idea here is: two users can like one trip
	var err error

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	var users = []models.User{
		models.User{
			Firstname:   "Dwayne",
			Lastname:    "Heu",
			Email:       "dwayne@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
		models.User{
			Firstname:   "Jen",
			Lastname:    "Stella",
			Email:       "jen@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Femme",
		},
	}

	input := "2020-09-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	trip := models.Trip{
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
	}

	err = server.DB.Model(&models.Trip{}).Create(&trip).Error
	if err != nil {
		log.Fatalf("cannot seed trip table: %v", err)
	}

	var likes = []models.Like{
		models.Like{
			UserID: 1,
			TripID: trip.ID,
		},
		models.Like{
			UserID: 2,
			TripID: trip.ID,
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		err = server.DB.Model(&models.Like{}).Create(&likes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed likes table: %v", err)
		}
	}
	return trip, users, likes, nil
}

func refreshUserTripAndCommentTable() error {
	err := server.DB.DropTableIfExists(&models.Comment{}, &models.Trip{}, &models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.Trip{}, &models.Comment{}, &models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user, trip and comment tables")
	return nil
}

func seedUsersTripsAndComments() (models.Trip, []models.User, []models.Comment, error) {
	// The idea here is: two users can comment one trip
	var err error

	input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)

	var users = []models.User{
		models.User{
			Firstname:   "Jon",
			Lastname:    "Doe",
			Email:       "doe@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Homme",
		},
		models.User{
			Firstname:   "Jenny",
			Lastname:    "Ellen",
			Email:       "jenny@gmail.com",
			Password:    "password",
			Accesslevel: 1,
			Dateofbirth: t,
			Sexe:        "Femme",
		},
	}

	input := "2020-09-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	trip := models.Trip{
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
	}

	err = server.DB.Model(&models.Trip{}).Create(&trip).Error
	if err != nil {
		log.Fatalf("cannot seed trip table: %v", err)
	}

	var comments = []models.Comment{
		models.Comment{
			Body:   "user 1 made this comment",
			UserID: 1,
			TripID: trip.ID,
		},
		models.Comment{
			Body:   "user 2 made this comment",
			UserID: 2,
			TripID: trip.ID,
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		err = server.DB.Model(&models.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}
	}
	return trip, users, comments, nil
}

func refreshUserAndResetPasswordTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.ResetPassword{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.ResetPassword{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user and resetpassword tables")
	return nil
}

// Seed the reset password table with the token
func seedResetPassword() (models.ResetPassword, error) {

	resetDetails := models.ResetPassword{
		Token: "awesometoken",
		Email: "pet@gmail.com",
	}
	err := server.DB.Model(&models.ResetPassword{}).Create(&resetDetails).Error
	if err != nil {
		return models.ResetPassword{}, err
	}
	return resetDetails, nil
}
