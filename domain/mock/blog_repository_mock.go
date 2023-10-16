package mock

import (
	"gogod/model"

	"github.com/stretchr/testify/mock"
)

type BlogRepository struct {
	mock.Mock
}

func (_m *BlogRepository) Create(blog *model.Blog) error {
	args := _m.Called()
	return args.Error(0)
}

func (_m *BlogRepository) All() ([]model.Blog, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Blog), args.Error(1)
}

func (_m *BlogRepository) GetByID(blogID string) (*model.Blog, error) {
	args := _m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Blog), args.Error(1)
}

func (_m *BlogRepository) UpdateOne(blog *model.BlogUpdateRequest) error {
	args := _m.Called()
	return args.Error(0)
}

func (_m *BlogRepository) Delete(blogID string) error {
	args := _m.Called()
	return args.Error(0)
}
