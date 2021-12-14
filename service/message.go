package service

import "chat-system/domain"

func (i *Info) CreateMessage(message *domain.Message) error {
	return i.database.Create(message).Error
}
