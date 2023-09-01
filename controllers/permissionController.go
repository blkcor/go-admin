package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllPermissions(ctx *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Preload("Permissions").Find(&permissions)
	return ctx.JSON(permissions)
}

func CreatePermission(ctx *fiber.Ctx) error {
	var permission models.Permission
	if err := ctx.BodyParser(&permission); err != nil {
		return err
	}

	database.DB.Create(&permission)
	return ctx.JSON(permission)
}

func GetPermission(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	permission := models.Permission{
		Id: uint(id),
	}
	database.DB.Find(&permission)

	return ctx.JSON(permission)
}

func UpdatePermission(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	permission := &models.Permission{
		Id: uint(id),
	}
	if err := ctx.BodyParser(&permission); err != nil {
		return err
	}

	database.DB.Model(&permission).Updates(permission)
	return ctx.JSON(permission)
}

func DeletePermission(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	permission := models.Permission{
		Id: uint(id),
	}
	database.DB.Delete(&permission)

	return ctx.JSON(fiber.Map{
		"message": "DELETE SUCCESSFULLY",
	})
}
