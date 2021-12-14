package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID          string      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Number      uint        `json:"number" gorm:"uniqueIndex:application_number_chat;default:1"`
	AppToken    string      `json:"appToken" gorm:"uniqueIndex:application_number_chat"`
	ChatNumber  string      `json:"chatNumber" gorm:"uniqueIndex:application_number_chat"`
	Application Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppToken;references:Token" validate:"required,nostructlevel"`
	Chat        Chat        `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatNumber;references:Number" validate:"required,nostructlevel"`
	CreatedAt   time.Time
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	m.ID = u.String()
	var max uint
	tx.Select("MAX(number) + 1").Table("messages").Where("app_token = ? AND chat_number = ?", m.AppToken, m.ChatNumber).Scan(&max)
	m.Number = max
	return
}
