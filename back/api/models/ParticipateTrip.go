package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// ParticipateTrip struct
type ParticipateTrip struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	TripID    uint64    `gorm:"not null" json:"trip_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Validate  int       `gorm:"default:0" json:"validate"`
	Trip      Trip      `json:"trip"`
	User      User      `json:"user"`
}

// SaveParticipation : function to participate to a trip
func (p *ParticipateTrip) SaveParticipation(db *gorm.DB) (*ParticipateTrip, error) {

	// Check if the auth user has participated to this trip before
	err := db.Debug().Model(&ParticipateTrip{}).Where("trip_id = ? AND user_id = ?", p.TripID, p.UserID).Take(&p).Error
	if err != nil {
		if err.Error() == "record not found" {
			// The user has not participated to this trip before, so lets save incomming participation
			err = db.Debug().Model(&ParticipateTrip{}).Create(&p).Error
			if err != nil {
				return &ParticipateTrip{}, err
			}
		}
	} else {
		// The user has participated it before, so create a custom error message
		err = errors.New("You already participated to this trip")
		return &ParticipateTrip{}, err
	}
	return p, nil
}

// DeleteParticipation : function to delete a participation of a trip
/*func (p *ParticipateTrip) DeleteParticipation(db *gorm.DB) (*ParticipateTrip, error) {
	var err error
	var deletedParticipation *ParticipateTrip

	err = db.Debug().Model(ParticipateTrip{}).Where("id = ?", p.ID).Take(&p).Error
	if err != nil {
		return &ParticipateTrip{}, err
	} else {
		//If the participation exist, save it in deleted participation and delete it
		deletedParticipation = p
		db = db.Debug().Model(&ParticipateTrip{}).Where("id = ?", p.ID).Take(&ParticipateTrip{}).Delete(&ParticipateTrip{})
		if db.Error != nil {
			fmt.Println("cant delete participation: ", db.Error)
			return &ParticipateTrip{}, db.Error
		}
	}
	return deletedParticipation, nil
}*/

// DeleteParticipationTripAndUser : function to delete a participation with trip id and user id
func (p *ParticipateTrip) DeleteParticipationTripAndUser(db *gorm.DB, pid uint64, uid uint64) (int64, error) {
	db = db.Debug().Model(&ParticipateTrip{}).Where("trip_id = ? and user_id = ?", pid, uid).Take(&ParticipateTrip{}).Delete(&ParticipateTrip{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Participation not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetParticipationInfo : get the infos
func (p *ParticipateTrip) GetParticipationInfo(db *gorm.DB, pid uint64) (*[]ParticipateTrip, error) {

	participations := []ParticipateTrip{}
	err := db.Debug().Model(&ParticipateTrip{}).Where("trip_id = ?", pid).Find(&participations).Error
	if err != nil {
		return &[]ParticipateTrip{}, err
	}
	if len(participations) > 0 {
		for i := range participations {
			err := db.Debug().Model(&User{}).Where("id = ?", participations[i].UserID).Take(&participations[i].User).Error
			if err != nil {
				return &[]ParticipateTrip{}, err
			}
			err2 := db.Debug().Model(&Trip{}).Where("id = ?", participations[i].TripID).Take(&participations[i].Trip).Error
			if err2 != nil {
				return &[]ParticipateTrip{}, err
			}
		}
	}
	return &participations, err
}

// GetParticipationsByUserID : get the infos
func (p *ParticipateTrip) GetParticipationsByUserID(db *gorm.DB, pid uint64) (*[]ParticipateTrip, error) {

	participations := []ParticipateTrip{}
	err := db.Debug().Model(&ParticipateTrip{}).Where("user_id = ?", pid).Find(&participations).Error
	if err != nil {
		return &[]ParticipateTrip{}, err
	}
	if len(participations) > 0 {
		for i := range participations {
			err := db.Debug().Model(&Trip{}).Where("id = ?", participations[i].TripID).Take(&participations[i].Trip).Error
			if err != nil {
				return &[]ParticipateTrip{}, err
			}
		}
	}
	return &participations, err
}

// DeleteUserParticipations : When a user is deleted, we also delete the participations that the user had
func (p *ParticipateTrip) DeleteUserParticipations(db *gorm.DB, uid uint64) (int64, error) {
	participations := []ParticipateTrip{}
	db = db.Debug().Model(&ParticipateTrip{}).Where("user_id = ?", uid).Find(&participations).Delete(&participations)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// DeleteTripParticipations : When a trip is deleted, we also delete the participations that the trip had
func (p *ParticipateTrip) DeleteTripParticipations(db *gorm.DB, pid uint64) (int64, error) {
	participations := []ParticipateTrip{}
	db = db.Debug().Model(&ParticipateTrip{}).Where("trip_id = ?", pid).Find(&participations).Delete(&participations)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
