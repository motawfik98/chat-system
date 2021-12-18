package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendApplication(app *domain.Application) error {
	bytes, err := json.Marshal(app)
	if err != nil {
		return err
	}
	return q.Channel.Publish(
		"",
		q.ApplicationQueue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bytes,
		})
}

func (q *Queues) ReceiveApplication(bytes []byte) (*domain.Application, error) {
	app := new(domain.Application)
	err := json.Unmarshal(bytes, app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
