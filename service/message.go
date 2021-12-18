package service

import (
	"chat-system/domain"
	"context"
	"errors"
	"fmt"
)

func (s *Store) CreateMessage(ctx context.Context, message *domain.Message) error {
	s.database.Select("id").Table("applications").Where("token = ?", message.AppToken).Scan(&message.AppID)
	s.database.Select("id").Table("chats").Where("app_id = ? AND number = ?", message.AppID, message.ChatNumber).Scan(&message.ChatID)
	if message.AppID == 0 || message.ChatID == 0 {
		return errors.New("unable to find specified app token and/or chat number")
	}
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
