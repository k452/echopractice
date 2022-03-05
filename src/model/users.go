package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
	"time"
)

func RegisterUser(d *RegisterUserType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(name, mail, level, pass, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(d.Name, d.Mail, 1, d.Pass, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func SelectUser(user_id int) []SelectUserType {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM user WHERE id=?;")
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

func SelectAllUser() []SelectUserType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user;")
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

func UpdateUser(d *UpdateUserType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET name=?, mail=?, level=?, pass=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Name, d.Mail, d.Level, d.Pass, time.Now(), d.User_id)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(user_id int) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		panic(err)
	}
}
