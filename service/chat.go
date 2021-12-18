package service

import (
	"chat-system/domain"
	"context"
)

func (s *Store) CreateChat(ctx context.Context, chat *domain.Chat) error {
	n, err := s.redis.HIncrBy(ctx, chat.AppToken, domain.MAX_CHAT_NUMBER, 1).Uint64()
	if err != nil {
		return err
	}
	chat.Number = uint(n)
	return nil
}

func (s *Store) GetChatsByApplication(appToken string) ([]domain.Chat, error) {
	chats := make([]domain.Chat, 0)
	err := s.database.Where("app_token = ?", appToken).Find(&chats).Error
	return chats, err
}

func (s *Store) GetChatByApplicationAndNumber(appToken string, number uint) (*domain.Chat, error) {
	chat := new(domain.Chat)
	err := s.database.Where("app_token = ? AND number = ?", appToken, number).First(&chat).Error
	return chat, err
}
