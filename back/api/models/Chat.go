package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Chat struct
type Chat struct {
	ID        	uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name  	  	string    `gorm:"text;not null; json:"name"`
	CreateBy  	uint64    `gorm:"null; json:"create_by"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Users     	[]*User   `gorm:"many2many:chat_users;association_foreignkey:id;foreignkey:id" json:"users,omitempty"`
	Messages  	[]*Message   `gorm:"many2many:chat_message;association_foreignkey:id;foreignkey:id" json:"messages,omitempty"`
}


// TableName : Gorm related
func (c *Chat) TableName() string {
	return "chat"
}

// BeforeUpdate is gorm hook that is triggered on every updated on vote struct
func (c *Chat) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// GetAllChat : function to find all Chats
func (m *Chat) GetAllChat(db *gorm.DB) (*[]Chat, error) {
	var err error
	chat := []Chat{}
	err = db.Debug().Model(&Chat{}).Limit(100).Find(&chat).Error
	if err != nil {
		return &[]Chat{}, err
	}
	return &chat, err
}

// Save Chat : Method Save Chat, triggered on every saved on trip struct
func (m *Chat) SaveChat(db *gorm.DB) (*Chat, error) {
	var err error
	err = db.Debug().Model(&Chat{}).Create(&m).Error
	if err != nil {
		return &Chat{}, err
	}
	if m.ID != 0 {
	//	err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Chat{}, err
		}
	}
	return m, nil
}/**/