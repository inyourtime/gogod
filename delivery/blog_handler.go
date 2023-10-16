package delivery

import (
	"gogod/domain"
	"gogod/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type blogHandler struct {
	blogUsecase domain.BlogUsecase
}

func NewBlogHandler(bu domain.BlogUsecase) *blogHandler {
	return &blogHandler{
		blogUsecase: bu,
	}
}

func (h *blogHandler) CreateBlog(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	if user_id == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}
	user_fullname := c.Locals("name").(string)
	if user_fullname == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}

	blog := new(model.Blog)
	if err := c.BodyParser(blog); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(blog); err != nil {
		return FiberError(c, err)
	}

	createBy := model.CreatedBy{
		UserID: user_id,
		Name:   user_fullname,
	}

	err := h.blogUsecase.CreateBlog(blog, createBy)
	if err != nil {
		return FiberError(c, err)
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "create blog success",
	})
}

func (h *blogHandler) GetAllBlogs(c *fiber.Ctx) error {
	blogs, err := h.blogUsecase.GetAllBlogs()
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(blogs)
}

func (h *blogHandler) GetBlog(c *fiber.Ctx) error {
	blogID := c.Params("blogID")
	if blogID == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}
	blog, err := h.blogUsecase.GetBlog(blogID)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(blog)
}

func (h *blogHandler) UpdateBlog(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	if user_id == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}

	req := new(model.BlogUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, err)
	}

	err := h.blogUsecase.UpdateBlog(req, user_id)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "update blog success",
	})
}

func (h *blogHandler) DeleteBlog(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	if user_id == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}

	blogID := c.Params("blogID")
	if blogID == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}
	err := h.blogUsecase.DeleteBlog(blogID, user_id)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "delete blog success",
	})
}
