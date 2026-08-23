package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURL string

	ElasticAddrs           string
	ElasticUser            string
	ElasticPassword        string
	ElasticCertFingerprint string
}

func New() *Env {
	godotenv.Load()

	e := &Env{
		DatabaseURL:            os.Getenv("DATABASE_URL"),
		ElasticAddrs:           os.Getenv("ELASTIC_ADDRS"),
		ElasticUser:            os.Getenv("ELASTIC_USER"),
		ElasticPassword:        os.Getenv("ELASTIC_PASSWORD"),
		ElasticCertFingerprint: os.Getenv("ELASTIC_CERT_FINGERPRINT"),
	}

	fmt.Println("✅ [ENV] Loaded")
	return e
}
