package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Michael19s/go_first_api/ent"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewEntClient() (*ent.Client, error) {
	//root:my-secret-pw@tcp(localhost:3306)/testbd?parseTime=True
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	//dsn := "root:my-secret-pw@tcp(localhost:3306)/testbd?parseTime=True"

	client, err := ent.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
