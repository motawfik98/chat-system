package domain

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type Application struct {
	ID        string `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Name      string `json:"name" validate:"required"`
	Token     string `json:"token" gorm:"uniqueIndex;size:36"`
	CreatedAt time.Time
}

func (a *Application) BeforeCreate(tx *gorm.DB) (err error) {
	u, e := uuid.NewRandom()
	if e != nil {
		err = e
	}
	a.ID = u.String()
	h := md5.New()
	io.WriteString(h, a.ID)
	a.Token = fmt.Sprintf("%x", h.Sum(nil))
	return
}
