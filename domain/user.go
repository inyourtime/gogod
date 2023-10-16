package domain

import (
	"gogod/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	GetByID(userID string, withPwd bool) (*model.User, error)
	GetByEmail(email string, withPwd bool) (*model.User, error)
	All() ([]model.User, error)
	UpdateOne(userID string, updateReq *model.UpdateUserRequest) error
}

type UserUsecase interface {
	GetProfile(userID string) (*model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(id string, req *model.UpdateUserRequest) error
}
