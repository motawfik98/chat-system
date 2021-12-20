package service

import (
	"bytes"
	"chat-system/domain"
	"chat-system/domain/elasticdomain"
	"context"
	"encoding/json"
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

func (s *Store) UpdateMessage(message *domain.Message) error {
	var appID uint
	s.database.Select("id").Table("applications").Where("token = ?", message.AppToken).Scan(&appID)
	chatIDSubQuery := s.database.Select("id").Table("chats").Where("app_id = ? AND number = ?", appID, message.ChatNumber)
	s.database.Select("id").Table("messages").Where("chat_id = (?) AND number = ?", chatIDSubQuery, message.Number).Scan(&message.ID)
	if message.ID == 0 {
		return errors.New("unable to find specified app token, chat number, and message number combination")
	}
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
		Where("chat_id = (?) AND number = ?", chatIDSubQuery, msgNumber).First(&message).Error
	return message, err
}

func (s *Store) SearchMessages(ctx context.Context, appToken string, number uint, message string) ([]domain.Message, error) {
	var id int64
	appIDSubQuery := s.database.Select("id").Table("applications").Where("token = ?", appToken)
	s.database.Select("id").Table("chats").Where("app_id = (?) AND number = ?", appIDSubQuery, number).Count(&id)
	if id == 0 {
		return nil, errors.New("unable to find specified app token and chat number combination")
	}
	var buf bytes.Buffer
	query := elasticdomain.CreateQuery(message)
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	// Perform the search request.
	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex("messages"),
		s.es.Search.WithBody(&buf),
		s.es.Search.WithTrackTotalHits(true),
		s.es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}
		// Print the response status and error information.
		return nil, errors.New(fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		))

	}

	response := new(elasticdomain.SuccessResponse)

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.GetMessages(appToken, number), nil
}
