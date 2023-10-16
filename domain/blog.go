package domain

import (
	"gogod/model"
)

type BlogRepository interface {
	Create(blog *model.Blog) error
	All() ([]model.Blog, error)
	GetByID(blogID string) (*model.Blog, error)
	UpdateOne(blog *model.BlogUpdateRequest) error
	Delete(blogID string) error
}

type BlogUsecase interface {
	CreateBlog(blog *model.Blog, createdBy model.CreatedBy) error
	GetAllBlogs() ([]model.Blog, error)
	GetBlog(blogID string) (*model.Blog, error)
	UpdateBlog(blog *model.BlogUpdateRequest, updatedBy string) error
	DeleteBlog(blogID string, deletedBy string) error
}
