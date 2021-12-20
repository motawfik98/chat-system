package rabbitmq

import (
	"github.com/streadway/amqp"
	"os"
)

type Queues struct {
	Channel          *amqp.Channel
	ApplicationQueue *amqp.Queue
	ChatQueue        *amqp.Queue
	MessageQueue     *amqp.Queue
}

type Action int

const (
	Create Action = iota
	Update
)

func Setup() (*amqp.Connection, *Queues, error) {
	conn, err := amqp.Dial(os.ExpandEnv("amqp://guest:guest@${RABBITMQ_HOST}:${RABBITMQ_PORT}/"))
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	appQueue, err := declareQueue("applications", ch)
	if err != nil {
		return nil, nil, err
	}
	chatQueue, err := declareQueue("chats", ch)
	if err != nil {
		return nil, nil, err
	}
	messageQueue, err := declareQueue("messages", ch)
	if err != nil {
		return nil, nil, err
	}
	queues := Queues{Channel: ch, ApplicationQueue: &appQueue, ChatQueue: &chatQueue, MessageQueue: &messageQueue}
	return conn, &queues, nil
}

func declareQueue(queueName string, channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}
