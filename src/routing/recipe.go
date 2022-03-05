package routing

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/model"
	"strconv"
)

func recipeRouting(e *echo.Echo) {
	g := e.Group("/recipe")
	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.SelectAllRecipe())
	})
	g.GET("/:id", func(c echo.Context) error {
		recipe_id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.SelectRecipe(recipe_id))
	})
	g.POST("", func(c echo.Context) error {
		var formData model.RegisterRecipeType
		if err := c.Bind(&formData); err != nil {
			logrus.Error(err)
		}
		model.RegisterRecipe(&formData)
		return c.String(http.StatusOK, "success create recipe")
	})
	g.PATCH("/:id", func(c echo.Context) error {
		var formData model.UpdateRecipeType
		formData.Recipe_id, _ = strconv.Atoi(c.Param("id"))
		formData.Title = c.FormValue("title")
		formData.Text = c.FormValue("text")
		formData.Sumbnail = c.FormValue("sumbnail")
		model.UpdateRecipe(&formData)
		return c.String(http.StatusOK, "success update recipe")
	})
	g.DELETE("/:id", func(c echo.Context) error {
		recipe_id, _ := strconv.Atoi(c.Param("id"))
		model.DeleteRecipe(recipe_id)
		return c.String(http.StatusOK, "success delete recipe")
	})
}
