package model

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"sake_io_api/db"
	"time"
)

func InsertData() {
	rand.Seed(time.Now().UnixNano())

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (name VARCHAR(100), age INT(5));")
	if err != nil {
		logrus.Error(err)
	}

	ins, err2 := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?);")
	if err2 != nil {
		logrus.Error(err2)
	}

	ins.Exec("Karasawa", rand.Intn(100))
}

func SelectData() []User {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		logrus.Error(err)
	}

	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Name, &user.Age); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}

type User struct {
	Name string
	Age  string
}
