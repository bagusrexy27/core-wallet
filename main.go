package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/core-wallet/app/config"
	"github.com/core-wallet/app/router"
	"github.com/core-wallet/app/utils"
)

func main() {
	config.LoadEnv()
	config.LoadConfig()
	err := config.InitDB()
	if err != nil {
		utils.ErrorLog("failed to connect database", err, true)
	}
	fmt.Println("Database connection success!")

	if err := config.InitRedis(); err != nil {
		utils.ErrorLog("failed to connect redis", err, true)
	}
	fmt.Println("Redis connection success!")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to IPROC API"))
	})
	router.SetupRoutes(app)

	app.Listen(":" + config.AppConfig.AppPort)
}
