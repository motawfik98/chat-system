package service

import "chat-system/domain"

func (i *Info) CreateMessage(message *domain.Message) error {
	return i.database.Create(message).Error
}

func (i *Info) GetMessagesByApplicationAndChat(appToken, chatNumber string) ([]domain.Message, error) {
	messages := make([]domain.Message, 0)
	err := i.database.Where("app_token = ? AND chat_number = ?", appToken, chatNumber).Find(&messages).Error
	return messages, err
}

func (i *Info) GetMessageByApplicationAndChatAndNumber(appToken string, number, msgNumber uint) (*domain.Message, error) {
	message := new(domain.Message)
	err := i.database.Where("app_token = ? AND chat_number = ? AND number = ?", appToken, number, msgNumber).First(&message).Error
	return message, err
}
