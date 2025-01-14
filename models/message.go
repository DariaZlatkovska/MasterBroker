package models

import "time"

type Message struct {
	ID        string    `json :"id" bson:"_id"`
	Content   string    `json:"content" bson:"content"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
