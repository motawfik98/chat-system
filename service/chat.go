package service

import (
	"chat-system/domain"
	"context"
)

func (i *Info) CreateChat(chat *domain.Chat) error {
	n, err := i.redis.HIncrBy(context.Background(), chat.AppToken, "number-of-chats", 1).Uint64()
	if err != nil {
		return err
	}
	chat.Number = uint(n)
	return nil
}

func (i *Info) GetChatsByApplication(appToken string) ([]domain.Chat, error) {
	chats := make([]domain.Chat, 0)
	err := i.database.Where("app_token = ?", appToken).Find(&chats).Error
	return chats, err
}

func (i *Info) GetChatByApplicationAndNumber(appToken string, number uint) (*domain.Chat, error) {
	chat := new(domain.Chat)
	err := i.database.Where("app_token = ? AND number = ?", appToken, number).First(&chat).Error
	return chat, err
}
