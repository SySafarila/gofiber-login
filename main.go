package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"mygo/handlers"
)

var Validate = validator.New()

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})
	app.Use(recover.New())

	app.Get("/", handlers.RootHandler)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/register", handlers.Register)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/auth/me", handlers.Me)

	err := app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
