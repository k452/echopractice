package routing

import (
	"github.com/labstack/echo"
	"sake_io_api/infra/persistence"
	"sake_io_api/presenter/http/handler"
	"sake_io_api/usecase"
)

func userRouting(e *echo.Echo) {
	uh := handler.NewUserHandler(usecase.NewUserUsecase(persistence.NewUserPersistence()))

	g := e.Group("/user")
	g.GET("", uh.GetUserAllHandle)
	g.GET("/:id", uh.GetUserHandle)
	g.POST("", uh.RegisterUserHandle)
	g.PATCH("/:id", uh.UpdateUserHandle)
	g.DELETE("/:id", uh.DeleteUserHandle)
}
