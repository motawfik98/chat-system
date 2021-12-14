package service

import "gorm.io/gorm"

type Info struct {
	database *gorm.DB
}

func NewInfoService(db *gorm.DB) *Info {
	return &Info{database: db}
}