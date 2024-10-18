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

// Listen and handle input 1 - 4
func GetInput() {
	fmt.Println("NOTE: Type help to refresh list of choices")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
		os.Exit(1)
	}

	switch input {
	case "help":
		PrintOptions()
	case "1":
		{
			// Get all entries
			entries, err := database.GetEntries(connection)
			if err != nil {
				fmt.Fprintf(os.Stderr, "GetEntries failed: %v\n", err)
			}

			// Print all entries
			for _, entry := range entries {
				PrintEntry(entry)
			}

			// Print total amount of entries
			total := len(entries)
			fmt.Println("Total amount of entries:", total)

			fmt.Println("Successfully got all entries!")
		}
	case "2":
		{
			// Get entry details
			username, email, password, err := GetMoreInput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "GetMoreInput failed: %v\n", err)
				os.Exit(1)
			}

			// Gen a unique ID
			uniqueID := rand.IntN(2147483647) // Postgres integer can hold -2147483648 to +2147483647

			//entry := database.Entry {
			//ID: uniqueID,
			//Username: username,
			//Email: email,
			//Password: password,
			//}

			// Add entry to database
			if err = database.AddEntry(uniqueID, username, email, password, connection); err != nil {
				fmt.Fprintf(os.Stderr, "AddEntry failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Successfully created entry!")
		}
	case "3":
		{
			// Get ID
			fmt.Println("Get, Enter ID:")
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

			// Print result
			PrintEntry(entry)

			fmt.Println("Successfully got entry!")
		}
	case "4":
		{
			// Get ID
			fmt.Println("Delete, Enter ID:")
			var inputID int

			if _, err := fmt.Scanln(&inputID); err != nil {
				fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
				os.Exit(1)
			}

			// Get entry
			if err := database.DeleteEntryByID(inputID, connection); err != nil {
				fmt.Fprintf(os.Stderr, "GetEntryByID failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Successfully deleted entry!")
		}
	default:
		fmt.Println("Invalid option")
		os.Exit(1)
	}
}

// Listen and handle input for creating a new entry
func GetMoreInput() (string, string, string, error) {
	var username, email, password string

	// Username
	fmt.Println("Enter a username:")
	if _, err := fmt.Scanln(&username); err != nil {
		return "", "", "", err
	}

	// Email
	fmt.Println("Enter a email:")
	if _, err := fmt.Scanln(&email); err != nil {
		return "", "", "", err
	}

	// Password
	fmt.Println("Enter a password:")
	if _, err := fmt.Scanln(&password); err != nil {
		return "", "", "", err
	}

	return username, email, password, nil
}

// Print entry in a custom format
func PrintEntry(entry database.Entry) {
	//
	fmt.Printf("Found data for ID: %d successfully!\n", entry.ID)
	fmt.Printf("ID:             %d\n", entry.ID)
	fmt.Printf("Username:       %s\n", entry.Username)
	fmt.Printf("Email:          %s\n", entry.Email)
	fmt.Printf("Password:       %s\n", entry.Password)
}
