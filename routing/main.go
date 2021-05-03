package routing

import (
	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {
	usersRouting(e)
}
