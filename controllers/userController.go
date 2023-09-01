package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(ctx *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return ctx.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	//TODO:这里好像是给一个默认的密码
	user.SetPassword(user.Password)

	//check if the email is in use
	var userTmp models.User
	if database.DB.Where("email = ?", user.Email).First(&userTmp); userTmp.Id != 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "The email is already in use",
		})
	}

	database.DB.Create(&user)
	return ctx.JSON(user)
}
