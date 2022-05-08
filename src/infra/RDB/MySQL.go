package RDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type MySQLConnectionInfo struct {
	USER     string
	PASSWORD string
	DB       string
	HOST     string
	PORT     string
}

func getMySQLConnectionInfo() *MySQLConnectionInfo {
	err := godotenv.Load(fmt.Sprintf("envFiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	info := &MySQLConnectionInfo{
		USER:     os.Getenv("USER"),
		PASSWORD: os.Getenv("PASSWORD"),
		DB:       os.Getenv("DB"),
		HOST:     os.Getenv("HOST"),
		PORT:     os.Getenv("PORT"),
	}

	return info
}

func getTestMySQLConnectionInfo() *MySQLConnectionInfo {
	envPath := fmt.Sprintf("/go/src/sake_io_api/envFiles/%s.env", os.Getenv("GO_ENV"))
	err := godotenv.Load(envPath)
	if err != nil {
		logrus.Fatalf("Error loading .env file loadPath → %s", envPath)
	}

	info := &MySQLConnectionInfo{
		USER:     os.Getenv("USER"),
		PASSWORD: os.Getenv("PASSWORD"),
		DB:       os.Getenv("TEST_DB"),
		HOST:     os.Getenv("HOST"),
		PORT:     os.Getenv("PORT"),
	}

	return info
}

func MySQLConnect(isProd bool) *sql.DB {
	var DBInfo *MySQLConnectionInfo
	if isProd {
		DBInfo = getMySQLConnectionInfo()
	} else {
		DBInfo = getTestMySQLConnectionInfo()
	}
	db, err := sql.Open("mysql", DBInfo.USER+":"+DBInfo.PASSWORD+"@tcp("+DBInfo.HOST+":"+DBInfo.PORT+")/"+DBInfo.DB+"?parseTime=true")

	if err != nil {
		logrus.Error(err)
	}

	return db
}
