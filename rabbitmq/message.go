package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendMessage(message *domain.Message, action Action) error {
	bytes, err := json.Marshal(struct {
		*domain.Message
		ID     uint
		ChatID uint
		Action Action
	}{message, message.ID, message.ChatID, action})
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

func (q *Queues) ReceiveMessage(bytes []byte) (*domain.Message, Action, error) {
	type messageWithIDs struct {
		*domain.Message
		ID     uint
		ChatID uint
		Action Action
	}
	message := new(messageWithIDs)
	err := json.Unmarshal(bytes, message)
	if err != nil {
		return nil, 0, err
	}
	message.Message.ID = message.ID
	message.Message.ChatID = message.ChatID

	return message.Message, message.Action, nil
}
