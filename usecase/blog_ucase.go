package usecase

import (
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/logger"
	"time"

	"github.com/google/uuid"
)

type blogUsecase struct {
	blogRepo domain.BlogRepository
}

func NewBlogUsecase(br domain.BlogRepository) domain.BlogUsecase {
	return &blogUsecase{
		blogRepo: br,
	}
}

func (u *blogUsecase) CreateBlog(blog *model.Blog, createdBy model.CreatedBy) error {
	newBlog := &model.Blog{
		BlogID:    uuid.NewString(),
		Title:     blog.Title,
		Desc:      blog.Desc,
		Image:     blog.Image,
		LikesBy:   []model.Like{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: createdBy,
	}
	err := u.blogRepo.Create(newBlog)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *blogUsecase) GetAllBlogs() ([]model.Blog, error) {
	blogs, err := u.blogRepo.All()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return blogs, nil
}

func (u *blogUsecase) GetBlog(blogID string) (*model.Blog, error) {
	blog, err := u.blogRepo.GetByID(blogID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if blog == nil {
		return nil, domain.ErrBlogNotFound
	}

	return blog, nil
}

func (u *blogUsecase) UpdateBlog(blog *model.BlogUpdateRequest, updatedBy string) error {
	currentBlog, err := u.blogRepo.GetByID(blog.BlogID)
	if err != nil {
		logger.Error(err)
		return err
	}
	if currentBlog == nil {
		return domain.ErrBlogNotFound
	}

	if currentBlog.CreatedBy.UserID != updatedBy {
		return domain.ErrNoAccessResource
	}
	blog.UpdatedAt = time.Now()
	err = u.blogRepo.UpdateOne(blog)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *blogUsecase) DeleteBlog(blogID string, deletedBy string) error {
	currentBlog, err := u.blogRepo.GetByID(blogID)
	if err != nil {
		logger.Error(err)
		return err
	}
	if currentBlog == nil {
		return domain.ErrBlogNotFound
	}
	if currentBlog.CreatedBy.UserID != deletedBy {
		return domain.ErrNoAccessResource
	}
	err = u.blogRepo.Delete(blogID)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
