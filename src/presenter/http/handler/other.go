package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type OtherHandler interface {
	HealthCheckHandle(c echo.Context) error
}

type otherHandler struct{}

func NewOtherHandler() OtherHandler {
	return &otherHandler{}
}

func (uh otherHandler) HealthCheckHandle(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
