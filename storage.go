package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDatabase() {
	dsn := "host=localhost port=5432 user=postgres password=fakha dbname=todolist sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&Task{}, &Category{})
	db = database
	fmt.Println("Database connected ✅")
}
