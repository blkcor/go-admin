package main

import (
	"fmt"
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.GetDatabaseConnection()
	fmt.Println(database.DB)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	err := app.Listen(":8888")
	if err != nil {
		fmt.Printf("Fail to start the serve: %s", err.Error())
		return
	}
}
