package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func New(config Config) *gorm.DB {
	// dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database: %v", config.Name, err)
		panic(err)
	}

	return db
}

func Load(db *gorm.DB, models ...interface{}) {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}
}

func Has(db *gorm.DB, model interface{}) bool {
	return db.Migrator().HasTable(model)
}

func Drop(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		if err := db.Migrator().DropTable(model); Has(db, model) && err != nil {
			log.Fatalf("Cannot drop table: %v", err)
		}
	}
}

func Seed(db *gorm.DB, models ...interface{}) {
	Drop(db, models...)
	Load(db, models...)
}
