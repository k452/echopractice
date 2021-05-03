package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
	"time"
)

func RegisterUserData() {
	db := db.Connect()
	defer db.Close()

	InitTable("users")

	ins, err := db.Prepare("INSERT INTO users(name, mail, lank, pass, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	ins.Exec("Karasawa", "karasawauec@gmail.com", 1, "password", time.Now(), time.Now())
}

func SelectAllUserData() []User {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		logrus.Error(err)
	}

	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Mail, &user.Lank, &user.Pass, &user.Created_at, &user.Updated_at); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}
