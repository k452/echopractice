package usecase

import (
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
)

type RecipeUsecase interface {
	GetRecipeAll() []model.Recipe
	GetRecipe(recipe_id int) []model.Recipe
	RegisterRecipe(u *model.RegisterRecipeType)
	UpdateRecipe(u *model.UpdateRecipeType)
	DeleteRecipe(recipe_id int)
}

//今はリポジトリのラッパのようになってるけど、よりビジネスロジックが複雑になった時に他の処理を組み合わせたりできるようにしていると解釈
type recipeUsecase struct {
	recipeRepository repository.RecipeRepository
}

func NewRecipeUsecase(rr repository.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{
		recipeRepository: rr,
	}
}

func (ru recipeUsecase) GetRecipeAll() []model.Recipe {
	return ru.recipeRepository.GetRecipeAll()
}

func (ru recipeUsecase) GetRecipe(recipe_id int) []model.Recipe {
	return ru.recipeRepository.GetRecipe(recipe_id)
}

func (ru recipeUsecase) RegisterRecipe(u *model.RegisterRecipeType) {
	ru.recipeRepository.RegisterRecipe(u)
}

func (ru recipeUsecase) UpdateRecipe(u *model.UpdateRecipeType) {
	ru.recipeRepository.UpdateRecipe(u)
}

func (ru recipeUsecase) DeleteRecipe(recipe_id int) {
	ru.recipeRepository.DeleteRecipe(recipe_id)
}
