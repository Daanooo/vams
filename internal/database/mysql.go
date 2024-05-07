package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		return nil, err
	}

	log.Printf("Database created, running auto migrations")

	RunMigrations(db)

	log.Printf("Migrations completed")

	return db, nil
}
