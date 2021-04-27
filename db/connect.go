package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"sake_io_api/conf"
)

func Connect() *sql.DB {
	DBInfo := conf.GetDBInfo()
	db, err := sql.Open("mysql", DBInfo.USER+":"+DBInfo.PASSWORD+"@tcp("+DBInfo.HOST+":"+DBInfo.PORT+")/"+DBInfo.DB)

	if err != nil {
		logrus.Error(err)
	}

	return db
}
