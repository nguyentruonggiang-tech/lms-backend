package ent_client

import (
	"context"
	"fmt"
	"log"

	"lms-api/ent"
	_ "lms-api/ent/runtime"
	"lms-api/internal/common/env"

	_ "github.com/go-sql-driver/mysql"
)

func New(e *env.Env) *ent.Client {
	client, err := ent.Open("mysql", e.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ [ENT] failed opening connection to mysql: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ [ENT] failed creating schema resources: %v", err)
	}

	fmt.Println("✅ [ENT] Connection To MySQL Successfully")
	return client
}
