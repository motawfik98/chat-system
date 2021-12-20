package api

import (
	"chat-system/domain"
	"chat-system/rabbitmq"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	if err := h.store.CreateApplication(application); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := h.queues.SendApplication(application, rabbitmq.Create); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}

func (h *Handler) HandleGetAllApplications(c echo.Context) error {
	applications, err := h.store.GetApplications()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, applications)
}

func (h *Handler) HandleGetApplicationByToken(c echo.Context) error {
	appToken := c.Param("token")
	application, err := h.store.GetApplicationByToken(appToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}

func (h *Handler) HandleUpdateApplication(c echo.Context) error {
	application := new(domain.Application)
	if err := c.Bind(application); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(application); err != nil {
		return err
	}
	if err := h.store.UpdateApplication(application); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := h.queues.SendApplication(application, rabbitmq.Update); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}
