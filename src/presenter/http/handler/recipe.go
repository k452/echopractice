package handler

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/domain/model"
	"sake_io_api/usecase"
	"strconv"
)

type RecipeHandler interface {
	GetRecipeAllHandle(c echo.Context) error
	GetRecipeHandle(c echo.Context) error
	RegisterRecipeHandle(c echo.Context) error
	UpdateRecipeHandle(c echo.Context) error
	DeleteRecipeHandle(c echo.Context) error
}

type recipeHandler struct {
	recipeUsecase usecase.RecipeUsecase
}

func NewRecipeHandler(ru usecase.RecipeUsecase) RecipeHandler {
	return &recipeHandler{
		recipeUsecase: ru,
	}
}

func (rh recipeHandler) GetRecipeAllHandle(c echo.Context) error {
	return c.JSON(http.StatusOK, rh.recipeUsecase.GetRecipeAll())
}

func (rh recipeHandler) GetRecipeHandle(c echo.Context) error {
	recipe_id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, rh.recipeUsecase.GetRecipe(recipe_id))
}

func (rh recipeHandler) RegisterRecipeHandle(c echo.Context) error {
	var formData model.RegisterRecipeType
	if err := c.Bind(&formData); err != nil {
		logrus.Error(err)
	}
	rh.recipeUsecase.RegisterRecipe(&formData)
	return c.String(http.StatusOK, "success create recipe")
}

func (rh recipeHandler) UpdateRecipeHandle(c echo.Context) error {
	var formData model.UpdateRecipeType
	formData.Recipe_id, _ = strconv.Atoi(c.Param("id"))
	formData.Title = c.FormValue("title")
	formData.Text = c.FormValue("text")
	formData.Sumbnail = c.FormValue("sumbnail")
	rh.recipeUsecase.UpdateRecipe(&formData)
	return c.String(http.StatusOK, "success update recipe")
}

func (rh recipeHandler) DeleteRecipeHandle(c echo.Context) error {
	recipe_id, _ := strconv.Atoi(c.Param("id"))
	rh.recipeUsecase.DeleteRecipe(recipe_id)
	return c.String(http.StatusOK, "success delete recipe")
}
