package utils

import "os"

func GetenvWithDefaultValue(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
