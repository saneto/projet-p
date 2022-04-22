package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/twinj/uuid"
)

// UserPreferences struct
type UserPreferences struct {
	UUID                     uuid.UUID `gorm:"type:uuid;unique_index;" json:"uuid"`
	ID                       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID                   uint64    `gorm:"not null" json:"user_id"`
	TravelerType             string    `gorm:"size:200;null" json:"traveler_type"`
	ShyOutgoing              string    `gorm:"size:200;null" json:"shy_outgoing"`
	CalmOrSpark              string    `gorm:"size:200;null" json:"calm_or_spark"`
	SundayMikeHornAdventurer string    `gorm:"size:200;null" json:"sunday_mike_horn_adventurer"`
	AthleteSportKezako       string    `gorm:"size:200;null" json:"athlete_sport_kezako"`
	LuxuryCharmingLodging    string    `gorm:"size:200;null" json:"luxury_charming_lodging"`
	FatMorning               string    `gorm:"size:200;null" json:"fat_morning"`
	Food                     string    `gorm:"size:200;null" json:"food"`
	ScooteredDisorganized    string    `gorm:"size:200;null" json:"scootered_disorganized"`
	TrekkingBeachTowel       string    `gorm:"size:200;null" json:"trekking_beach_towel"`
	User                     User      `json:"user"`
	Users                    []*User   `gorm:"many2many:userpreferences_users;association_foreignkey:id;foreignkey:id" json:"users,omitempty"`
	CreatedAt                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName : Gorm related
func (p *UserPreferences) TableName() string {
	return "userpreferences"
}

// BeforeSave : Method before Save
func (p *UserPreferences) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on vote struct
func (p *UserPreferences) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Prepare : prepare statements
func (p *UserPreferences) Prepare() {
	p.ID = 0
	p.TravelerType = html.UnescapeString(strings.TrimSpace(p.TravelerType))
	p.ShyOutgoing = html.UnescapeString(strings.TrimSpace(p.ShyOutgoing))
	p.CalmOrSpark = html.UnescapeString(strings.TrimSpace(p.CalmOrSpark))
	p.SundayMikeHornAdventurer = html.UnescapeString(strings.TrimSpace(p.SundayMikeHornAdventurer))
	p.AthleteSportKezako = html.UnescapeString(strings.TrimSpace(p.AthleteSportKezako))
	p.LuxuryCharmingLodging = html.UnescapeString(strings.TrimSpace(p.LuxuryCharmingLodging))
	p.FatMorning = html.UnescapeString(strings.TrimSpace(p.FatMorning))
	p.Food = html.UnescapeString(strings.TrimSpace(p.Food))
	p.ScooteredDisorganized = html.UnescapeString(strings.TrimSpace(p.ScooteredDisorganized))
	p.TrekkingBeachTowel = html.UnescapeString(strings.TrimSpace(p.TrekkingBeachTowel))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// SaveUserPreferences : function to save a user preferences linked to a user
func (p *UserPreferences) SaveUserPreferences(db *gorm.DB) (*UserPreferences, error) {
	err := db.Debug().Create(&p).Error
	if err != nil {
		return &UserPreferences{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &UserPreferences{}, err
		}
	}
	return p, nil
}

// FindAllUserPreferences : function to find all user preferences
func (p *UserPreferences) FindAllUserPreferences(db *gorm.DB) (*[]UserPreferences, error) {
	var err error
	userPreferences := []UserPreferences{}
	err = db.Debug().Model(&UserPreferences{}).Limit(100).Preload("Users").Order("created_at desc").Find(&userPreferences).Error
	if err != nil {
		return &[]UserPreferences{}, err
	}
	if len(userPreferences) > 0 {
		for i := range userPreferences {
			err := db.Debug().Model(&User{}).Where("id = ?", userPreferences[i].UserID).Take(&userPreferences[i].User).Error
			if err != nil {
				return &[]UserPreferences{}, err
			}
		}
	}
	return &userPreferences, nil
}

// FindUserPreferenceByID : function to find a user preferences with an ID
func (p *UserPreferences) FindUserPreferenceByID(db *gorm.DB, pid uint64) (*UserPreferences, error) {
	var err error
	err = db.Debug().Model(&UserPreferences{}).Where("id = ?", pid).Preload("Users").Take(&p).Error
	if err != nil {
		return &UserPreferences{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &UserPreferences{}, err
		}
	}
	return p, nil
}

// UpdateAUserPreferences : function to update a user preferences
func (p *UserPreferences) UpdateAUserPreferences(db *gorm.DB, pid uint64) (*UserPreferences, error) {
	var err error
	db = db.Debug().Model(&UserPreferences{}).Where("id = ?", pid).Take(&UserPreferences{}).UpdateColumns(
		map[string]interface{}{
			"traveler_type":               p.TravelerType,
			"shy_outgoing":                p.ShyOutgoing,
			"calm_or_spark":               p.CalmOrSpark,
			"sunday_mike_horn_adventurer": p.SundayMikeHornAdventurer,
			"athlete_sport_kezako":        p.AthleteSportKezako,
			"luxury_charming_lodging":     p.LuxuryCharmingLodging,
			"fat_morning":                 p.FatMorning,
			"food":                        p.Food,
			"scootered_disorganized":      p.ScooteredDisorganized,
			"trekking_beach_towel":        p.TrekkingBeachTowel,
			"updated_at":                  time.Now(),
		},
	)
	err = db.Debug().Model(&UserPreferences{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &UserPreferences{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &UserPreferences{}, err
		}
	}
	return p, nil
}

// DeleteAUserPreferences : function to delete a user preferences
func (p *UserPreferences) DeleteAUserPreferences(db *gorm.DB, pid uint64, uid uint64) (int64, error) {
	db = db.Debug().Model(&UserPreferences{}).Where("id = ? and user_id = ?", pid, uid).Take(&UserPreferences{}).Delete(&UserPreferences{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("User Preferences not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// FindUserPreferenceByUserID : function to get all user preferences for a user
func (p *UserPreferences) FindUserPreferenceByUserID(db *gorm.DB, uid uint64) (*[]UserPreferences, error) {

	var err error
	userPref := []UserPreferences{}
	err = db.Debug().Model(&UserPreferences{}).Where("user_id = ?", uid).Limit(100).Order("created_at desc").Find(&userPref).Error
	if err != nil {
		return &[]UserPreferences{}, err
	}
	if len(userPref) > 0 {
		for i := range userPref {
			err := db.Debug().Model(&User{}).Where("id = ?", userPref[i].UserID).Take(&userPref[i].User).Error
			if err != nil {
				return &[]UserPreferences{}, err
			}
		}
	}
	return &userPref, nil
}

// DeleteUserPreferencesUserDeleted : When a user is deleted, we also delete the user preference that the user had
func (p *UserPreferences) DeleteUserPreferencesUserDeleted(db *gorm.DB, uid uint64) (int64, error) {
	userPref := []UserPreferences{}
	db = db.Debug().Model(&UserPreferences{}).Where("user_id = ?", uid).Find(&userPref).Delete(&userPref)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
