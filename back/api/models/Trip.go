package models

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/twinj/uuid"
)

const MappingTripConst = `
{
	"mappings":{
		"trip":{
			"properties":{
				"country":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"base64":{
					"type":"keyword"
				},
				"start_date":{
					"type":"date"
				},
				"end_date":{
					"type":"date"
				},
				"description":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"lodging":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"title":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"budget":{
					"type":"integer"
				},
				"nb_days":{
					"type":"integer"
				},
				"nb_traveler":{
					"type":"integer"
				}
			}
		}
	}
}`

// Trip Struct
type Trip struct {
	UUID         uuid.UUID  `gorm:"type:uuid;unique_index;" json:"uuid"`
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Country      string     `gorm:"size:255;not null;" json:"country"`
	Longitude    string     `gorm:"text;null;" json:"longitude"`
	Latitude     string     `gorm:"text;null;" json:"latitude"`
	Title        string     `gorm:"size:255;not null;" json:"title"`
	TripType     string     `gorm:"size:255;not null;" json:"trip_type"`
	IsActive     bool       `gorm:"boolean;not null;" json:"is_active"`
	Description  string     `gorm:"text;not null;" json:"description"`
	StartDate    time.Time  `gorm:"not null;" json:"start_date"`
	EndDate      time.Time  `gorm:"not null;" json:"end_date"`
	NbDays       int        `gorm:"not null;" json:"nb_days"`
	MiddleAge    int        `gorm:"not null;" json:"middle_age"`
	NbTraveler   int        `gorm:"not null;" json:"nb_traveler"`
	ImageTrip    string     `gorm:"size:255;null;" json:"image_trip"`
	Program      string     `gorm:"text;null;" json:"program"`
	ImageProgram string     `gorm:"size:255;null;" json:"image_program"`
	Base64       string     `gorm:"type:text;null" json:"base64"`
	Lodging      string     `gorm:"text;null;" json:"lodging"`
	Budget       int        `gorm:"not null;" json:"budget"`
	Author       User       `json:"author"`
	AuthorID     uint64     `gorm:"not null" json:"author_id"`
	Comments     []*Comment `gorm:"many2many:trip_comments;association_foreignkey:id;foreignkey:id" json:"comments,omitempty"`
	Users        []*User    `gorm:"many2many:trip_users;association_foreignkey:id;foreignkey:id" json:"users,omitempty"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName : Gorm related
func (t *Trip) TableName() string {
	return "trips"
}

// BeforeSave : Method before Save
func (t *Trip) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on vote struct
func (t *Trip) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Validate : function to check the data
func (t *Trip) Validate() map[string]string {
	var err error
	var errorMessages = make(map[string]string)

	if t.Title == "" {
		err = errors.New("Titre requis")
		errorMessages["Required_title"] = err.Error()
	}
	if t.Country == "" {
		err = errors.New("Pays de destination requis")
		errorMessages["Required_country"] = err.Error()
	}
	if t.TripType == "" {
		err = errors.New("Type de voyage requis")
		errorMessages["Required_trip_type"] = err.Error()
	}
	if t.Description == "" {
		err = errors.New("Description requis")
		errorMessages["Required_description"] = err.Error()
	}
	if t.StartDate.IsZero() {
		err = errors.New("Date de départ requis")
		errorMessages["Required_start_date"] = err.Error()
	}
	if t.EndDate.IsZero() {
		err = errors.New("Date de fin requis")
		errorMessages["Required_end_date"] = err.Error()
	}
	if t.AuthorID < 1 {
		err = errors.New("Required Author")
		errorMessages["Required_author"] = err.Error()
	}
	if t.NbDays == 0 {
		err = errors.New("Nombre de jours requis")
		errorMessages["Required_nb_days"] = err.Error()
	}
	if t.Budget == 0 {
		err = errors.New("Budget requis")
		errorMessages["Required_budget"] = err.Error()
	}
	return errorMessages
}

//Prepare : prepare a trip
func (t *Trip) Prepare() {
	t.Country = html.EscapeString(strings.TrimSpace(t.Country))
	t.Title = html.UnescapeString(strings.TrimSpace(t.Title))
	t.Description = html.UnescapeString(strings.TrimSpace(t.Description))
	t.Program = html.UnescapeString(strings.TrimSpace(t.Program))
	t.Lodging = html.UnescapeString(strings.TrimSpace(t.Lodging))
	t.Longitude = html.EscapeString(strings.TrimSpace(t.Longitude))
	t.Latitude = html.EscapeString(strings.TrimSpace(t.Latitude))
	t.Author = User{}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

// SaveTrip : Method Save Trip, triggered on every saved on trip struct
func (t *Trip) SaveTrip(db *gorm.DB) (*Trip, error) {
	var err error
	err = db.Model(&Trip{}).Create(&t).Error
	if err != nil {
		fmt.Println(err)
		return &Trip{}, err
	}
	/*if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Trip{}, err
		}
	}*/
	return t, nil
}

// FindAllTrips : function to find all trips
func (t *Trip) FindAllTrips(db *gorm.DB) (*[]Trip, error) {
	var err error
	trips := []Trip{}
	err = db.Debug().Model(&Trip{}).Where("is_active = ?", true).Limit(100).Preload("Users").Preload("Comments").Order("created_at desc").Find(&trips).Error
	if err != nil {
		return &[]Trip{}, err
	}
	if len(trips) > 0 {
		for i := range trips {
			err := db.Debug().Model(&User{}).Where("id = ?", trips[i].AuthorID).Take(&trips[i].Author).Error
			if err != nil {
				return &[]Trip{}, err
			}
		}
	}
	return &trips, nil
}

// FindTripByID : function to find a trip with an ID
func (t *Trip) FindTripByID(db *gorm.DB, pid uint64) (*Trip, error) {
	var err error
	err = db.Debug().Model(&Trip{}).Where("id = ?", pid).Preload("Users").Preload("Comments").Take(&t).Error
	if err != nil {
		return &Trip{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Trip{}, err
		}
	}
	return t, nil
}

// UpdateATrip : function to update a trip
func (t *Trip) UpdateATrip(db *gorm.DB, pid uint64) (*Trip, error) {

	db = db.Debug().Model(&Trip{}).Where("id = ?", pid).Take(&Trip{}).UpdateColumns(
		map[string]interface{}{
			"title":       t.Title,
			"country":     t.Country,
			"longitude":   t.Longitude,
			"latitude":    t.Latitude,
			"description": t.Description,
			"trip_type":   t.TripType,
			"is_active":   t.IsActive,
			"start_date":  t.StartDate,
			"end_date":    t.EndDate,
			"nb_days":     t.NbDays,
			"middle_age":  t.MiddleAge,
			"nb_traveler": t.NbTraveler,
			"program":     t.Program,
			"lodging":     t.Lodging,
			"budget":      t.Budget,
			"updated_at":  time.Now(),
			"base64":      t.Base64,
		},
	)
	if db.Error != nil {
		return &Trip{}, db.Error
	}
	return t, nil
}

// DeleteATrip : function to delete a trip
func (t *Trip) DeleteATrip(db *gorm.DB, pid uint64, uid uint64) (int64, error) {
	db = db.Debug().Model(&Trip{}).Where("id = ? and author_id = ?", pid, uid).Take(&Trip{}).Delete(&Trip{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Trip not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// FindUserTrips : function to get all trips for a user
func (t *Trip) FindUserTrips(db *gorm.DB, uid uint64) (*[]Trip, error) {

	var err error
	trips := []Trip{}
	err = db.Debug().Model(&Trip{}).Where("author_id = ?", uid).Limit(100).Order("created_at desc").Find(&trips).Error
	if err != nil {
		return &[]Trip{}, err
	}
	if len(trips) > 0 {
		for i := range trips {
			err := db.Debug().Model(&User{}).Where("id = ?", trips[i].AuthorID).Take(&trips[i].Author).Error
			if err != nil {
				return &[]Trip{}, err
			}
		}
	}
	return &trips, nil
}

// DeleteUserTrips : When a user is deleted, we also delete the trip that the user had
func (t *Trip) DeleteUserTrips(db *gorm.DB, uid uint64) (int64, error) {
	trips := []Trip{}
	db = db.Debug().Model(&Trip{}).Where("author_id = ?", uid).Find(&trips).Delete(&trips)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
