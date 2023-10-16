package usecase_test

import (
	"errors"
	"gogod/domain"
	"gogod/domain/mock"
	"gogod/model"
	"gogod/usecase"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

var mockBlogs = []model.Blog{
	{
		BlogID: gofakeit.UUID(),
		Title:  gofakeit.BookTitle(),
		Desc:   gofakeit.Phrase(),
		Image:  gofakeit.ImageURL(5, 5),
		LikesBy: []model.Like{
			{
				CreatedBy: model.CreatedBy{
					UserID: gofakeit.UUID(),
					Name:   gofakeit.Name(),
				},
				CreatedAt: gofakeit.Date(),
			},
		},
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
		CreatedBy: model.CreatedBy{
			UserID: gofakeit.UUID(),
			Name:   gofakeit.Name(),
		},
	},
	{
		BlogID: gofakeit.UUID(),
		Title:  gofakeit.BookTitle(),
		Desc:   gofakeit.Phrase(),
		Image:  gofakeit.ImageURL(5, 5),
		LikesBy: []model.Like{
			{
				CreatedBy: model.CreatedBy{
					UserID: gofakeit.UUID(),
					Name:   gofakeit.Name(),
				},
				CreatedAt: gofakeit.Date(),
			},
		},
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
		CreatedBy: model.CreatedBy{
			UserID: gofakeit.UUID(),
			Name:   gofakeit.Name(),
		},
	},
}

func TestCreateBlog(t *testing.T) {
	t.Run("create blog success", func(t *testing.T) {
		mockBlog := mockBlogs[0]
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("Create").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.CreateBlog(&mockBlog, mockBlog.CreatedBy)
		assert.NoError(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("create blog fail", func(t *testing.T) {
		mockBlog := mockBlogs[0]
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("Create").Return(errors.New("erir"))
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.CreateBlog(&mockBlog, mockBlog.CreatedBy)
		assert.Error(t, err)
		mockBlogRepo.AssertExpectations(t)
	})
}

func TestGetAllBlogs(t *testing.T) {
	t.Run("get all blogs success", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("All").Return(mockBlogs, nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		result, err := u.GetAllBlogs()
		assert.NoError(t, err)
		assert.Equal(t, mockBlogs, result)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("get all blogs fail", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("All").Return(nil, errors.New("dkub"))
		u := usecase.NewBlogUsecase(mockBlogRepo)
		result, err := u.GetAllBlogs()
		assert.Error(t, err)
		assert.Nil(t, result)
		mockBlogRepo.AssertExpectations(t)
	})
}

func TestGetBlog(t *testing.T) {
	t.Run("get blog success", func(t *testing.T) {
		mockBlog := mockBlogs[0]
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		result, err := u.GetBlog(mockBlog.BlogID)
		assert.NoError(t, err)
		assert.Equal(t, mockBlog, *result)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("get blog fail", func(t *testing.T) {
		mockBlog := mockBlogs[0]
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, errors.New("ff"))
		u := usecase.NewBlogUsecase(mockBlogRepo)
		_, err := u.GetBlog(mockBlog.BlogID)
		assert.Error(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("get blog not found", func(t *testing.T) {
		mockBlog := mockBlogs[0]
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		_, err := u.GetBlog(mockBlog.BlogID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrBlogNotFound)
		mockBlogRepo.AssertExpectations(t)
	})
}

func TestUpdateBlog(t *testing.T) {
	mockBlog := mockBlogs[0]
	mockBlogUpdateReq := model.BlogUpdateRequest{
		BlogID:    mockBlog.BlogID,
		Title:     gofakeit.BookTitle(),
		Desc:      gofakeit.Phrase(),
		Image:     gofakeit.ImageURL(5, 5),
		UpdatedAt: gofakeit.Date(),
	}

	t.Run("update blog success", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("UpdateOne").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.UpdateBlog(&mockBlogUpdateReq, mockBlog.CreatedBy.UserID)
		assert.NoError(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("update blog fail: update repo err", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("UpdateOne").Return(errors.New("err"))
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.UpdateBlog(&mockBlogUpdateReq, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("update blog fail: get repo err", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, errors.New("err"))
		mockBlogRepo.On("UpdateOne").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.UpdateBlog(&mockBlogUpdateReq, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "UpdateOne", 0)
	})

	t.Run("update blog fail: blog not found", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, nil)
		mockBlogRepo.On("UpdateOne").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.UpdateBlog(&mockBlogUpdateReq, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrBlogNotFound)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "UpdateOne", 0)
	})

	t.Run("update blog fail: not owner", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("UpdateOne").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.UpdateBlog(&mockBlogUpdateReq, gofakeit.UUID())
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrNoAccessResource)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "UpdateOne", 0)
	})
}

func TestDeleteBlog(t *testing.T) {
	mockBlog := mockBlogs[0]

	t.Run("delete blog success", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("Delete").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.DeleteBlog(mockBlog.BlogID, mockBlog.CreatedBy.UserID)
		assert.NoError(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("delete blog fail: delete repo err", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("Delete").Return(errors.New("err"))
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.DeleteBlog(mockBlog.BlogID, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		mockBlogRepo.AssertExpectations(t)
	})

	t.Run("delete blog fail: get repo err", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, errors.New("err"))
		mockBlogRepo.On("Delete").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.DeleteBlog(mockBlog.BlogID, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "Delete", 0)
	})

	t.Run("delete blog fail: blog not found", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(nil, nil)
		mockBlogRepo.On("Delete").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.DeleteBlog(mockBlog.BlogID, mockBlog.CreatedBy.UserID)
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrBlogNotFound)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "Delete", 0)
	})

	t.Run("delete blog fail: not owner", func(t *testing.T) {
		mockBlogRepo := new(mock.BlogRepository)
		mockBlogRepo.On("GetByID").Return(&mockBlog, nil)
		mockBlogRepo.On("Delete").Return(nil)
		u := usecase.NewBlogUsecase(mockBlogRepo)
		err := u.DeleteBlog(mockBlog.BlogID, gofakeit.UUID())
		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrNoAccessResource)
		mockBlogRepo.AssertNumberOfCalls(t, "GetByID", 1)
		mockBlogRepo.AssertNumberOfCalls(t, "Delete", 0)
	})
}
