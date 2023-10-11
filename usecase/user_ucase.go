package usecase

import (
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
