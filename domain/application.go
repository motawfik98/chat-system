package domain

import (
	"time"
)

type Application struct {
	ID         string     `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Name       string     `json:"name" validate:"required"`
	Token      string     `json:"token" gorm:"uniqueIndex;size:36"`
	ChatsCount uint       `json:"chatsCount"`
	CreatedAt  *time.Time `json:",omitempty"`
}
