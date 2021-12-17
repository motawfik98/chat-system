package rabbitmq

import "github.com/streadway/amqp"

type Queues struct {
	Channel          *amqp.Channel
	ApplicationQueue *amqp.Queue
}

func SetupRabbitMQ() (*amqp.Connection, *Queues, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	appQueue, err := ch.QueueDeclare(
		"applications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	queues := Queues{Channel: ch, ApplicationQueue: &appQueue}
	return conn, &queues, nil
}
