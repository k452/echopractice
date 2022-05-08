package utils

import (
	"time"
)

func StringToTime(str string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", str)
	return t
}
