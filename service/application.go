package service

import "chat-system/domain"

func (i *Info) CreateApplication(application *domain.Application) error {
	return i.database.Create(application).Error
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
