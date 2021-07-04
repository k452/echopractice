package routing

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/model"
	"strconv"
)

func usersRouting(e *echo.Echo) {
	g := e.Group("/users")
	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.SelectAllUserData())
	})
	g.POST("", func(c echo.Context) error {
		var formData model.RegisterUserType
		if err := c.Bind(&formData); err != nil {
			logrus.Error(err)
		}
		model.RegisterUser(&formData)
		return c.String(http.StatusOK, "Success")
	})
	g.GET("/:id", func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.SelectUserData(user_id))
	})
}
