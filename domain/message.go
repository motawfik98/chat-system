package domain

import (
	"time"
)

const (
	MAX_MESSAGE_NUMBER = "max-message-number"
	TOTAL_MESSAGES     = "total-messages"
)

type Message struct {
	ID         uint       `json:"-" gorm:"primaryKey"`
	ChatID     uint       `json:"-" gorm:"uniqueIndex:chatID_messageNo"`
	Chat       Chat       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatID;references:ID" validate:"required,nostructlevel"`
	Number     uint       `json:"number" param:"msg" gorm:"uniqueIndex:chatID_messageNo"`
	Message    string     `json:"message" validate:"required"`
	AppToken   string     `json:"appToken,omitempty" gorm:"-:migration;<-:false" param:"token"`
	ChatNumber uint       `json:"chatNumber,omitempty" gorm:"-:migration;<-:false" param:"number"`
	CreatedAt  *time.Time `json:",omitempty"`
	UpdatedAt  *time.Time `json:",omitempty"`
}
