package service

import (
	"chat-system/domain"
	"context"
	"errors"
	"fmt"
)

func (s *Store) CreateMessage(ctx context.Context, message *domain.Message) error {
	var appID uint
	s.database.Select("id").Table("applications").Where("token = ?", message.AppToken).Scan(&appID)
	s.database.Select("id").Table("chats").Where("app_id = ? AND number = ?", appID, message.ChatNumber).Scan(&message.ChatID)
	if appID == 0 || message.ChatID == 0 {
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
	appIDSubQuery := s.database.Select("id").Table("applications").Where("token = ?", appToken)
	chatIDSubQuery := s.database.Select("id").Table("chats").Where("app_id = (?) AND number = ?", appIDSubQuery, chatNumber)
	err := s.database.Select("*, ? AS app_token, ? AS chat_number", appToken, chatNumber).
		Where("chat_id = (?)", chatIDSubQuery).Find(&messages).Error
	return messages, err
}

func (s *Store) GetMessageByApplicationAndChatAndNumber(appToken string, number, msgNumber uint) (*domain.Message, error) {
	message := new(domain.Message)
	var appID uint
	s.database.Select("id").Table("applications").Where("token = ?", appToken).Scan(&appID)
	chatIDSubQuery := s.database.Select("id").Table("chats").Where("app_id = ? AND number = ?", appID, number)
	err := s.database.Select("*, ? AS app_token, ? AS chat_number", appToken, number).
		Where("app_id = ? AND chat_id = (?) AND number = ?", appID, chatIDSubQuery, msgNumber).First(&message).Error
	return message, err
}
