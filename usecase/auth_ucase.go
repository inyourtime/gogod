package usecase

import (
	"gogod/domain"
	"gogod/model"

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
