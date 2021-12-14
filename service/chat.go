package service

import "chat-system/domain"

func (i *Info) CreateChat(chat *domain.Chat) error {
	return i.database.Create(chat).Error
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
