package workers

import (
	"chat-system/domain"
	"chat-system/rabbitmq"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
)

func ConsumeApplications(queues *rabbitmq.Queues, db *gorm.DB, rds *redis.Client) {
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
						rds.HSetNX(context.Background(), app.Token, domain.TOTAL_CHATS, 0)
						rds.HSetNX(context.Background(), app.Token, domain.MAX_CHAT_NUMBER, 0)
						d.Ack(false)
					}
				}
			}(d)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
