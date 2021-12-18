package domain

import (
	"time"
)

const (
	MAX_MESSAGE_NUMBER = "max-message-number"
	TOTAL_MESSAGES     = "total-messages"
)

type Message struct {
	ID         uint        `json:"-" gorm:"primaryKey"`
	AppID      uint        `json:"-"`
	App        Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppID;references:ID" validate:"required,nostructlevel"`
	ChatID     uint        `json:"-" gorm:"uniqueIndex:appID_chatNo_messageNo"`
	Chat       Chat        `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatID;references:ID" validate:"required,nostructlevel"`
	Number     uint        `json:"number" gorm:"uniqueIndex:appID_chatNo_messageNo"`
	Message    string      `json:"message" validate:"required"`
	AppToken   string      `json:"appToken" gorm:"-"`
	ChatNumber uint        `json:"chatNumber" gorm:"-"`
	CreatedAt  *time.Time  `json:",omitempty"`
}
