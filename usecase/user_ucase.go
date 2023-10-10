package usecase

import (
	"gogod/domain"
	"gogod/model"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) GetProfile(id string) (*model.User, error) {
	return nil, nil
}
