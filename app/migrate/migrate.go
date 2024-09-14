package main

import (
	"app/config"
	"app/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", config.ENV.DB_USER, config.ENV.DB_PASSWORD, config.ENV.DB_HOST, config.ENV.DB_PORT, config.ENV.DB_DATABASE, "%2F")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Blog{}, &models.Comment{})
	log.Println("Migration Completed")
}
