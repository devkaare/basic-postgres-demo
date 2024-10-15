package database

import (
	"context"
	"fmt"
	"os"

	"github.com/devkaare/basic-postgres-demo/config"
	"github.com/jackc/pgx/v5"
)

type Entry struct {
	UserID   int
	Username string
	Email    string
	Password string
}

// Connect to Postgress DB.
func Connect() {
	// Load the DB connection URL
	connURL := config.Config("BASIC_POSTGRES_DEMO_DATABASE_URL")

	// Open DB connection
	conn, err := pgx.Connect(context.Background(), connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

// Fetch all entries in the DB.
func FetchEntries() []Entry {
	return []Entry{}
}

// Fetch Entry (if found) with given UserID.
func FetchEntryByID(id int) Entry {
	return Entry{}
}

// Delete Entry (if found) with given UserID.
// Returns true if Entry was deleted, and false if not.
func DeleteEntryByID(id int) bool {
	return false
}

// Insert Entry into the DB.
// Returns true if Entry was inserted, and false if not.
func InsertEntry(u *Entry) bool {
	return false
}
