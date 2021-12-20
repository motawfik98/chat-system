package domain

import (
	"time"
)

type Application struct {
	ID         uint       `json:"-" gorm:"primaryKey"`
	Name       string     `json:"name" validate:"required"`
	Token      string     `json:"token" param:"token" gorm:"uniqueIndex;size:36"`
	ChatsCount uint       `json:"chatsCount"`
	CreatedAt  *time.Time `json:",omitempty"`
}
