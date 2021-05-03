package routing

import (
	"github.com/labstack/echo"
	//"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/model"
)

func managementRouting(e *echo.Echo) {
	g := e.Group("/management")
	g.GET("/:user_id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.SelectAllManagement(c.Param("user_id")))
	})
	/*
		g.POST("", func(c echo.Context) error {
			var formData model.RegisterUserType
			if err := c.Bind(&formData); err != nil {
				logrus.Error(err)
			}
			model.RegisterUser(&formData)
			return c.String(http.StatusOK, "Success")
		})
		g.GET("/:user_id", func(c echo.Context) error {
			return c.JSON(http.StatusOK, model.SelectUserData(c.Param("user_id")))
		})
	*/
}
