package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID         string `json:"-" gorm:"primaryKey;autoIncrement:false"`
	AppToken   string `json:"appToken" gorm:"uniqueIndex:application_number_chat;size:36"`
	ChatNumber uint   `json:"chatNumber" gorm:"uniqueIndex:application_number_chat"`
	Number     uint   `json:"number" gorm:"uniqueIndex:application_number_chat;default:1"`
	ChatID     string `json:"-" gorm:"size:36"`
	Chat       Chat   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatID;references:ID" validate:"required,nostructlevel"`
	Message    string `json:"message" validate:"required"`
	CreatedAt  time.Time
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	m.ID = u.String()
	tx.Select("MAX(number) + 1").Table("messages").Where("app_token = ? AND chat_number = ?", m.AppToken, m.ChatNumber).Scan(&m.Number)
	tx.Select("DISTINCT(ID)").Table("chats").Where("app_token = ? AND number = ?", m.AppToken, m.ChatNumber).Find(&m.ChatID)
	return
}
