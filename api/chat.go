package api

import (
	"chat-system/domain"
	"github.com/labstack/echo/v4"
	"net/http"
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
