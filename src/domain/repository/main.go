package repository

import (
	"sake_io_api/domain/model"
)

type UserRepository interface {
	GetUserAll() []model.User
	GetUser(user_id int) []model.User
	RegisterUser(u *model.RegisterUserType)
	UpdateUser(u *model.UpdateUserType)
	DeleteUser(user_id int)
}

type ManagementRepository interface {
	GetManagementAll() []model.Management
	GetManagement(user_id int) []model.Management
	RegisterManagement(m *model.RegisterManagementType)
	UpdateManagement(m *model.UpdateManagementType)
	DeleteManagement(user_id int)
}

type RecipeRepository interface {
	GetRecipeAll() []model.Recipe
	GetRecipe(user_id int) []model.Recipe
	RegisterRecipe(r *model.RegisterRecipeType)
	UpdateRecipe(r *model.UpdateRecipeType)
	DeleteRecipe(user_id int)
}
