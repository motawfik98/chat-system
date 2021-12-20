package service

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Store struct {
	database *gorm.DB
	redis    *redis.Client
	es       *elasticsearch.Client
}

func NewInfoService(db *gorm.DB, redis *redis.Client, es *elasticsearch.Client) *Store {
	return &Store{database: db, redis: redis, es: es}
}
