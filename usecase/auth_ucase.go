package usecase

import (
	"gogod/config"
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func (u *authUsecase) Login(req *model.AuthLoginRequest) (*model.AuthLoginResponse, error) {
	currentUser, err := u.userRepo.GetByEmail(req.Email, true)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if currentUser == nil {
		return nil, domain.ErrEmailPwdIncorrect
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(req.Password))
	if err != nil {
		return nil, domain.ErrEmailPwdIncorrect
	}

	token, err := u.authRepo.SignUserToken(currentUser)
	if err != nil {
		logger.Error(err)
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

func (u *authUsecase) Register(req *model.User) (*model.User, error) {
	// check user exist
	currentUser, err := u.userRepo.GetByEmail(req.Email, false)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if currentUser != nil {
		return nil, domain.ErrEmailExist
	}
	// hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}
	// inst
	newUser := &model.User{
		UserID:    uuid.NewString(),
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
		logger.Error(err)
		return nil, err
	}
	return response, nil
}

func (u *authUsecase) Refresh(refreshToken string) (*model.Token, error) {
	claim := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.ENV.Jwt.RefreshSecret), nil
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	userID := claim["user_id"].(string)
	user, err := u.userRepo.GetByID(userID, false)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	token, err := u.authRepo.SignUserToken(user)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return token, nil
}

func (u *authUsecase) Google(info model.GoogleInfo) (*model.AuthLoginResponse, error) {
	currentUser, err := u.userRepo.GetByEmail(info.Email, false)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if currentUser == nil {
		newUser := &model.User{
			UserID:    uuid.NewString(),
			Provider:  model.GoogleProvider,
			Email:     info.Email,
			Password:  "",
			Firstname: info.GivenName,
			Lastname:  info.FamilyName,
			Avatar:    info.Picture,
			Role:      model.UserRole,
			GoogleID:  info.ID,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		currentUser, err = u.userRepo.Create(newUser)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	} else {
		if currentUser.GoogleID != info.ID {
			return nil, domain.ErrEmailExist
		}
	}

	token, err := u.authRepo.SignUserToken(currentUser)
	if err != nil {
		logger.Error(err)
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
