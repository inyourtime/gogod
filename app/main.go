package main

import (
	"context"
	"gogod/config"
	"gogod/delivery/middleware"
	"gogod/delivery/route"
	"gogod/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_flogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// load config
	cfg := config.LoadConfig()
	_ = config.LoadGoogleConfig(cfg)

	// connect database
	mc := database.MongoDBConnect(cfg)
	defer mc.Disconnect(context.TODO())

	app := fiber.New(fiber.Config{
		RequestMethods: fiber.DefaultMethods,
		ErrorHandler:   fiber.DefaultErrorHandler,
	})

	// middleware here
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: cors.ConfigDefault.AllowMethods,
	}))
	app.Use(_flogger.New(_flogger.Config{
		TimeZone: "Asia/Bangkok",
	}))
	app.Use(middleware.Recover())

	// register route
	route.SetupRoute(app)

	log.Fatal(app.Listen(":" + cfg.App.ServerPort))
}
