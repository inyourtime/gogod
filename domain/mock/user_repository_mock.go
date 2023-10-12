package mock

import (
	"gogod/model"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) Create(user *model.User) (*model.User, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (_m *UserRepository) GetByID(_id primitive.ObjectID, withPwd bool) (*model.User, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (_m *UserRepository) GetByEmail(email string, withPwd bool) (*model.User, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (_m *UserRepository) All() ([]model.User, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), args.Error(1)
}
