package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Africa%2FCairo")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}