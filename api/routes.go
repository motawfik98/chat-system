package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Routes() {
	h.router.GET("hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	})

	h.router.GET("/applications", h.HandleGetAllApplications)
	h.router.GET("/applications/:token", h.HandleGetApplicationByToken)
	h.router.POST("/applications", h.HandleCreateApplication)
	h.router.PUT("/applications/:token", h.HandleUpdateApplication)
	h.router.POST("/chat", h.HandleCreateChat)
	h.router.POST("/message", h.HandleCreateMessage)
}
