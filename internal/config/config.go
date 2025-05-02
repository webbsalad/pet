package config

import (
	"log"
	"os"
)

type Config struct {
	STORAGE_PATH string `yaml:"storage_path"`
}

func NewConfig() Config {
	storage_path := os.Getenv("STORAGE_PATH")
	if storage_path == "" {
		log.Fatal("environment variable STORAGE_PATH is required")
	}

	return Config{
		STORAGE_PATH: storage_path,
	}
}
