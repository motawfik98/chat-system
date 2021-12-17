package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Info struct {
	database *gorm.DB
	redis    *redis.Client
}

func NewInfoService(db *gorm.DB, redis *redis.Client) *Info {
	return &Info{database: db, redis: redis}
}
