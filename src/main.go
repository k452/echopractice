package main

import (
	"sake_io_api/presenter/http/middleWare"
	"sake_io_api/presenter/http/routing"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	middleWare.UseMiddleWare(e)
	routing.Routing(e)
	e.Logger.Fatal(e.Start(":8080"))
}
