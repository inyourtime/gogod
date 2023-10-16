package route

import (
	"gogod/delivery"
	"gogod/delivery/middleware"
	"gogod/pkg/database"
	"gogod/repository"
	"gogod/usecase"

	"github.com/gofiber/fiber/v2"
)

func blogRoute(router fiber.Router) {
	blog := router.Group("/blog")

	blogRepo := repository.NewBlogRepository(database.MC)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogHandler := delivery.NewBlogHandler(blogUsecase)

	{
		blog.Post("/", middleware.JwtGuard(), blogHandler.CreateBlog)
		blog.Get("/", middleware.JwtGuard(), blogHandler.GetAllBlogs)
		blog.Get("/:blogID", middleware.JwtGuard(), blogHandler.GetBlog)
		blog.Put("/", middleware.JwtGuard(), blogHandler.UpdateBlog)
		blog.Delete("/:blogID", middleware.JwtGuard(), blogHandler.DeleteBlog)
	}
}
