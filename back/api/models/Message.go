package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const MappingMessageConst = `
{
	"mappings":{
		"message":{
			"properties":{
				"body":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"created_at":{
					"type":"keyword"
				},
				"SenderName":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"SenderId":{
					"type":"integer"
				},
				"conversation_id":{
					"type":"integer"
				},
				"group":{
					"type":"boolean"
				}
			}
		}
	}
}`


// Message struct
type ELKMessage struct {
	Body      	string   	` json:"body"`
	CreatedAt 	time.Time 	` json:"created_at"`
	SenderId 	uint64		` json:"sender_id"`
	SenderName 	string		` json:"sender_name"`
	Groupe 		bool		` json:"group"`
	ConversationId 	uint64	` json:"conversation_id"`
	ConversationName 	string		` json:"conversation_name"`
	ConversationParticipantIds []uint64  ` json:"conversation_participant_ids"`
}


// Message struct
type Message struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Body      string    `gorm:"text;not null;" json:"body"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy uint64	`gorm:"default:0" json:"created_by"`
}


// TableName : Gorm related
func (m *Message) TableName() string {
	return "message"
}


// GetAllMessage : function to find all Messages
func (m *Message) GetAllMessage(db *gorm.DB) (*[]Message, error) {
	var err error
	message := []Message{}
	err = db.Debug().Model(&Message{}).Limit(100).Find(&message).Error
	if err != nil {
		return &[]Message{}, err
	}
	return &message, err
}

// Save message : Method Save message, triggered on every saved on trip struct
func (m *Message) SaveMessage(db *gorm.DB) (*Message, error) {
	var err error
	err = db.Debug().Model(&Message{}).Create(&m).Error
	if err != nil {
		return &Message{}, err
	}
	if m.ID != 0 {
	//	err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Message{}, err
		}
	}
	return m, nil
}