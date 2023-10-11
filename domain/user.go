package domain

import (
	"gogod/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface{
	Create(user *model.User) (*model.User, error)
	GetByID(_id primitive.ObjectID, withPwd bool) (*model.User, error)
	GetByEmail(email string, withPwd bool) (*model.User, error)
}

type UserUsecase interface{
	GetProfile(id string) (*model.User, error)
}
