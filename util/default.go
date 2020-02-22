package util

import (
	"strconv"

	"github.com/pkg/errors"
)

func DefaultInt(m map[string]string, key string, d int) int {
	item, ok := m[key]
	if ok {
		value, err := strconv.Atoi(item)
		if err == nil {
			return value
		}
	}
	return d
}

func GetStringKey(m map[string]string, key string) (string, error) {
	item, ok := m[key]
	if ok {
		return item, nil
	}
	return "", errors.New(key + "is required")
}
