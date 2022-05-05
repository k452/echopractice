package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
	"time"
)

func RegisterManagement(d *RegisterManagementType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO management(user_id, sake_name, amount, date, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.User_id, d.Sake_name, d.Amount, d.Date, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func SelectManagement(user_id int) []SelectManagementType {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM management WHERE user_id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)
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

func SelectAllManagement() []SelectManagementType {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM management;")
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

func UpdateManagement(d *UpdateManagementType) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE management SET amount=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Amount, time.Now(), d.Management_id)
	if err != nil {
		panic(err)
	}
}

func DeleteManagement(management_id int) {
	db := db.Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM management WHERE id=?")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(management_id)
	if err != nil {
		panic(err)
	}
}
