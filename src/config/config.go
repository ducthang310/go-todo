package config

import (
	"fmt"
	"github.com/ducthang310/go-todo/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	fmt.Println("--calling ConnectDB")
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("Error while migrating")
	}
	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	fmt.Println("--calling DisconnectDB")
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
