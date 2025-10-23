package utils

import (
	"os"
)

func GetConfig() *Config {
	return &Config{
		Port:   ThreeTermString(len(os.Getenv("PORT")) > 0, os.Getenv("PORT"), "3000"),
		Secret: ThreeTermString(len(os.Getenv("SECRET")) > 0, os.Getenv("SECRET"), "SECRET"),
	}
}
