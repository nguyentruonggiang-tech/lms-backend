package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURL string
}

func New() *Env {
	godotenv.Load()

	e := &Env{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	fmt.Println("✅ [ENV] Loaded")
	return e
}
