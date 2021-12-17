package service

import "chat-system/domain"

func (s *Store) CreateMessage(message *domain.Message) error {
	return s.database.Create(message).Error
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
