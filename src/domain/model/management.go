package model

import (
	"time"
)

type Management struct {
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
