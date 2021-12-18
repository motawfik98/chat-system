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

func ConsumeChatsMessages(queues *rabbitmq.Queues, db *gorm.DB, rds *redis.Client) {
	messages, err := queues.Channel.Consume(
		queues.ChatQueue.Name,
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
				chat, err := queues.ReceiveChat(d.Body)
				if err == nil {
					db.Select("id").Table("applications").Where("token = ?", chat.AppToken).Take(&chat.AppID)
					err = db.Create(chat).Error
					if err == nil {
						chatsCount := rds.HIncrBy(context.Background(), chat.AppToken, domain.TOTAL_CHATS, 1).Val()
						db.Table("applications").Where("id = ?", chat.AppID).Update("chats_count", chatsCount)
						rds.HSetNX(context.Background(), fmt.Sprintf("%s-%d", chat.AppToken, chat.Number), "number-of-messages", 0)
						d.Ack(false)
					}
				}
			}(d)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
