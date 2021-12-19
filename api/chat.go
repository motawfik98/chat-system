package api

import (
	"chat-system/domain"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (h *Handler) HandleCreateChat(c echo.Context) error {
	chat := new(domain.Chat)
	if err := c.Bind(chat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(chat); err != nil {
		return err
	}
	if err := h.store.CreateChat(c.Request().Context(), chat); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := h.queues.SendChat(chat); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chat)
}

func (h *Handler) HandleGetAllChatsByApplication(c echo.Context) error {
	appToken := c.Param("token")
	chats, err := h.store.GetChatsByApplication(appToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chats)
}

func (h *Handler) HandleGetChatByAppTokenAndNumber(c echo.Context) error {
	appToken := c.Param("token")
	number, err := strconv.ParseUint(c.Param("number"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	chats, err := h.store.GetChatByApplicationAndNumber(appToken, uint(number))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chats)
}
