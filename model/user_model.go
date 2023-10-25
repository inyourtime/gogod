package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type providerType string
type roleType string

const (
	LocalProvider  providerType = "local"
	GoogleProvider providerType = "google"
)

const (
	AdminRole roleType = "admin"
	UserRole  roleType = "user"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId,omitempty"`
	Provider  providerType       `json:"provider,omitempty" bson:"provider"`
	Email     string             `json:"email,omitempty" bson:"email" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname" validate:"required"`
	Avatar    string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Role      roleType           `json:"role,omitempty" bson:"role"`
	GoogleID  string             `json:"googleId,omitempty" bson:"googleId,omitempty"`
	IsActive  bool               `json:"isActive,omitempty" bson:"isActive"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func NewUser(user User) User {
	return User{
		Provider:  user.Provider,
		Email:     user.Email,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Avatar:    user.Avatar,
		GoogleID:  user.GoogleID,
		Role:      UserRole,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type AuthLoginResponse struct {
	Email     string    `json:"email,omitempty"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Token     *Token    `json:"token"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

type UpdateUserRequest struct {
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Avatar    string `json:"avatar,omitempty" bson:"avatar,omitempty"`
}

type GoogleInfo struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}
