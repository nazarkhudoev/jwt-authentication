package database

import (
	"jwt-authentication/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB//We created global var in order to use for saving data in our controllers

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:my-password@/toydam"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})

}