package usecase

import (
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"
	"time"

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
