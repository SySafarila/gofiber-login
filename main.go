package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mygo/database"
	"mygo/handlers"
)

func main() {
	dsn := "host=localhost user=postgres password='' dbname=slime port=5432"
	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic(errDb)
	}

	//errMigrate := db.AutoMigrate(&models.User{})
	//if errMigrate != nil {
	//	panic(errMigrate)
	//}
	// assign db
	database.DB = db

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
