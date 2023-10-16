package usecase_test

import (
	"errors"
	"gogod/domain/mock"
	"gogod/model"
	"gogod/usecase"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func TestGetProfile(t *testing.T) {
	mockID := primitive.NewObjectID()
	mockUserID := uuid.NewString()
	mockEmail := gofakeit.Email()
	mockFirstname := gofakeit.FirstName()
	mockLastname := gofakeit.LastName()
	mockProvider := model.LocalProvider
	mockUserRole := model.UserRole
	mockPassword := gofakeit.Password(true, true, true, true, false, 10)
	mockPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
	mockUserRegistered := model.User{
		ID:        mockID,
		UserID:    mockUserID,
		Provider:  mockProvider,
		Email:     mockEmail,
		Password:  string(mockPasswordHash),
		Firstname: mockFirstname,
		Lastname:  mockLastname,
		Role:      mockUserRole,
		IsActive:  true,
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
	}

	t.Run("Get profile success", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByID").Return(&mockUserRegistered, nil)
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetProfile(mockID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockUserRegistered, *result)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Get profile fail: repo error", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByID").Return(nil, errors.New("repo error"))
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetProfile(mockID.Hex())

		assert.Nil(t, result)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Get profile fail: user not found", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByID").Return(nil, nil)
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetProfile(mockID.Hex())

		assert.Nil(t, result)
		if assert.Error(t, err) && assert.ErrorIs(t, err, err.(*fiber.Error)) {
			assert.Equal(t, fiber.StatusNotFound, err.(*fiber.Error).Code)
		}

		mockUserRepo.AssertExpectations(t)
	})
}

func TestGellAll(t *testing.T) {
	mockUsersRegistered := []model.User{
		{
			ID:        primitive.NewObjectID(),
			UserID:    uuid.NewString(),
			Provider:  model.LocalProvider,
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, true, true, true, false, 10),
			Firstname: gofakeit.FirstName(),
			Lastname:  gofakeit.LastName(),
			Role:      model.UserRole,
			IsActive:  true,
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		{
			ID:        primitive.NewObjectID(),
			UserID:    uuid.NewString(),
			Provider:  model.LocalProvider,
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, true, true, true, false, 10),
			Firstname: gofakeit.FirstName(),
			Lastname:  gofakeit.LastName(),
			Role:      model.UserRole,
			IsActive:  true,
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
	}

	t.Run("Get all success", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("All").Return(mockUsersRegistered, nil)
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetAllUser()

		assert.NoError(t, err)
		assert.Equal(t, len(mockUsersRegistered), len(result))
		assert.Equal(t, mockUsersRegistered, result)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Get all: no user", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("All").Return([]model.User{}, nil)
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetAllUser()

		assert.NoError(t, err)
		assert.Equal(t, 0, len(result))
		assert.Equal(t, []model.User{}, result)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Get all: repo error", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("All").Return(nil, errors.New("repo error"))
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetAllUser()

		assert.Error(t, err)
		assert.Nil(t, result)

		mockUserRepo.AssertExpectations(t)
	})
}
