package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
	"time"
)

func SelectUserData(user_id int) []SelectUserType {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM users WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}

	rows, err := stmt.Query(user_id)
	if err != nil {
		logrus.Error(err)
	}

	var users []SelectUserType
	for rows.Next() {
		user := SelectUserType{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Mail, &user.Level, &user.Pass, &user.Created_at, &user.Updated_at); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}

func RegisterUser(d *RegisterUserType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(name, mail, level, pass, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(d.Name, d.Mail, 1, d.Pass, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func SelectAllUserData() []SelectUserType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		logrus.Error(err)
	}

	var users []SelectUserType
	for rows.Next() {
		user := SelectUserType{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Mail, &user.Lank, &user.Pass, &user.Created_at, &user.Updated_at); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}
