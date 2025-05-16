package db

import (
	"BlogHub/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=blog_hub port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	err = DB.AutoMigrate(&models.Blog{}, &models.Comment{}, &models.User{})

	if err != nil{
		log.Fatal("AutoMigrate failed: ", err)
	}
}
