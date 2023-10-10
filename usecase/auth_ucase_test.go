package usecase_test

import (
	"errors"
	"gogod/domain/mock"
	"gogod/model"
	"gogod/usecase"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	mockEmail := faker.Email()
	mockPassword := faker.Password()
	mockPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
	mockUser := model.User{
		ID:        primitive.NewObjectID(),
		Provider:  model.LocalProvider,
		Email:     mockEmail,
		Password:  string(mockPasswordHash),
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		Role:      model.UserRole,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Login success", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockAuthRepo.On("SignUserToken").Return(&model.Token{AccessToken: faker.Jwt(), RefreshToken: faker.Jwt()}, nil)
		mockUserRepo.On("GetByEmail").Return(&mockUser, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: mockPassword})

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.NotNil(t, result)
		assert.IsType(t, &model.AuthLoginResponse{}, result)

		mockAuthRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Login fail: repo error", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, errors.New("repo error"))
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: mockPassword})

		assert.Error(t, err)
		assert.Nil(t, result)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Login fail: user not found", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: mockPassword})

		assert.Nil(t, result)
		if assert.Error(t, err) && assert.ErrorIs(t, err, err.(*fiber.Error)) {
			assert.Equal(t, fiber.StatusUnauthorized, err.(*fiber.Error).Code)
		}

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Login fail: password incorrect", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(&mockUser, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: faker.Password()})

		assert.Nil(t, result)
		if assert.Error(t, err) && assert.ErrorIs(t, err, err.(*fiber.Error)) {
			assert.Equal(t, fiber.StatusUnauthorized, err.(*fiber.Error).Code)
		}

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Login fail: sign token error", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockAuthRepo.On("SignUserToken").Return(nil, errors.New("sign token error"))
		mockUserRepo.On("GetByEmail").Return(&mockUser, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: mockPassword})

		assert.Nil(t, result)
		assert.Error(t, err)

		mockAuthRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}
