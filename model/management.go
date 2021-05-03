package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
)

/*
func SelectManagement(user_id string) []SelectUserType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE user_id=%s;", "users", user_id))
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

func RegisterManagement(d *RegisterUserType) {
	db := db.Connect()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO users(name, mail, lank, pass, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	ins.Exec(d.Name, d.Mail, 1, d.Pass, time.Now(), time.Now())
}
*/
func SelectAllManagement(user_id string) []SelectManagementType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE user_id=%s;", "management", user_id))
	if err != nil {
		logrus.Error(err)
	}

	var managements []SelectManagementType
	for rows.Next() {
		management := SelectManagementType{}
		if err := rows.Scan(&management.Management_id, &management.User_id, &management.Sake_name, &management.Amount, &management.Date, &management.Created_at, &management.Updated_at); err != nil {
			logrus.Error(err)
		}
		managements = append(managements, management)
	}
	return managements
}
