package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"sake_io_api/model"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		x := model.SelectData()
		fmt.Println(x)
		return c.String(http.StatusOK, "Select")
	})
	e.GET("/i", func(c echo.Context) error {
		model.InsertData()
		return c.String(http.StatusOK, "Insert")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
