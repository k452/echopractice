package model

import (
	"time"
)

type Recipe struct {
	Recipe_id  int       `json:"recipe_id"`
	User_id    int       `json:"user_id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	Sumbnail   string    `json:"sumbnail"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type RegisterRecipeType struct {
	User_id  int    `json:"user_id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Sumbnail string `json:"sumbnail"`
}

type UpdateRecipeType struct {
	Recipe_id int    `json:"recipe_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Sumbnail  string `json:"sumbnail"`
}
