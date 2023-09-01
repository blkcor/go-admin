package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllUsers(ctx *fiber.Ctx) error {
	var users []models.User
	database.DB.Preload("Role").Find(&users)
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

func GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Preload("Role").Find(&user)

	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(&user)

	return ctx.JSON(fiber.Map{
		"message": "DELETE SUCCESSFULLY",
	})
}
