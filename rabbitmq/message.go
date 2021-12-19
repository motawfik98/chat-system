package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendMessage(message *domain.Message) error {
	bytes, err := json.Marshal(struct {
		*domain.Message
		ChatID uint
	}{message, message.ChatID})
	if err != nil {
		return err
	}
	return q.Channel.Publish(
		"",
		q.MessageQueue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bytes,
		})
}

func (q *Queues) ReceiveMessage(bytes []byte) (*domain.Message, error) {
	type messageWithIDs struct {
		*domain.Message
		AppID  uint
		ChatID uint
	}
	message := new(messageWithIDs)
	err := json.Unmarshal(bytes, message)
	if err != nil {
		return nil, err
	}
	message.Message.ChatID = message.ChatID

	return message.Message, nil
}
