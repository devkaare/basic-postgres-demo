package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/devkaare/basic-postgres-demo/database"
)

var connection database.Connection

func main() {
	// Required so the compiler doesn't complain when calling database.Connect
	var err error

	// Connect to database and get a connection
	connection, err = database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connecting to database failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	// Print options
	PrintOptions()

	// Start listening for input
	for {
		GetInput()
	}
}

// Print available options 1 - 4
func PrintOptions() {
	fmt.Print(
		"What would you like to do?\n",
		"1: Get all entries\n",
		"2: Create new entry\n",
		"3: Get entry by ID\n",
		"4: Delete entry by ID\n",
		"Enter the number displayed beside your choice:\n",
	)
}

// Print entry in a custom format
func PrintEntry(entry database.Entry) {
	fmt.Printf("ID:             %d\n", entry.ID)
	fmt.Printf("Username:       %s\n", entry.Username)
	fmt.Printf("Email:          %s\n", entry.Email)
	fmt.Printf("Password:       %s\n", entry.Password)
}

// Listen and handle input 1 - 4
func GetInput() {
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
		os.Exit(1)
	}

	// Check if input was help early so it doesn't need to check all the switch cases
	if input == "help" {
		PrintOptions()
		return
	}

	// Check and execute selected option
	switch input {
	case "1":
		{
			fmt.Println("SELECTED: Get all entries")

			// Get all entries
			entries, err := database.GetEntries(connection)
			if err != nil {
				fmt.Fprintf(os.Stderr, "GetEntries failed: %v\n", err)
			}

			// Print all entries
			for _, entry := range entries {
				fmt.Printf("%+v\n", entry)
			}

			// Print total amount of entries
			fmt.Println("INFO: Total amount of entries:", len(entries))

			fmt.Println("SUCCESS: Successfully got all entries!")
		}
	case "2":
		{
			fmt.Println("SELECTED: Create new entry")

			// Create database entry struct
			entry := database.Entry{}

			// Populate entry with ID, Username, Email and Password
			// Get ID
			entry.ID = rand.Int32()

			// Get username
			fmt.Println("INPUT: Enter a username:")
			if _, err := fmt.Scanln(&entry.Username); err != nil {
				os.Exit(1)
			}

			// Get email
			fmt.Println("INPUT: Enter a email:")
			if _, err := fmt.Scanln(&entry.Email); err != nil {
				os.Exit(1)
			}

			// Get password
			fmt.Println("INPUT: Enter a password:")
			if _, err := fmt.Scanln(&entry.Password); err != nil {
				fmt.Fprintf(os.Stderr, "CreateEntry failed: %v\n", err)
				os.Exit(1)
			}

			// Add entry to database
			if err := database.AddEntry(entry, connection); err != nil {
				fmt.Fprintf(os.Stderr, "AddEntry failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("INFO: New ID:", entry.ID)

			fmt.Println("SUCCESS: Successfully created entry!")
		}
	case "3":
		{
			fmt.Println("SELECTED: Get entry by ID")

			// Get ID
			fmt.Println("INPUT: Enter ID:")
			var inputID int

			if _, err := fmt.Scanln(&inputID); err != nil {
				fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
				os.Exit(1)
			}

			// Get entry
			entry, err := database.GetEntryByID(inputID, connection) // Pass the database connection
			if err != nil {
				fmt.Fprintf(os.Stderr, "GetEntryByID failed: %v\n", err)
				os.Exit(1)
			}

			// Print entry
			fmt.Printf("%+v\n", entry)

			fmt.Println("SUCCESS: Successfully got entry!")
		}
	case "4":
		{
			fmt.Println("SELECTED: Delete entry by ID")

			// Get ID
			fmt.Println("INPUT: Enter ID:")
			var inputID int

			if _, err := fmt.Scanln(&inputID); err != nil {
				fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
				os.Exit(1)
			}

			// Get entry
			if err := database.DeleteEntryByID(inputID, connection); err != nil {
				fmt.Fprintf(os.Stderr, "DeleteEntryByID failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("SUCCESS: Successfully deleted entry!")
		}
	default:
		{
			fmt.Println("INVALID: Invalid option entered!")
		}
	}
}
