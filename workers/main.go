package main

import (
	"chat-system/database"
	"chat-system/rabbitmq"
	"chat-system/redis"
	"chat-system/workers/workers"
)

func main() {
	rabbitConn, queues, err := rabbitmq.Setup()
	if err != nil {
		panic(err)
	}
	defer rabbitConn.Close()
	defer queues.Channel.Close()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	rds := redis.Setup()

	forever := make(chan bool)

	go workers.ConsumeApplications(queues, db, rds)
	go workers.ConsumeChats(queues, db, rds)
	go workers.ConsumeMessages(queues, db, rds)

	<-forever

}
