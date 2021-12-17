package workers

import (
	"chat-system/rabbitmq"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
)

func ConsumeApplicationsMessages(queues *rabbitmq.Queues, db *gorm.DB, rds *redis.Client) {
	messages, err := queues.Channel.Consume(
		queues.ApplicationQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			go func(d amqp.Delivery) {
				app, err := queues.ReceiveApplication(d.Body)
				if err == nil {
					err = db.Create(app).Error
					if err == nil {
						rds.HSetNX(context.Background(), app.Token, "number-of-chats", 0)
						d.Ack(false)
					}
				}
			}(d)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
