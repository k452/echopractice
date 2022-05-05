package routing

import (
	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {
	userRouting(e)
	managementRouting(e)
	recipeRouting(e)
	otherRouting(e)
}
