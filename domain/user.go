package domain

import (
	"gogod/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface{
	Create(user *model.User) (*model.User, error)
	GetByID(_id primitive.ObjectID, withPwd bool) (*model.User, error)
	GetByEmail(email string, withPwd bool) (*model.User, error)
	All() ([]model.User, error)
	UpdateOne(_id primitive.ObjectID, updateReq *model.UpdateUserRequest) error
}

type UserUsecase interface{
	GetProfile(id string) (*model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(id string, req *model.UpdateUserRequest) error
}
