package Request

import (
	"time"
)

type CreatePostRequest struct {
	Title     string    `json:"title" bson:"title"  binding:"required"`
	Content   string    `json:"content" bson:"content"  binding:"required"`
	Image     string    `json:"image" bson:"image"  binding:"required"`
	User      string    `json:"user" bson:"user"  binding:"required"`
	CreateAt  time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type UpdatePost struct {
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	User      string             `json:"user,omitempty" bson:"user,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

