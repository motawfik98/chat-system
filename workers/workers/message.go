package workers

import (
	"chat-system/domain"
	"chat-system/rabbitmq"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
)

func ConsumeMessages(queues *rabbitmq.Queues, db *gorm.DB, rds *redis.Client) {
	messages, err := queues.Channel.Consume(
		queues.MessageQueue.Name,
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
				message, err := queues.ReceiveMessage(d.Body)
				if err == nil {
					db.Select("id").Table("applications").Where("token = ?", message.AppToken).Take(&message.AppID)
					db.Select("id").Table("chats").Where("app_id = ? AND number = ?", message.AppID, message.ChatNumber).Take(&message.ChatID)
					err = db.Create(message).Error
					if err == nil {
						messagesCount := rds.HIncrBy(context.Background(), fmt.Sprintf("%s-%d", message.AppToken, message.ChatNumber), domain.TOTAL_MESSAGES, 1).Val()
						db.Table("chats").Where("id = ?", message.ChatID).Update("messages_count", messagesCount)
						d.Ack(false)
					}
				}
			}(d)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
