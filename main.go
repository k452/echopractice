package main

import (
	"github.com/labstack/echo"
	"net/http"
	"sake_io_api/model"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		x := model.SelectAllUserData()
		return c.JSON(http.StatusOK, x)
	})
	e.GET("/r", func(c echo.Context) error {
		model.RegisterUserData()
		return c.String(http.StatusOK, "Insert")
	})
	e.GET("/i", func(c echo.Context) error {
		model.InitTable("management")
		return c.String(http.StatusOK, "Insert")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
