package api

import (
	"chat-system/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) HandleCreateMessage(c echo.Context) error {
	message := new(domain.Message)
	if err := c.Bind(message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(message); err != nil {
		return err
	}
	if err := h.dnConn.CreateMessage(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, message)
}