package rabbitmq

import (
	"chat-system/domain"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *Queues) SendApplication(app *domain.Application, action Action) error {
	bytes, err := json.Marshal(struct {
		*domain.Application
		ID     uint
		Action Action
	}{app, app.ID, action})
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

func (q *Queues) ReceiveApplication(bytes []byte) (*domain.Application, error, Action) {
	type app struct {
		*domain.Application
		ID     uint
		Action Action
	}
	application := new(app)
	err := json.Unmarshal(bytes, application)
	if err != nil {
		return nil, err, 0
	}
	application.Application.ID = application.ID
	return application.Application, nil, application.Action
}
