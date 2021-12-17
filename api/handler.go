package api

import (
	"chat-system/rabbitmq"
	"chat-system/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store  *service.Store
	queues *rabbitmq.Queues
	router *echo.Echo
}

func RegisterAPIHandler(info *service.Store, queues *rabbitmq.Queues, router *echo.Echo) {
	handler := &Handler{
		store:  info,
		queues: queues,
		router: router,
	}
	handler.Routes()
}
