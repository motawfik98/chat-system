package main

import (
	"chat-system/api"
	"chat-system/database"
	"chat-system/rabbitmq"
	"chat-system/redis"
	"chat-system/service"
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

	info := service.NewInfoService(db, redisClient)
	api.RegisterAPIHandler(info, queues, e)

	e.Logger.Fatal(e.Start(":3000"))
}
