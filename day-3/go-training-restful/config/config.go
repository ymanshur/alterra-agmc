package config

import (
	"fmt"

	"go-training-restful/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_USER": "root",
		"DB_PASS": "",
		"DB_PORT": "3306",
		"DB_HOST": "127.0.0.1",
		"DB_NAME": "agmc",
	}
	// dsn
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USER"],
		config["DB_PASS"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	// InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
