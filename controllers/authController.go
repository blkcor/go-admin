package controllers

import (
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	user := models.User{
		Id:        0,
		FirstName: "blkcor",
		LastName:  "smart",
		Email:     "blkcor.dev@gmail.com",
		Password:  "Woshiguaiwu123",
	}
	return ctx.JSON(user)
}
