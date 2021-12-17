package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (q *Queues) SendChat(chat *domain.Chat) error {
	bytes, err := json.Marshal(chat)
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
	chat := new(domain.Chat)
	err := json.Unmarshal(bytes, chat)
	if err != nil {
		return nil, err
	}
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	chat.ID = u.String()
	return chat, nil
}
