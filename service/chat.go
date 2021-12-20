package service

import (
	"chat-system/domain"
	"context"
	"errors"
)

func (s *Store) CreateChat(ctx context.Context, chat *domain.Chat) error {
	s.database.Select("id").Table("applications").Where("token = ?", chat.AppToken).Scan(&chat.AppID)
	if chat.AppID == 0 {
		return errors.New("unable to find specified app token")
	}
	n, err := s.redis.HIncrBy(ctx, chat.AppToken, domain.MAX_CHAT_NUMBER, 1).Uint64()
	if err != nil {
		return err
	}
	chat.Number = uint(n)
	return nil
}

func (s *Store) UpdateChat(chat *domain.Chat) error {
	subQuery := s.database.Table("applications").Select("id").Where("token = ?", chat.AppToken)
	s.database.Select("id").Table("chats").Where("app_id = (?) AND number = ?", subQuery, chat.Number).Scan(&chat.ID)
	if chat.ID == 0 {
		return errors.New("cannot find specified application/chat combination")
	}
	return nil
}

func (s *Store) GetChatsByApplication(appToken string) ([]domain.Chat, error) {
	chats := make([]domain.Chat, 0)
	subQuery := s.database.Table("applications").Select("id").Where("token = ?", appToken)
	err := s.database.Select("*, ? AS app_token", appToken).Where("app_id = (?)", subQuery).Find(&chats).Error
	return chats, err
}

func (s *Store) GetChatByApplicationAndNumber(appToken string, number uint) (*domain.Chat, error) {
	chat := new(domain.Chat)
	subQuery := s.database.Table("applications").Select("id").Where("token = ?", appToken)
	err := s.database.Select("*, ? AS app_token", appToken).Where("app_id = (?) AND number = ?", subQuery, number).First(&chat).Error
	return chat, err
}
