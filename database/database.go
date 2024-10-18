package database

import (
	"context"
	"errors"

	"github.com/devkaare/basic-postgres-demo/config"
	"github.com/jackc/pgx/v5"
)

type Entry struct {
	ID       int32 // Postgres integer can hold -2147483648 to +2147483647 so using int32
	Username string
	Email    string
	Password string
}

// Required so main.go can access *pgx.Conn
var (
	Connection  *pgx.Conn
	ErrNotFound = errors.New("entry with provided id was not found")
)

// Connect to Postgress database
func Connect() (*pgx.Conn, error) {
	// Load the database connection url and connect to database
	connection, err := pgx.Connect(context.Background(), config.Config("BASIC_POSTGRES_DEMO_DATABASE_URL")) // Load the database url from .env
	if err != nil {
		return nil, err
	}
	// NOTE: This program doesn't use any concurrency so passing around the database connection is fine!!!
	//defer connection.Close(context.Background())

	// Return database connection
	return connection, nil
}

// Get all entries in the database
func GetEntries(connection *pgx.Conn) ([]Entry, error) {
	var entries []Entry

	// Send query to database and get a row with results
	rows, err := connection.Query(context.Background(), "select * from entry")
	if err != nil {
		return entries, err
	}
	defer rows.Close()

	// Loop over rows, create entry and append it to entries
	for rows.Next() {
		var entry Entry

		if err := rows.Scan(&entry.ID, &entry.Username, &entry.Email, &entry.Password); err != nil {
			return entries, err
		}

		entries = append(entries, entry)
	}

	// Check rows error after its closed
	if err = rows.Err(); err != nil {
		return entries, err
	}

	return entries, nil
}

// Get entry with given ID
func GetEntryByID(id int32, connection *pgx.Conn) (Entry, error) {
	var entry Entry

	// Populate entry with fields returned by database
	err := connection.QueryRow(context.Background(), "select * from entry where id = $1", id).Scan(&entry.ID, &entry.Username, &entry.Email, &entry.Password)
	if err != nil && err != pgx.ErrNoRows { // If err is not ErrNoRows return err
		return entry, err
	}

	// If err is ErrNoRows return ErrNotFound
	if err != nil {
		return entry, ErrNotFound
	}

	return entry, nil
}

// Delete entry with given ID
func DeleteEntryByID(id int, connection *pgx.Conn) error {
	result, err := connection.Exec(context.Background(), "delete from entry where id = $1", id)
	if err != nil {
		return err
	}

	// If rows affacted is NOT 1 return ErrNotFound
	if result.RowsAffected() != 1 {
		return ErrNotFound
	}

	return nil
}

// Add entry to the database
func AddEntry(entry Entry, connection *pgx.Conn) error {
	if _, err := connection.Exec(context.Background(), "insert into entry (id, username, email, password) values ($1, $2, $3, $4)", entry.ID, entry.Username, entry.Email, entry.Password); err != nil {
		return err
	}

	return nil
}
