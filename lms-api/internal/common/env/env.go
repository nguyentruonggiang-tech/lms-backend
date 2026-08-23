package env

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	IsProduction bool
	Host         string
	Port         string
	DatabaseURL  string

	SecretAccessToken    string
	ExpiresAtAccessToken time.Duration

	SecretRefreshToken    string
	ExpiresAtRefreshToken time.Duration

	RedisAddr string
	RedisPass string

	ElasticAddrs           string
	ElasticUser            string
	ElasticPassword        string
	ElasticCertFingerprint string

	RabbitMQURL string
}

func New() *Env {
	godotenv.Load()

	e := &Env{
		IsProduction:          os.Getenv("IS_PRODUCTION") == "true",
		Host:                  os.Getenv("HOST"),
		Port:                  os.Getenv("PORT"),
		DatabaseURL:           os.Getenv("DATABASE_URL"),
		SecretAccessToken:     os.Getenv("SECRET_ACCESS_TOKEN"),
		ExpiresAtAccessToken:  parseDuration("EXPIRES_AT_ACCESS_TOKEN"),
		SecretRefreshToken:    os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiresAtRefreshToken: parseDuration("EXPIRES_AT_REFRESH_TOKEN"),
		RedisAddr:              os.Getenv("REDIS_ADDR"),
		RedisPass:              os.Getenv("REDIS_PASS"),
		ElasticAddrs:           os.Getenv("ELASTIC_ADDRS"),
		ElasticUser:            os.Getenv("ELASTIC_USER"),
		ElasticPassword:        os.Getenv("ELASTIC_PASSWORD"),
		ElasticCertFingerprint: os.Getenv("ELASTIC_CERT_FINGERPRINT"),
		RabbitMQURL:            os.Getenv("RABBIT_MQ_URL"),
	}

	fmt.Println("✅ [ENV] Loaded")
	return e
}

func parseDuration(key string) time.Duration {
	d, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		log.Fatalf("invalid duration for %s: %v", key, err)
	}
	return d
}
