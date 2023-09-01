package routes

import (
	"github.blkcor.go-admin/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)
	app.Get("/auth/user", controllers.User)
}
