package persistence

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
	"sake_io_api/infra/RDB"
	"time"
)

type recipePersistence struct{}

func NewRecipePersistence() repository.RecipeRepository {
	return &recipePersistence{}
}

func (up recipePersistence) GetRecipeAll() []model.Recipe {
	db := RDB.MySQLConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM recipe;")
	if err != nil {
		logrus.Error(err)
	}

	var recipes []model.Recipe
	for rows.Next() {
		recipe := model.Recipe{}
		if err := rows.Scan(&recipe.Recipe_id, &recipe.User_id, &recipe.Title, &recipe.Text, &recipe.Sumbnail, &recipe.Created_at, &recipe.Updated_at); err != nil {
			logrus.Error(err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func (up recipePersistence) GetRecipe(recipe_id int) []model.Recipe {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM recipe WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}

	rows, err := stmt.Query(recipe_id)
	if err != nil {
		logrus.Error(err)
	}

	var recipes []model.Recipe
	for rows.Next() {
		recipe := model.Recipe{}
		if err := rows.Scan(&recipe.Recipe_id, &recipe.User_id, &recipe.Title, &recipe.Text, &recipe.Sumbnail, &recipe.Created_at, &recipe.Updated_at); err != nil {
			logrus.Error(err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func (up recipePersistence) RegisterRecipe(r *model.RegisterRecipeType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO recipe(user_id, title, text, sumbnail, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(r.User_id, r.Title, r.Text, r.Sumbnail, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func (up recipePersistence) UpdateRecipe(r *model.UpdateRecipeType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE recipe SET title=?, text=?, sumbnail=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Title, r.Text, r.Sumbnail, time.Now(), r.Recipe_id)
	if err != nil {
		panic(err)
	}
}

func (up recipePersistence) DeleteRecipe(recipe_id int) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM recipe WHERE id=?")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(recipe_id)
	if err != nil {
		panic(err)
	}
}
