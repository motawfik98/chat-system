package domain

import (
	"time"
)

const (
	MAX_CHAT_NUMBER = "max-chat-number" // hold the maximum chat number per application
	TOTAL_CHATS     = "total-chats"     // hold the total chats number per application
)

type Chat struct {
	ID            string      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	AppToken      string      `json:"appToken" gorm:"uniqueIndex:number_token;size:36"`
	Number        uint        `json:"number" gorm:"uniqueIndex:number_token"`
	AppID         string      `json:"-" gorm:"size:36"`
	App           Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppID;references:ID" validate:"required,nostructlevel"`
	Title         string      `json:"title" validate:"required"`
	MessagesCount uint        `json:"messagesCount"`
	CreatedAt     *time.Time  `json:",omitempty"`
}
