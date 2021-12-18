package domain

import (
	"time"
)

const (
	MAX_MESSAGE_NUMBER = "max-message-number"
	TOTAL_MESSAGES     = "total-messages"
)

type Message struct {
	ID         string      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	AppToken   string      `json:"appToken" gorm:"uniqueIndex:application_number_chat;size:36"`
	ChatNumber uint        `json:"chatNumber" gorm:"uniqueIndex:application_number_chat"`
	Number     uint        `json:"number" gorm:"uniqueIndex:application_number_chat"`
	ChatID     string      `json:"-" gorm:"size:36"`
	Chat       Chat        `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatID;references:ID" validate:"required,nostructlevel"`
	AppID      string      `json:"-" gorm:"size:36"`
	App        Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppID;references:ID" validate:"required,nostructlevel"`
	Message    string      `json:"message" validate:"required"`
	CreatedAt  *time.Time  `json:",omitempty"`
}
