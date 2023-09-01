package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.blkcor.go-admin/util"
	"github.com/gofiber/fiber/v2"
)

func User(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	issuer, _ := util.ParseJWT(cookie)

	var user models.User
	database.DB.Where("id", issuer).Find(&user)
	if user.Id == 0 {
		ctx.Status(fiber.StatusConflict)
		return ctx.JSON(fiber.Map{
			"message": "Couldn't find the user, maybe you account has been deleted!Please associate with the admin if you have further issues!",
		})
	}
	return ctx.JSON(user)
}
