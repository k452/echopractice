package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/db"
)

func InitTable(tableName string) {
	db := db.Connect()
	defer db.Close()

	if tableName == "users" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (user_id INT NOT NULL AUTO_INCREMENT, name VARCHAR(100), mail VARCHAR(100), lank INT(10), pass VARCHAR(1000), created_at DATETIME, updated_at DATETIME, primary key(user_id));")
		if err != nil {
			logrus.Error(err)
		}
	} else if tableName == "recipe" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS recipe (id VARCHAR(1000), title VARCHAR(100), text VARCHAR(100000), sumbnail VARCHAR(100000), recipe_id VARCHAR(1000), created_at DATETIME, updateed_at DATETIME);")
		if err != nil {
			logrus.Error(err)
		}
	} else if tableName == "management" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS management (user_id VARCHAR(1000), sake_name VARCHAR(100), date DATE);")
		if err != nil {
			logrus.Error(err)
		}
	}
}
