package usecase

import (
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	// retrive object_id
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidObjID
	}
	// retrive user
	currentUser, err := u.userRepo.GetByID(objectID, false)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if currentUser == nil {
		return nil, domain.ErrUserNotFound
	}
	return currentUser, nil
}

func (u *userUsecase) GetAllUser() ([]model.User, error) {
	users, err := u.userRepo.All()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) UpdateUser(id string, req *model.UpdateUserRequest) error {
	// retrive object_id
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidObjID
	}

	if req.Password != "" {
		// hash
		bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return err
		}
		req.Password = string(bytes)
	}

	err = u.userRepo.UpdateOne(objectID, req)
	if err != nil {
		if err == domain.ErrUserNotFound {
			return err
		}
		logger.Error(err)
		return err
	}
	return nil
}
