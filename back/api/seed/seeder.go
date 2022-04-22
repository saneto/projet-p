package seed

import (
	"log"
	//"time"

	"github.com/jinzhu/gorm"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)

// Load : Validation and Join
func Load(db *gorm.DB) {

	/*input1 := "1996-02-08"
	layout1 := "2006-01-02"
	t, _ := time.Parse(layout1, input1)*/

/*	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input2 := "2020-10-01"
	layout2 := "2006-01-02"
	end, _ := time.Parse(layout2, input2)

	var users = []models.User{
		models.User{
			Firstname:        "Steve",
			Lastname:         "Victor",
			Email:            "steve@gmail.com",
			Accesslevel:      1,
			Password:         "password",
			Dateofbirth:      "1996-02-08",
			Sexe:             "Homme",
			City:             "Paris",
			IsGuide:          "false",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Developer Go",
		},
		models.User{
			Firstname:        "Kevin",
			Lastname:         "Feige",
			Email:            "feige@gmail.com",
			Accesslevel:      1,
			Password:         "password",
			Dateofbirth:      "1996-02-08",
			Sexe:             "Homme",
			City:             "Paris",
			IsGuide:          "true",
			CountryGuide:     "France",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Student IT",
		},
		models.User{
			Firstname:        "Karim",
			Lastname:         "Benzema",
			Email:            "k@benzema.io",
			Accesslevel:      1,
			Password:         "password",
			Dateofbirth:      "1996-02-08",
			Sexe:             "Homme",
			City:             "Paris",
			IsGuide:          "false",
			PhoneNumber:      "0625321458",
			DepartureAirport: "Orly",
			Description:      "Soccer player",
		},
	}

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
			IsActive:    true,
			TripType:    "Culturel",
			Program:     "On partirai de Nice et nous longerons toute la côte Est jusqu'à Menton.",
			Lodging:     "Principalement du airbnb",
			Budget:      350,
			AuthorID:    1,
			Latitude:    "46.227638",
			Longitude:   "2.213749",
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
			IsActive:    true,
			TripType:    "Détente",
			Program:     "Pas encore fait, mais si vous êtes intéressé, envoyer moi un MP et on y réfléchira ensemble.",
			Lodging:     "Principalement du airbnb ou auberge de jeunesse.",
			Budget:      500,
			AuthorID:    2,
			Latitude:    "40.463667",
			Longitude:   "-3.74922",
		},
		models.Trip{
			Country:     "Italie",
			Title:       "Trek en Italie, ça vous dit ?",
			Description: "Voyage pour les fans de trekking et de nature !",
			StartDate:   start,
			EndDate:     end,
			NbDays:      7,
			MiddleAge:   28,
			NbTraveler:  3,
			TripType:    "Circuit",
			IsActive:    true,
			Program:     "Pas encore décidé.",
			Lodging:     "Une tente !",
			Budget:      250,
			AuthorID:    1,
			Latitude:    "41.87194",
			Longitude:   "12.56738",
		},
	}

	var comments = []models.Comment{
		models.Comment{
			Body:   "Voyage intéressant",
			UserID: 1,
			TripID: 1,
		},
		models.Comment{
			Body:   "Pouvez-vous nous fournir plus de détails ?",
			UserID: 2,
			TripID: 1,
		},
		models.Comment{
			Body:   "Excellent programme !",
			UserID: 1,
			TripID: 2,
		},
	}

	var userPreferences = []models.UserPreferences{
		models.UserPreferences{
			UserID:                   1,
			TravelerType:             "Tripizien Accompli",
			ShyOutgoing:              "Timide",
			CalmOrSpark:              "Un vrai piment de Cayenne",
			SundayMikeHornAdventurer: "Aventurier comme Mike Horn",
			AthleteSportKezako:       "Le sport c'est la vie",
			LuxuryCharmingLodging:    "Un petit lieu de charme",
			FatMorning:               "Frais et paré à partir à 7h12 !",
			Food:                     "Aucun régime",
			ScooteredDisorganized:    "Organisé comme un scout",
			TrekkingBeachTowel:       "Partir à l’aventure !",
		},
		models.UserPreferences{
			UserID:                   2,
			TravelerType:             "Tripizien des temps modernes",
			ShyOutgoing:              "Extraverti",
			CalmOrSpark:              "Un vrai piment de Cayenne",
			SundayMikeHornAdventurer: "Aventurier du dimanche",
			AthleteSportKezako:       "Le sport ? Ca se mange ?",
			LuxuryCharmingLodging:    "Plus y’a d’étoiles, plus je brille",
			FatMorning:               "Grasse mat’ c’est le kiff !",
			Food:                     "Vegan",
			ScooteredDisorganized:    "Désorganisé depuis que je suis né",
			TrekkingBeachTowel:       "Serviette, Plage et Bronzette",
		},
		models.UserPreferences{
			UserID:                   3,
			TravelerType:             "Tripizien culturel",
			ShyOutgoing:              "Timide",
			CalmOrSpark:              "Une petite perle de lait",
			SundayMikeHornAdventurer: "Aventurier comme Mike Horn",
			AthleteSportKezako:       "Le sport c'est la vie",
			LuxuryCharmingLodging:    "Un petit lieu de charme",
			FatMorning:               "Grasse mat’ c’est le kiff !",
			Food:                     "Halal",
			ScooteredDisorganized:    "Organisé comme un scout",
			TrekkingBeachTowel:       "Serviette, Plage et Bronzette",
		},
	}*/


	err := db.Debug().DropTableIfExists(&models.Comment{}, &models.ParticipateTrip{}, &models.Like{}, &models.UserPreferences{}, &models.Trip{}, &models.User{}, &models.Message{}, &models.Chat{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().DropTableIfExists(&models.Chat{}, "users").Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.ParticipateTrip{}, &models.Like{}, &models.Trip{}, &models.Comment{}, &models.UserPreferences{}, &models.Message{}, &models.Chat{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Trip{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.UserPreferences{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Comment{}).AddForeignKey("trip_id", "trips(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.ParticipateTrip{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.ParticipateTrip{}).AddForeignKey("trip_id", "trips(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Like{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Like{}).AddForeignKey("trip_id", "trips(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

/*	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		trips[i].AuthorID = users[i].ID
		//trips[i].Author = users[i]
		err = db.Debug().Model(&models.Trip{}).Create(&trips[i]).Association("Author").Error
		if err != nil {
			log.Fatalf("cannot seed trips table: %v", err)
		}

		err = db.Debug().Model(&models.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}

		err = db.Debug().Model(&models.UserPreferences{}).Create(&userPreferences[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user preferences table: %v", err)
		}
	}*/
}
