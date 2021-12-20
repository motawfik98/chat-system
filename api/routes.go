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

	h.router.GET("/applications/:token/chats", h.HandleGetAllChatsByApplication)
	h.router.GET("/applications/:token/chats/:number", h.HandleGetChatByAppTokenAndNumber)
	h.router.POST("/applications/:token/chats", h.HandleCreateChat)
	h.router.PUT("/applications/:token/chats/:number", h.HandleUpdateChat)

	h.router.GET("/applications/:token/chats/:number/messages", h.HandleGetAllMessagesByApplicationAndChat)
	h.router.GET("/applications/:token/chats/:number/messages/:msg", h.HandleGetMessageByApplicationAndChatAndNumber)
	h.router.POST("/applications/:token/chats/:number/messages", h.HandleCreateMessage)
	h.router.POST("/applications/:token/chats/:number/search", h.HandleSearchMessages)
}
