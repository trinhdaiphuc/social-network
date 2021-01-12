package config

import "os"

func env(key, defaultValue string) (value string) {
	if value = os.Getenv(key); value == "" {
		value = defaultValue
	}
	return
}