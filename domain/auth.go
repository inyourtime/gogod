package domain

import (
	"gogod/model"
)

type AuthRepository interface {
	SignUserToken(user *model.User) (*model.Token, error)
}

type AuthUsecase interface {
	Login(req *model.AuthLoginRequest) (*model.AuthLoginResponse, error)
	Register(req *model.User) (*model.User, error)
	Refresh(refreshToken string) (*model.Token, error)
	Google(info model.GoogleInfo) (*model.AuthLoginResponse, error)
}
