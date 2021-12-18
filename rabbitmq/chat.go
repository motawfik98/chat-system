package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendChat(chat *domain.Chat) error {
	// custom Marshal to send AppID to queue
	bytes, err := json.Marshal(struct {
		*domain.Chat
		AppID uint
	}{chat, chat.AppID})
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

func (q *Queues) ReceiveChat(bytes []byte) (*domain.Chat, error) {
	type chatWithIDs struct {
		*domain.Chat
		AppID uint
	}
	chat := new(chatWithIDs)
	err := json.Unmarshal(bytes, &chat)

	if err != nil {
		return nil, err
	}
	chat.Chat.AppID = chat.AppID
	return chat.Chat, nil

}
