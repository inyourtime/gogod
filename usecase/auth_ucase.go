package usecase

import (
	"gogod/domain"
	"gogod/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepo domain.AuthRepository
	userRepo domain.UserRepository
}

func NewAuthUsecase(ar domain.AuthRepository, ur domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		authRepo: ar,
		userRepo: ur,
	}
}

func (u *authUsecase) Login(req model.AuthLoginRequest) (*model.AuthLoginResponse, error) {
	currentUser, err := u.userRepo.GetByEmail(req.Email, true)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Email or Password are not correct 🥲")
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(req.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Email or Password are not correct 🥲")
	}

	token, err := u.authRepo.SignUserToken(currentUser)
	if err != nil {
		return nil, err
	}

	response := &model.AuthLoginResponse{
		Email:     currentUser.Email,
		Firstname: currentUser.Firstname,
		Lastname:  currentUser.Lastname,
		UpdatedAt: currentUser.UpdatedAt,
		Token:     token,
	}
	return response, nil
}

func (u *authUsecase) Register(req model.User) (*model.User, error) {
	// check user exist
	currentUser, err := u.userRepo.GetByEmail(req.Email, false)
	if err != nil {
		return nil, err
	}
	if currentUser != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "email already exist 😜")
	}
	// hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}
	// inst
	newUser := model.User{
		Provider:  req.Provider,
		Email:     req.Email,
		Password:  string(bytes),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
		Role:      model.UserRole,
		GoogleID:  req.GoogleID,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	response, err := u.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return response, nil
}
