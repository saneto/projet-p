package controllers

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"os"

	"gitlab.com/esgi/lftrip/BackLFTrip/api/middlewares"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olahol/melody"

	"github.com/olivere/elastic"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	
)

// Server : Struct
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine 
	Melody *melody.Melody
	Elk *elastic.Client
	Context context.Context
}


var errList = make(map[string]string)

// Initialize : connection to our database
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	
	var err error

	server.Context = context.Background()
/*	server.Elk, err = elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")),
	elastic.SetSniff(false))*/
	server.Elk, err = elastic.NewClient(
		elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")),	
		elastic.SetBasicAuth(os.Getenv("ELASTICSEARCH_USER"), os.Getenv("ELASTICSEARCH_PASSWORD")),
		 elastic.SetSniff(false))
	if err != nil {
		fmt.Println("error%s ", err)
	}else{
		exists, err := server.Elk.IndexExists(os.Getenv("ELK_TRIPS_INDEX")).Do(server.Context)
		if err != nil {
			server.Elk.DeleteIndex(os.Getenv("ELK_TRIPS_INDEX")).Do(server.Context)
			exists = false
		}
		if !exists {
			_, err := server.Elk.CreateIndex(os.Getenv("ELK_TRIPS_INDEX")).BodyString(models.MappingTripConst).Do(server.Context)
			if err != nil {
				fmt.Println("error create index %s ", err)
			}
		}
		exists, err = server.Elk.IndexExists(os.Getenv("ELK_MESSAGE_INDEX")).Do(server.Context)
		if err != nil {
			server.Elk.DeleteIndex(os.Getenv("ELK_MESSAGE_INDEX")).Do(server.Context)
			exists = false
		}
		if !exists {
			_, err := server.Elk.CreateIndex(os.Getenv("ELK_MESSAGE_INDEX")).BodyString(models.MappingMessageConst).Do(server.Context)
			if err != nil {
				fmt.Println("error%s ", err)
			}
		}

		exists, err = server.Elk.IndexExists(os.Getenv("ELK_NOTIF_INDEX")).Do(server.Context)
		if err != nil {
			server.Elk.DeleteIndex(os.Getenv("ELK_NOTIF_INDEX")).Do(server.Context)
			exists = false
		}
		if !exists {
			_, err := server.Elk.CreateIndex(os.Getenv("ELK_NOTIF_INDEX")).BodyString(models.MappingNotifConst).Do(server.Context)
			if err != nil {
				fmt.Printf("error%s ", err)
			}
		}
	}

	server.Melody = melody.New()
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error connecting to postgres:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.ResetPassword{},
		&models.Like{},
		&models.Comment{},
		&models.ParticipateTrip{},
		&models.UserPreferences{},
	)

	
	//server.Elk.DeleteIndex(os.Getenv("ELK_TRIPS_INDEX")).Do(server.Context)

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())
	server.initializeRoutes()
}

// Run : Status
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
