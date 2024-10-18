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

// Required so main.go can access *pgx.Conn
type Connection *pgx.Conn

// Connect to Postgress database
func Connect() (Connection, error) {
    var connection Connection
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
func GetEntries(connection *pgx.Conn) ([]Entry, error){
    rows, err := connection.Query(context.Background(), "select * from entry")
    if err != nil {
        return []Entry{}, err
    }
    defer rows.Close()

    var entries []Entry 

    for rows.Next() {
        var entry Entry

        if err := rows.Scan(&entry.ID, &entry.Username, &entry.Email, &entry.Password); err != nil {
            return entries, err
        }

        entries = append(entries, entry)
    }
    if err = rows.Err(); err != nil {
        return entries, err
    }

    return []Entry{}, nil
}

// Get entry (if found) with given UserID
func GetEntryByID(id int, connection *pgx.Conn) (Entry, error) {
    var entry Entry

	// Populate entry with fields returned by database
	err := connection.QueryRow(context.Background(), "select * from entry where id = $1", id).Scan(
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
func DeleteEntryByID(id int, connection *pgx.Conn) error {
    _, err := connection.Exec(context.Background(), "delete from entry where id = $1", id)
    // I don't need result for anything, hence _ above
    if err != nil {
        return err
    }

    return nil
}

// Add entry into the database
func AddEntry(id int, username, email, password string, connection *pgx.Conn) error {
    _, err := connection.Exec(context.Background(), "insert into entry (id, username, email, password) values ($1, $2, $3, $4)",
        id,
        username,
        email,
        password,
    )
    if err != nil {
        return err
    }

    return nil
}
