package db

import (
	"log"

	"github.com/WellDoneIT90/Golang_project/30postgresql/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://welldoneit:welldoneit@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Book{})

	return db
}
