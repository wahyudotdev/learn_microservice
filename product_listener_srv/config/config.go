package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

var Db = func() *gorm.DB {
	if os.Getenv("ENV") == "DEV" {
		db, err := gorm.Open(sqlite.Open("../product_q_srv/product.sqlite"), &gorm.Config{})
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

var MessageBrokerUrl = func() string {
	return os.Getenv("AMQP_SERVER_URL")
}()

var Timeout = func() time.Duration {
	return 90 * time.Second
}()
