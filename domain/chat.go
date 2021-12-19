package domain

import (
	"time"
)

const (
	MAX_CHAT_NUMBER = "max-chat-number" // hold the maximum chat number per application
	TOTAL_CHATS     = "total-chats"     // hold the total chats number per application
)

type Chat struct {
	ID            uint        `json:"-" gorm:"primaryKey"`
	AppID         uint        `json:"-" gorm:"uniqueIndex:appID_chatNo"`
	Number        uint        `json:"number" gorm:"uniqueIndex:appID_chatNo"`
	App           Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppID;references:ID" validate:"required,nostructlevel"`
	AppToken      string      `json:"appToken,omitempty" gorm:"-:migration;<-:false" param:"token"`
	Title         string      `json:"title" validate:"required"`
	MessagesCount uint        `json:"messagesCount"`
	CreatedAt     *time.Time  `json:",omitempty"`
}
