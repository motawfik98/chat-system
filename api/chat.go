package api

import (
	"chat-system/domain"
	"github.com/labstack/echo/v4"
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
	if err := h.dnConn.CreateChat(chat); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chat)
}

func (h *Handler) HandleGetAllChatsByApplication(c echo.Context) error {
	appToken := c.Param("token")
	chats, err := h.dnConn.GetChatsByApplication(appToken)
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
	chats, err := h.dnConn.GetChatByApplicationAndNumber(appToken, uint(number))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chats)
}
