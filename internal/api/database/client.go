package database

import (
	"context"
	"log"

	"github.com/gmarcha/goswagger-ent-app/ent"
)

func Connect() *ent.Client {

	client, err := ent.Open("postgres", "host=192.168.0.20 port=26668 dbname=api_db user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to psql: %v", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
