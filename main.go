package main

import (
	"fmt"
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = database.GetDatabaseConnection()
	app := fiber.New()
	routes.Setup(app)
	err := app.Listen(":8888")
	if err != nil {
		fmt.Printf("Fail to start the serve: %s", err.Error())
		return
	}
}
