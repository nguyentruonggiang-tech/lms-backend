package ent_client

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"lms-worker/ent"
	_ "lms-worker/ent/runtime"
	"lms-worker/internal/common/env"
)

func New(e *env.Env) *ent.Client {
	client, err := ent.Open("mysql", e.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ [ENT] failed opening connection to mysql: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ [ENT] failed creating schema resources: %v", err)
	}

	fmt.Println("✅ [ENT] Connected successfully")
	return client
}
