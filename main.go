package main

import (
	"github.com/labstack/echo"
	"sake_io_api/routing"
)

func main() {
	e := echo.New()
	routing.Routing(e)
	e.Logger.Fatal(e.Start(":8080"))
}
