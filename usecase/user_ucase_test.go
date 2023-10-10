package usecase_test

import (
	"errors"
	"gogod/domain/mock"
	"gogod/model"
	"gogod/usecase"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func TestGetProfile(t *testing.T) {
	mockID := primitive.NewObjectID()
	mockEmail := gofakeit.Email()
	mockFirstname := gofakeit.FirstName()
	mockLastname := gofakeit.LastName()
	mockProvider := model.LocalProvider
	mockUserRole := model.UserRole
	mockPassword := gofakeit.Password(true, true, true, true, false, 10)
	mockPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
	mockUserRegistered := model.User{
		ID:        mockID,
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

	t.Run("Get profile fail: invalid object id", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		u := usecase.NewUserUsecase(mockUserRepo)

		result, err := u.GetProfile(gofakeit.LetterN(20))

		assert.Nil(t, result)
		if assert.Error(t, err) && assert.ErrorIs(t, err, err.(*fiber.Error)) {
			assert.Equal(t, fiber.StatusBadRequest, err.(*fiber.Error).Code)
		}

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
