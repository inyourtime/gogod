package main

import (
	"context"
	"gogod/config"
	"gogod/delivery/route"
	"gogod/pkg/database"
	"gogod/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_flogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// load config
	cfg := config.LoadConfig()
	_ = logger.NewLogger(cfg)

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

	// register route
	route.SetupRoute(app)

	log.Fatal(app.Listen(":" + cfg.App.ServerPort))
}
