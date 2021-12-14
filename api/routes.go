package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Routes() {
	h.router.GET("hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	})

	h.router.POST("/application", h.HandleCreateApplication)
	h.router.POST("/chat", h.HandleCreateChat)
	h.router.POST("/message", h.HandleCreateMessage)
}
