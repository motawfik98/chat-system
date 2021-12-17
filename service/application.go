package service

import (
	"chat-system/domain"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
)

func (i *Info) CreateApplication(application *domain.Application) error {
	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	random := u.String()
	h := md5.New()
	_, err = io.WriteString(h, random)
	if err != nil {
		return err
	}
	application.Token = fmt.Sprintf("%x", h.Sum(nil))
	return nil
}

func (i *Info) UpdateApplication(application *domain.Application, token string) error {
	return i.database.Model(&domain.Application{}).Where("token = ?", token).Update("name", application.Name).Scan(application).Error
}

func (i *Info) GetApplications() ([]domain.Application, error) {
	applications := make([]domain.Application, 0)
	err := i.database.Find(&applications).Error
	return applications, err
}

func (i *Info) GetApplicationByToken(appToken string) (*domain.Application, error) {
	application := new(domain.Application)
	err := i.database.Where("token = ?", appToken).First(&application).Error
	return application, err
}
