package config

import (
	"fmt"
	"go-restful/model"

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

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migration
	// InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}
