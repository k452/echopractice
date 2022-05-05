package routing

import (
	"github.com/labstack/echo"
	"sake_io_api/infra/persistence"
	"sake_io_api/presenter/http/handler"
	"sake_io_api/usecase"
)

func managementRouting(e *echo.Echo) {
	mh := handler.NewManagementHandler(usecase.NewManagementUsecase(persistence.NewManagementPersistence()))

	g := e.Group("/management")
	g.GET("", mh.GetManagementAllHandle)
	g.GET("/:id", mh.GetManagementHandle)
	g.POST("/:id", mh.RegisterManagementHandle)
	g.PATCH("/:id", mh.UpdateManagementHandle)
	g.DELETE("/:id", mh.DeleteManagementHandle)
}
