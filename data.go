package main

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
)

func GetEvent(db *surrealdb.DB) {

}

func setupDB() *surrealdb.DB {
	url := os.Getenv("SURREALDB_URL")
	if url == "" {
		url = "ws://localhost:8000/rpc"
	}

	db, err := surrealdb.New(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating database: %v\n", err)
		os.Exit(1)
	}
}
