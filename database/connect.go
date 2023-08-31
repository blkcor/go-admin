package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:12345678@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database!" + err.Error())
	}
	return db
}
