package models

import (
	"time"
)
// Message struct
type ELKNotification struct {
	Message     string   	` json:"message"`
	Type      	string   	` json:"type"`
	TargetName  string   	` json:"target_name"`
	CreatedAt 	time.Time 	` json:"created_at"`
	TargetID 	uint64		` json:"target_id"`
	TargetEmail string		` json:"target_email"`
}

const MappingNotifConst = `
{
	"mappings":{
		"notification":{
			"properties":{
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"type":{
					"type":"string",
					"store": true,
					"fielddata": true
				},
				"target_name":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"created_at":{
					"type":"keyword"
				},
				"target_id":{
					"type":"integer"
				},
				"target_email":{
					"type":"string",
				}
			}
		}
	}
}`
