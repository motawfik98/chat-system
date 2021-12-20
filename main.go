package main

import (
	"chat-system/api"
	"chat-system/database"
	"chat-system/rabbitmq"
	"chat-system/redis"
	"chat-system/service"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &api.CustomValidator{Validator: validator.New()}

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	rabbitConn, queues, err := rabbitmq.Setup()
	if err != nil {
		panic(err)
	}
	defer rabbitConn.Close()
	defer queues.Channel.Close()

	redisClient := redis.Setup()

	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	info := service.NewInfoService(db, redisClient, esClient)
	api.RegisterAPIHandler(info, queues, e)

	e.Logger.Fatal(e.Start(":3000"))
}
