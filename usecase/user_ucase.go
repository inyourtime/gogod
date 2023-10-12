package usecase

import (
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"

	"github.com/gofiber/fiber/v2"
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
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid object id")
	}
	// retrive user
	currentUser, err := u.userRepo.GetByID(objectID, false)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if currentUser == nil {
		return nil, fiber.ErrNotFound
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
		return fiber.NewError(fiber.StatusBadRequest, "Invalid object id")
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
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		logger.Error(err)
		return err
	}
	return nil
}
