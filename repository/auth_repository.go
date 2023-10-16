package repository

import (
	"gogod/config"
	"gogod/domain"
	"gogod/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	client *mongo.Client
}

func NewAuthRepository(c *mongo.Client) domain.AuthRepository {
	return &authRepository{
		client: c,
	}
}

func (r *authRepository) SignUserToken(user *model.User) (*model.Token, error) {
	claims := model.UserClaims{
		ID:    user.UserID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
		},
	}
	// access token
	acstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ssA, err := acstoken.SignedString([]byte(config.ENV.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	// refresh token
	claims.RegisteredClaims.ID = uuid.NewString()
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour))
	claims.RegisteredClaims.Issuer = "refresh_token"
	claims.RegisteredClaims.Subject = "users_refresh_token"
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ssR, err := refToken.SignedString([]byte(config.ENV.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	res := &model.Token{
		AccessToken:  ssA,
		RefreshToken: ssR,
	}
	return res, nil
}
