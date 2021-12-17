package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Store struct {
	database *gorm.DB
	redis    *redis.Client
}

func NewInfoService(db *gorm.DB, redis *redis.Client) *Store {
	return &Store{database: db, redis: redis}
}
