package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(ctx *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Preload("Permissions").Find(&roles)
	return ctx.JSON(roles)
}

func CreateRole(ctx *fiber.Ctx) error {
	var roleDto fiber.Map
	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	database.DB.Create(&role)
	// 设置角色与权限之间的关联关系
	database.DB.Model(&role).Association("Permissions").Append(permissions)
	return ctx.JSON(role)
}

func GetRole(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Preload("Permissions").Find(&role)

	return ctx.JSON(role)
}

func UpdateRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var roleDto fiber.Map
	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result []models.Permission
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Permissions: permissions,
		Name:        roleDto["name"].(string),
	}

	database.DB.Model(&role).Updates(role)

	return ctx.JSON(role)
}

func DeleteRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Delete(&role)

	return ctx.JSON(fiber.Map{
		"message": "DELETE SUCCESSFULLY",
	})
}
