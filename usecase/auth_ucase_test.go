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

func TestLogin(t *testing.T) {
	mockEmail := gofakeit.Email()
	mockPassword := gofakeit.Password(true, true, true, true, false, 10)
	mockPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
	mockUser := model.User{
		ID:        primitive.NewObjectID(),
		Provider:  model.LocalProvider,
		Email:     mockEmail,
		Password:  string(mockPasswordHash),
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Role:      model.UserRole,
		IsActive:  true,
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
	}

	t.Run("Login success", func(t *testing.T) {
		mockAuthRepo := new(mock.AuthRepository)
		mockUserRepo := new(mock.UserRepository)
		mockAuthRepo.On("SignUserToken").Return(&model.Token{AccessToken: gofakeit.LetterN(20), RefreshToken: gofakeit.LetterN(20)}, nil)
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

		result, err := u.Login(model.AuthLoginRequest{Email: mockEmail, Password: gofakeit.Password(true, true, true, true, false, 10)})

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

func TestRegister(t *testing.T) {
	mockEmail := gofakeit.Email()
	mockFirstname := gofakeit.FirstName()
	mockLastname := gofakeit.LastName()
	mockProvider := model.LocalProvider
	mockUserRole := model.UserRole
	mockPassword := gofakeit.Password(true, true, true, true, false, 10)
	mockPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
	mockUserRegistered := model.User{
		ID:        primitive.NewObjectID(),
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

	mockAuthRepo := new(mock.AuthRepository)

	t.Run("Register success", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, nil)
		mockUserRepo.On("Create").Return(&mockUserRegistered, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Register(model.User{Provider: mockProvider, Email: mockEmail, Password: mockEmail, Firstname: mockFirstname, Lastname: mockLastname})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockUserRegistered, *result)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Register fail: get user repo error", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, errors.New("repo error"))
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Register(model.User{Provider: mockProvider, Email: mockEmail, Password: mockEmail, Firstname: mockFirstname, Lastname: mockLastname})

		assert.Nil(t, result)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Register fail: email exist", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(&mockUserRegistered, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Register(model.User{Provider: mockProvider, Email: mockEmail, Password: mockEmail, Firstname: mockFirstname, Lastname: mockLastname})

		assert.Nil(t, result)
		if assert.Error(t, err) && assert.ErrorIs(t, err, err.(*fiber.Error)) {
			assert.Equal(t, fiber.StatusUnprocessableEntity, err.(*fiber.Error).Code)
		}

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Register fail: gen password error", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, nil)
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Register(model.User{Provider: mockProvider, Email: mockEmail, Password: gofakeit.Password(true, true, true, true, false, 80), Firstname: mockFirstname, Lastname: mockLastname})

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.ErrorIs(t, err, bcrypt.ErrPasswordTooLong)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Register fail: create repo error", func(t *testing.T) {
		mockUserRepo := new(mock.UserRepository)
		mockUserRepo.On("GetByEmail").Return(nil, nil)
		mockUserRepo.On("Create").Return(nil, errors.New("repo error"))
		u := usecase.NewAuthUsecase(mockAuthRepo, mockUserRepo)

		result, err := u.Register(model.User{Provider: mockProvider, Email: mockEmail, Password: mockEmail, Firstname: mockFirstname, Lastname: mockLastname})

		assert.Nil(t, result)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
