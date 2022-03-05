package routing

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/model"
	"strconv"
)

func userRouting(e *echo.Echo) {
	g := e.Group("/user")
	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.SelectAllUser())
	})
	g.GET("/:id", func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.SelectUser(user_id))
	})
	g.POST("", func(c echo.Context) error {
		var formData model.RegisterUserType
		if err := c.Bind(&formData); err != nil {
			logrus.Error(err)
		}
		model.RegisterUser(&formData)
		return c.String(http.StatusOK, "success create user")
	})
	g.PATCH("/:id", func(c echo.Context) error {
		var formData model.UpdateUserType
		formData.User_id, _ = strconv.Atoi(c.Param("id"))
		formData.Name = c.FormValue("name")
		formData.Mail = c.FormValue("mail")
		formData.Level, _ = strconv.Atoi(c.FormValue("level"))
		formData.Pass = c.FormValue("pass")
		model.UpdateUser(&formData)
		return c.String(http.StatusOK, "success update user")
	})
	g.DELETE("/:id", func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("id"))
		model.DeleteUser(user_id)
		return c.String(http.StatusOK, "success delete user")
	})
}
