package config

import (
	"fmt"
	"toko-elektronik/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB_USER := "root"
	DB_PASS := ""
	DB_HOST := "127.0.0.1"
	DB_PORT := "3306"
	DB_NAME := "toko_elektronik"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.Item{})
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.User{})
}
