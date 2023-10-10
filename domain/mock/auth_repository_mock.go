package mock

import (
	"gogod/model"

	"github.com/stretchr/testify/mock"
)

type AuthRepository struct {
	mock.Mock
}

func (_m *AuthRepository) SignUserToken(user *model.User) (*model.Token, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Token), args.Error(1)
}
