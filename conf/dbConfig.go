package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/joho/godotenv"
)

type DBConnect struct {
	USER     string
	PASSWORD string
	DB       string
	HOST     string
	PORT     string
}

func GetDBInfo() *DBConnect {
	err := godotenv.Load(fmt.Sprintf("envFile/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	DBInfo := &DBConnect{
		USER:     os.Getenv("USER"),
		PASSWORD: os.Getenv("PASSWORD"),
		DB:       os.Getenv("DB"),
		HOST:     os.Getenv("HOST"),
		PORT:     os.Getenv("PORT"),
	}

	return DBInfo
}
