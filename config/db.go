package config

import (
	"final-project-mygram/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "rajawali02"
	port     = "5432"
	dbname   = "mygram_app"
	db       *gorm.DB
	err 	error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s  sslmode=disable", host, user, password, port, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	defer fmt.Println("Successfully connected to database!")

	// db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
	db.Debug().AutoMigrate(model.User{}, model.Photo{})
}

func GetDB() *gorm.DB {
	return db
}