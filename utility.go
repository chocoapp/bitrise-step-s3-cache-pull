package main

import (
	"log"
	"os"
)

func GetEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing variable '%s'\n", key)
	}
	return value
}
