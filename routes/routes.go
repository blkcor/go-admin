package routes

import (
	"github.blkcor.go-admin/controllers"
	"github.blkcor.go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)
	//the routes below need to be authenticated!
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
	app.Post("api/updateInfo", controllers.UpdateInfo)
	app.Post("api/updatePassword", controllers.UpdatePassword)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/role/:id", controllers.GetRole)
	app.Put("/api/role/:id", controllers.UpdateRole)
	app.Delete("/api/role/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)
	app.Post("/api/permissions", controllers.CreatePermission)
	app.Get("/api/permission/:id", controllers.GetPermission)
	app.Put("/api/permission/:id", controllers.UpdatePermission)
	app.Delete("/api/permission/:id", controllers.DeletePermission)
}
