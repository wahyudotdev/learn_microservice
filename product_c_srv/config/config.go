package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Port = func() string {
	return ":" + os.Getenv("PORT")
}()

var Db = func() *gorm.DB {
	if os.Getenv("ENV") == "DEV" {
		db, err := gorm.Open(sqlite.Open("product.sqlite"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db

	} else {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/absen?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db
	}

}()

var MessageBrokerUrl = func() string {
	return os.Getenv("AMQP_SERVER_URL")
}()
