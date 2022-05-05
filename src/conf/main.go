package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/joho/godotenv"
)

func GetDomain() string {
	err := godotenv.Load(fmt.Sprintf("envFiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	return os.Getenv("DOMAIN")
}
