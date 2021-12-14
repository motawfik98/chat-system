package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	dnConn *gorm.DB
	router *echo.Echo
}

func RegisterAPIHandler(dbConn *gorm.DB, router *echo.Echo) {
	handler := &Handler{
		dnConn: dbConn,
		router: router,
	}
	handler.Routes()
}
