package database

import (
	"fmt"
	"github.blkcor.go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDatabaseConnection() {
	db, err := gorm.Open(mysql.Open("root:12345678@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database!" + err.Error())
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Something went wrong when connecting to the database: ", err.Error())
	}
	DB = db
}
