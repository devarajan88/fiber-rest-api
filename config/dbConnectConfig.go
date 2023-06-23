package config

import (
	"fiber-rest-api/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Welcome to golang rest api with fiber framework")

	godotenv.Load()

	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", dbuser, dbpassword, dbhost, dbname)

	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection failed.")
	}
	DB = db
	fmt.Println("DB connected successfully.")
	autoMigrate(db)
}

func autoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(
		&model.Student{},
		&model.Department{},
		&model.CsvRecord{},
	)
	if err != nil {
		panic("DB table  migration is failed due to unexpected error.")
	}
}
