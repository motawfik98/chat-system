package service

import (
	"chat-system/domain"
	"context"
	"fmt"
)

func (s *Store) CreateMessage(ctx context.Context, message *domain.Message) error {
	n, err := s.redis.HIncrBy(ctx, fmt.Sprintf("%s-%d", message.AppToken, message.ChatNumber), domain.MAX_MESSAGE_NUMBER, 1).Uint64()
	if err != nil {
		return err
	}
	message.Number = uint(n)
	return nil
}

func (s *Store) GetMessagesByApplicationAndChat(appToken, chatNumber string) ([]domain.Message, error) {
	messages := make([]domain.Message, 0)
	err := s.database.Where("app_token = ? AND chat_number = ?", appToken, chatNumber).Find(&messages).Error
	return messages, err
}

func (s *Store) GetMessageByApplicationAndChatAndNumber(appToken string, number, msgNumber uint) (*domain.Message, error) {
	message := new(domain.Message)
	err := s.database.Where("app_token = ? AND chat_number = ? AND number = ?", appToken, number, msgNumber).First(&message).Error
	return message, err
}
