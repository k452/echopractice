package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
	"time"
)

func RegisterRecipe(d *RegisterRecipeType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO recipe(user_id, title, text, sumbnail, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(d.User_id, d.Title, d.Text, d.Sumbnail, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func SelectRecipe(recipe_id int) []SelectRecipeType {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM recipe WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}

	rows, err := stmt.Query(recipe_id)
	if err != nil {
		logrus.Error(err)
	}

	var recipes []SelectRecipeType
	for rows.Next() {
		recipe := SelectRecipeType{}
		if err := rows.Scan(&recipe.Recipe_id, &recipe.User_id, &recipe.Title, &recipe.Text, &recipe.Sumbnail, &recipe.Created_at, &recipe.Updated_at); err != nil {
			logrus.Error(err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func SelectAllRecipe() []SelectRecipeType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM recipe;")
	if err != nil {
		logrus.Error(err)
	}

	var recipes []SelectRecipeType
	for rows.Next() {
		recipe := SelectRecipeType{}
		if err := rows.Scan(&recipe.Recipe_id, &recipe.User_id, &recipe.Title, &recipe.Text, &recipe.Sumbnail, &recipe.Created_at, &recipe.Updated_at); err != nil {
			logrus.Error(err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func UpdateRecipe(d *UpdateRecipeType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE recipe SET title=?, text=?, sumbnail=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Title, d.Text, d.Sumbnail, time.Now(), d.Recipe_id)
	if err != nil {
		panic(err)
	}
}

func DeleteRecipe(recipe_id int) {
	db := db.Connect()
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
