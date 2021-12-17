package api

import (
	"chat-system/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) HandleCreateMessage(c echo.Context) error {
	appToken := c.Param("token")
	chatNumber, err := strconv.ParseUint(c.Param("number"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	message := &domain.Message{
		AppToken:   appToken,
		ChatNumber: uint(chatNumber),
	}
	if err := c.Bind(message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(message); err != nil {
		return err
	}
	if err := h.store.CreateMessage(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, message)
}

func (h *Handler) HandleGetAllMessagesByApplicationAndChat(c echo.Context) error {
	appToken := c.Param("token")
	chatNumber := c.Param("number")
	messages, err := h.store.GetMessagesByApplicationAndChat(appToken, chatNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, messages)
}

func (h *Handler) HandleGetMessageByApplicationAndChatAndNumber(c echo.Context) error {
	appToken := c.Param("token")
	number, err := strconv.ParseUint(c.Param("number"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	msgNumber, err := strconv.ParseUint(c.Param("msg"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	chats, err := h.store.GetMessageByApplicationAndChatAndNumber(appToken, uint(number), uint(msgNumber))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chats)
}
