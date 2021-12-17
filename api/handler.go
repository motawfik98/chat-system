package api

import (
	"chat-system/rabbitmq"
	"chat-system/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	dnConn *service.Info
	queues *rabbitmq.Queues
	router *echo.Echo
}

func RegisterAPIHandler(info *service.Info, queues *rabbitmq.Queues, router *echo.Echo) {
	handler := &Handler{
		dnConn: info,
		queues: queues,
		router: router,
	}
	handler.Routes()
}
