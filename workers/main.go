package main

import (
	"chat-system/database"
	"chat-system/rabbitmq"
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

	forever := make(chan bool)

	go workers.ConsumeApplicationsMessages(queues, db)

	<-forever

}
