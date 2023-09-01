package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}

func User(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("blkcor"), nil
	})
	if err != nil || !token.Valid {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims
	userId, _ := claims.GetIssuer()

	var user models.User
	database.DB.Where("id", userId).Find(&user)
	if user.Id == 0 {
		ctx.Status(fiber.StatusConflict)
		return ctx.JSON(fiber.Map{
			"message": "Couldn't find the user, maybe you account has been deleted!Please associate with the admin if you have further issues!",
		})
	}
	return ctx.JSON(user)
}
