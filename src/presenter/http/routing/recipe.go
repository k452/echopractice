package routing

import (
	"github.com/labstack/echo"
	"sake_io_api/infra/persistence"
	"sake_io_api/presenter/http/handler"
	"sake_io_api/usecase"
)

func recipeRouting(e *echo.Echo) {
	rh := handler.NewRecipeHandler(usecase.NewRecipeUsecase(persistence.NewRecipePersistence()))

	g := e.Group("/recipe")
	g.GET("", rh.GetRecipeAllHandle)
	g.GET("/:id", rh.GetRecipeHandle)
	g.POST("", rh.RegisterRecipeHandle)
	g.PATCH("/:id", rh.UpdateRecipeHandle)
	g.DELETE("/:id", rh.DeleteRecipeHandle)
}
