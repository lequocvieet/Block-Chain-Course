package db

import (
	"golang_service/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=test_online port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dbURL := "postgres://postgres:postgres@localhost:5432/test_online"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Exam{})
	db.AutoMigrate(&models.UserExam{})
	db.AutoMigrate(&models.User{})

	return db
}
