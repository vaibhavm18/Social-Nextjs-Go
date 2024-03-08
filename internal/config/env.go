package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	fmt.Println(path)
	err = godotenv.Load(filepath.Join(path, ".env"))

	if err != nil {
		return err
	}

	return nil
}
