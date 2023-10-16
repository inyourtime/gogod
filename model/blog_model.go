package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	BlogID    string             `json:"blogId" bson:"blogId"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Desc      string             `json:"desc" bson:"desc" validate:"required"`
	Image     string             `json:"image" bson:"image" validate:"required"`
	LikesBy   []Like             `json:"likesBy" bson:"likesBy"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedBy CreatedBy          `json:"createdBy" bson:"createdBy"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CommentID string             `json:"commentId" bson:"commentId"`
	Content   string             `json:"content" bson:"content"`
	LikesBy   []Like             `json:"likesBy" bson:"likesBy"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedBy CreatedBy          `json:"createdBy" bson:"createdBy"`
	BlogID    string             `json:"-" bson:"blogId"`
}

type Like struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	CreatedBy CreatedBy `json:"createdBy" bson:"createdBy"`
}

type CreatedBy struct {
	UserID string `json:"userId" bson:"userId"`
	Name   string `json:"name" bson:"name"`
}

type BlogUpdateRequest struct {
	BlogID    string    `json:"blogId" bson:"blogId" validate:"required"`
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Desc      string    `json:"desc,omitempty" bson:"desc,omitempty"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
