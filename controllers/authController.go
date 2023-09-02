package controllers

import (
	"github.blkcor.go-admin/database"
	"github.blkcor.go-admin/models"
	"github.blkcor.go-admin/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Register(ctx *fiber.Ctx) error {
	var data map[string]interface{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "The password entered twice is inconsistent",
		})
	}
	//注册的用户默认分配角色
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     data["email"].(string),
		RoleId:    1,
	}
	user.SetPassword(data["password"].(string))

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

func Login(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	//校验用户是否存在
	user := models.User{}
	if database.DB.Where("email", data["email"]).First(&user); user.Id == 0 {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//校验用户的密码是否正确
	if err := user.ComparePassword(data["password"]); err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "wrong email or password",
		})
	}

	//FIXME:air bug：当我们在return的时候想要把用户的信息变成json数据返回出去的时候 air启动不报错 并且提示位置有问题

	//用户登录成功
	ctx.Status(200)
	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//set cookie to the request ctx
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "LOGIN SUCCESSFULLY",
	})
}

func Logout(ctx *fiber.Ctx) error {
	//清除cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "LOGOUT SUCCESSFULLY",
	})
}

func UpdateInfo(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	cookie := ctx.Cookies("jwt")
	issuer, _ := util.ParseJWT(cookie)

	database.DB.Model(&user).Where("id", issuer).Updates(user)

	return ctx.JSON(user)
}

func UpdatePassword(ctx *fiber.Ctx) error {
	var data map[string]interface{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	cookie := ctx.Cookies("jwt")
	issuer, _ := util.ParseJWT(cookie)
	if data["password"] != data["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "The password entered twice is inconsistent",
		})
	}
	user := models.User{}
	user.SetPassword(data["password"].(string))

	database.DB.Model(&models.User{}).Where("id", issuer).Updates(user)

	return ctx.JSON(user)
}
