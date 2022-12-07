package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DBPost struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"  binding:"required"`
	Content   string             `json:"content" bson:"content"  binding:"required"`
	Image     string             `json:"image" bson:"image"  binding:"required"`
	User      string             `json:"user" bson:"user"  binding:"required"`
	CreateAt  time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}


