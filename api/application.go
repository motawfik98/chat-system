package api

import (
	"chat-system/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) HandleCreateApplication(c echo.Context) error {
	application := new(domain.Application)
	if err := c.Bind(application); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(application); err != nil {
		return err
	}
	if err := h.dnConn.CreateApplication(application); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}
