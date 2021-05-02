package routing

import (
	"github.com/labstack/echo"
	"net/http"
	"sake_io_api/model"
)

func usersRouting(e *echo.Echo) {
	e.GET("/users", func(c echo.Context) error {
		x := model.SelectAllUserData()
		return c.JSON(http.StatusOK, x)
	})
	e.POST("/users", func(c echo.Context) error {
		model.RegisterUserData()
		return c.String(http.StatusOK, "Insert")
	})
}
