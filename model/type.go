package model

import (
	"time"
)

type SelectUserType struct {
	User_id    int       `json:"user_id"`
	Name       string    `json:"name"`
	Mail       string    `json:"mail"`
	Lank       int       `json:"lank"`
	Pass       string    `json:"pass"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type RegisterUserType struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type SelectManagementType struct {
	Management_id int       `json:"management_id"`
	User_id       int       `json:"user_id"`
	Sake_name     string    `json:"sake_name"`
	Amount        int       `json:"amount"`
	Date          time.Time `json:"date"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type RegisterManagementType struct {
	User_id   int       `json:"user_id"`
	Sake_name string    `json:"sake_name"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date"`
}

type UpdateManagementType struct {
	Management_id int `json:"management_id"`
	Amount        int `json:"amount"`
}
