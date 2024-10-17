package database

import (
	"context"

	"github.com/devkaare/basic-postgres-demo/config"
	"github.com/jackc/pgx/v5"
)

type Entry struct {
	ID       int
	Username string
	Email    string
	Password string
}

// Connect to Postgress database
func Connect() (*pgx.Conn, error) {
	// Load the database connection url and connect to database
	connection, err := pgx.Connect(context.Background(), config.Config("BASIC_POSTGRES_DEMO_DATABASE_URL")) // Load the database url from .env
	if err != nil {
		return nil, err
	}
	// NOTE: This program doesn't use any concurrency so passing around the database connection is fine!!!
	//defer connection.Close(context.Background())

	if err := connection.Ping(context.Background()); err != nil {
		return nil, err
	}

	// Return database connection
	return connection, nil
}

// Get all entries in the database
func GetEntries() []Entry {
	return []Entry{}
}

// Get entry (if found) with given UserID
func GetEntryByID(id int, connection *pgx.Conn) (*Entry, error) {
	entry := &Entry{}

	// Populate entry with fields returned by database
	err := connection.QueryRow(context.Background(), "select * from entry where id=$1", 1).Scan(
		&entry.ID,
		&entry.Username,
		&entry.Email,
		&entry.Password,
	)
	if err != nil {
		return entry, err
	}

	return entry, nil

}

// Delete entry (if found) with given UserID
func DeleteEntryByID(id int) {
}

// Add entry into the database
func AddEntry(u *Entry) {
}
