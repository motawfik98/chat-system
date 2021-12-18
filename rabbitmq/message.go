package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (q *Queues) SendMessage(message *domain.Message) error {
	bytes, err := json.Marshal(message)
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
	message := new(domain.Message)
	err := json.Unmarshal(bytes, message)
	if err != nil {
		return nil, err
	}
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	message.ID = u.String()
	return message, nil
}
