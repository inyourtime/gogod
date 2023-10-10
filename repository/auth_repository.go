package repository

import (
	"gogod/domain"
	"gogod/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	col *mongo.Collection
}

func NewAuthRepository(c *mongo.Collection) domain.AuthRepository {
	return &authRepository{
		col: c,
	}
}

func (r *authRepository) SignUserToken(user *model.User) (*model.Token, error) {
	return nil, nil
}
