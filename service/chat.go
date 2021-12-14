package service

import "chat-system/domain"

func (i *Info) CreateChat(chat *domain.Chat) error {
	return i.database.Create(chat).Error
}
