package service

import (
	"chat-system/domain"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
)

func (s *Store) CreateApplication(application *domain.Application) error {
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

func (s *Store) UpdateApplication(application *domain.Application, token string) error {
	return s.database.Model(&domain.Application{}).Where("token = ?", token).Update("name", application.Name).Scan(application).Error
}

func (s *Store) GetApplications() ([]domain.Application, error) {
	applications := make([]domain.Application, 0)
	err := s.database.Find(&applications).Error
	return applications, err
}

func (s *Store) GetApplicationByToken(appToken string) (*domain.Application, error) {
	application := new(domain.Application)
	err := s.database.Where("token = ?", appToken).First(&application).Error
	return application, err
}
