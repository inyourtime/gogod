package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Desc   string             `json:"desc" bson:"desc"`
	Image  string             `json:"image" bson:"image"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
}

type Comment struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}
