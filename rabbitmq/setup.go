package rabbitmq

import "github.com/streadway/amqp"

type Queues struct {
	Channel          *amqp.Channel
	ApplicationQueue *amqp.Queue
	ChatQueue        *amqp.Queue
	MessageQueue     *amqp.Queue
}

func Setup() (*amqp.Connection, *Queues, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
