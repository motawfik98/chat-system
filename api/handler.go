package api

import (
	"chat-system/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	dnConn *service.Info
	router *echo.Echo
}

func RegisterAPIHandler(info *service.Info, router *echo.Echo) {
	handler := &Handler{
		dnConn: info,
		router: router,
	}
	handler.Routes()
}
