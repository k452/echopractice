package routing

import (
	"github.com/labstack/echo"
	"sake_io_api/presenter/http/handler"
)

func otherRouting(e *echo.Echo) {
	oh := handler.NewOtherHandler()

	e.GET("/", oh.HealthCheckHandle)
}
