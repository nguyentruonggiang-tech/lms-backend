package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	IsProduction bool
	Host         string
	Port         string
	DatabaseURL  string
}

func New() *Env {
	godotenv.Load()

	e := &Env{
		IsProduction: os.Getenv("IS_PRODUCTION") == "true",
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
	}

	fmt.Println("✅ [ENV] Loaded")
	return e
}
