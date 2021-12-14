package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID        string      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	AppToken  string      `json:"appToken" gorm:"uniqueIndex:number_token;size:36"`
	Number    uint        `json:"number" gorm:"uniqueIndex:number_token;default:1"`
	AppID     string      `json:"-" gorm:"size:36"`
	App       Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppID;references:ID" validate:"required,nostructlevel"`
	Title     string      `json:"title" validate:"required"`
	CreatedAt time.Time
}

func (c *Chat) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	c.ID = u.String()
	tx.Select("MAX(number) + 1").Table("chats").Where("app_token = ?", c.AppToken).Scan(&c.Number)
	tx.Select("DISTINCT(ID)").Table("applications").Where("token = ?", c.AppToken).Scan(&c.AppID)
	return
}
