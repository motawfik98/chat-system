package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Application struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Name      string `json:"name" validate:"required"`
	Token     string `json:"token" gorm:"uniqueIndex;size:36"`
	CreatedAt time.Time
}

func (a *Application) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	a.Token = u.String()
	return
}
