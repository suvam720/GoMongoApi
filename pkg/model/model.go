package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskBody struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Text      string             `json:"text,omitempty"`
	Completed bool               `json:"completed,omitempty"`
}
