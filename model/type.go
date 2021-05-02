package model

import (
	"time"
)

type User struct {
	User_id    int       `json:"user_id"`
	Name       string    `json:"name"`
	Mail       string    `json:"mail"`
	Lank       int       `json:"lank"`
	Pass       string    `json:"pass"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
