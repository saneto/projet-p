package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Like struct
type Like struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	TripID    uint64    `gorm:"not null" json:"trip_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Trip      Trip      `json:"trip"`
	User      User      `json:"user"`
}

// SaveLike : function to like a trip
func (l *Like) SaveLike(db *gorm.DB) (*Like, error) {

	// Check if the auth user has liked this trip before:
	err := db.Debug().Model(&Like{}).Where("trip_id = ? AND user_id = ?", l.TripID, l.UserID).Take(&l).Error
	if err != nil {
		if err.Error() == "record not found" {
			// The user has not liked this trip before, so lets save incomming like:
			err = db.Debug().Model(&Like{}).Create(&l).Error
			if err != nil {
				return &Like{}, err
			}
		}
	} else {
		// The user has liked it before, so create a custom error message
		err = errors.New("You already liked this trip")
		return &Like{}, err
	}
	return l, nil
}

// DeleteLike : function to delete a like of a trip
/*func (l *Like) DeleteLike(db *gorm.DB) (*Like, error) {
	var err error
	var deletedLike *Like

	err = db.Debug().Model(Like{}).Where("id = ?", l.ID).Take(&l).Error
	if err != nil {
		return &Like{}, err
	} else {
		//If the like exist, save it in deleted like and delete it
		deletedLike = l
		db = db.Debug().Model(&Like{}).Where("id = ?", l.ID).Take(&Like{}).Delete(&Like{})
		if db.Error != nil {
			fmt.Println("cant delete like: ", db.Error)
			return &Like{}, db.Error
		}
	}
	return deletedLike, nil
}*/

// DeleteLikeTripAndUser : function to delete a like with trip id and user id
func (l *Like) DeleteLikeTripAndUser(db *gorm.DB, pid uint64, uid uint64) (int64, error) {
	db = db.Debug().Model(&Like{}).Where("trip_id = ? and user_id = ?", pid, uid).Take(&Like{}).Delete(&Like{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Like not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetLikesInfo : get the infos
func (l *Like) GetLikesInfo(db *gorm.DB, pid uint64) (*[]Like, error) {

	likes := []Like{}
	err := db.Debug().Model(&Like{}).Where("trip_id = ?", pid).Find(&likes).Error
	if err != nil {
		return &[]Like{}, err
	}
	if len(likes) > 0 {
		for i := range likes {
			err := db.Debug().Model(&User{}).Where("id = ?", likes[i].UserID).Take(&likes[i].User).Error
			if err != nil {
				return &[]Like{}, err
			}
			err2 := db.Debug().Model(&Trip{}).Where("id = ?", likes[i].TripID).Take(&likes[i].Trip).Error
			if err2 != nil {
				return &[]Like{}, err
			}
		}
	}
	return &likes, err
}

// GetLikesByUserID : get the infos
func (l *Like) GetLikesByUserID(db *gorm.DB, pid uint64) (*[]Like, error) {

	likes := []Like{}
	err := db.Debug().Model(&Like{}).Where("user_id = ?", pid).Find(&likes).Error
	if err != nil {
		return &[]Like{}, err
	}
	if len(likes) > 0 {
		for i := range likes {
			err := db.Debug().Model(&Trip{}).Where("id = ?", likes[i].TripID).Take(&likes[i].Trip).Error
			if err != nil {
				return &[]Like{}, err
			}
		}
	}
	return &likes, err
}

// DeleteUserLikes : When a user is deleted, we also delete the likes that the user had
func (l *Like) DeleteUserLikes(db *gorm.DB, uid uint64) (int64, error) {
	likes := []Like{}
	db = db.Debug().Model(&Like{}).Where("user_id = ?", uid).Find(&likes).Delete(&likes)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// DeleteTripLikes : When a trip is deleted, we also delete the likes that the trip had
func (l *Like) DeleteTripLikes(db *gorm.DB, pid uint64) (int64, error) {
	likes := []Like{}
	db = db.Debug().Model(&Like{}).Where("trip_id = ?", pid).Find(&likes).Delete(&likes)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// FindAllLikes : function to find all likes
func (l *Like) FindAllLikes(db *gorm.DB) (*[]Like, error) {
	var err error
	likes := []Like{}
	err = db.Debug().Model(&Like{}).Limit(100).Order("created_at desc").Find(&likes).Group("trip_id").Error
	if err != nil {
		return &[]Like{}, err
	}

	return &likes, nil
}
