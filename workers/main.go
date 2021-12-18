package main

import (
	"chat-system/database"
	"chat-system/rabbitmq"
	"chat-system/redis"
	"chat-system/workers/workers"
)

func main() {
	rabbitConn, queues, err := rabbitmq.SetupRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer rabbitConn.Close()
	defer queues.Channel.Close()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	rds := redis.SetupRedis()

	forever := make(chan bool)

	go workers.ConsumeApplicationsMessages(queues, db, rds)
	go workers.ConsumeChatsMessages(queues, db, rds)
	go workers.ConsumeMessages(queues, db, rds)

	<-forever

}
