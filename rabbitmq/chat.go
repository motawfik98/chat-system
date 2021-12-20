package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendChat(chat *domain.Chat, action Action) error {
	// custom Marshal to send AppID to queue
	bytes, err := json.Marshal(struct {
		*domain.Chat
		ID     uint
		AppID  uint
		Action Action
	}{chat, chat.ID, chat.AppID, action})
	if err != nil {
		return err
	}
	return q.Channel.Publish(
		"",
		q.ChatQueue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bytes,
		})
}

func (q *Queues) ReceiveChat(bytes []byte) (*domain.Chat, Action, error) {
	type chatWithIDs struct {
		*domain.Chat
		ID     uint
		AppID  uint
		Action Action
	}
	chat := new(chatWithIDs)
	err := json.Unmarshal(bytes, &chat)

	if err != nil {
		return nil, 0, err
	}
	chat.Chat.AppID = chat.AppID
	chat.Chat.ID = chat.ID
	return chat.Chat, chat.Action, nil

}
