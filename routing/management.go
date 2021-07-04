package routing

import (
	"github.com/labstack/echo"
	"net/http"
	"sake_io_api/model"
	"strconv"
	"time"
)

func managementRouting(e *echo.Echo) {
	g := e.Group("/management")
	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.SelectAllManagement())
	})
	g.GET("/:id", func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.SelectManagement(user_id))
	})
	g.POST("/:id", func(c echo.Context) error {
		var formData model.RegisterManagementType
		formData.User_id, _ = strconv.Atoi(c.Param("id"))
		formData.Sake_name = c.FormValue("sake_name")
		formData.Amount, _ = strconv.Atoi(c.FormValue("amount"))
		formData.Date, _ = time.Parse("2006-01-02", c.FormValue("date"))
		model.RegisterManagement(&formData)
		return c.String(http.StatusOK, "success create management")
	})
	g.PATCH("/:id", func(c echo.Context) error {
		var formData model.UpdateManagementType
		formData.Management_id, _ = strconv.Atoi(c.Param("id"))
		formData.Amount, _ = strconv.Atoi(c.FormValue("amount"))
		model.UpdateManagement(&formData)
		return c.String(http.StatusOK, "success update management")
	})
	g.DELETE("/:id", func(c echo.Context) error {
		management_id, _ := strconv.Atoi(c.Param("id"))
		model.DeleteManagement(management_id)
		return c.String(http.StatusOK, "success delete management")
	})
}
