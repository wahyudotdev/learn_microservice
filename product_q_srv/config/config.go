package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db = func() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("ENV") == "DEV" {
		db, err := gorm.Open(sqlite.Open("product.sqlite"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db

	} else {
		user := os.Getenv("MYSQL_USER")
		pass := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/app?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db
	}

}()

var Port = func() string {
	return ":" + os.Getenv("PORT")
}()
