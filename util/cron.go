package util

import (
	"errors"
	"strings"
)

func CronExpression(hoursMinutes string) (string, error) {
	hm := strings.Split(hoursMinutes, ":")
	if len(hm) != 2 {
		return "", errors.New("Invalid input")
	}
	return hm[1] + " " + hm[0] + " * * 1,2,3,4,5", nil
}
