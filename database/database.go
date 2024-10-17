package database

import (
	"context"
	"fmt"
	//"os"

	"github.com/devkaare/basic-postgres-demo/config"
	"github.com/jackc/pgx/v5"
)

type Entry struct {
	ID       int
	Username string
	Email    string
	Password string
}

var DB *pgx.Conn

// Connect to Postgress database.
func Connect() error {
	// Load the database URL.
	connURL := config.Config("BASIC_POSTGRES_DEMO_DATABASE_URL")

	DB, err := pgx.Connect(context.Background(), connURL)
	if err != nil {
		return err
		//fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//os.Exit(1)
	}
	defer DB.Close(context.Background())

	//CreateEntryTable()

	fmt.Println("Connection Opend to Database")
	var greeting string
	err = DB.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
		//fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		//os.Exit(1)
	}

	fmt.Println(greeting)
	return nil
}

//create table if not exists entries (
//id integer primary key,
//username text unique not null,
//email text not null,
//password text not null
//)

// Get all entries in the database.
func GetEntries() []Entry {
	return []Entry{}
}

// Get entry (if found) with given UserID.
func GetEntryByID(id int) Entry {
	return Entry{}
}

// Delete entry (if found) with given UserID.
func DeleteEntryByID(id int) {
}

// Add entry into the database.
func AddEntry(u *Entry) {
}
