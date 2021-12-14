package service

import "chat-system/domain"

func (i *Info) CreateApplication(application *domain.Application) error {
	return i.database.Create(application).Error
}
