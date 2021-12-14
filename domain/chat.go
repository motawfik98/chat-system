package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID        string      `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Number    uint        `json:"number" gorm:"uniqueIndex:number_token"`
	AppToken  string      `json:"appToken" gorm:"uniqueIndex:number_token;size:36" validate:"required"`
	App       Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppToken;references:Token" validate:"required,nostructlevel"`
	CreatedAt time.Time
}

func (c *Chat) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	c.ID = u.String()
	return
}
