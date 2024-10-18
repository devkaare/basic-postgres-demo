package main

import (
	"fmt"
	"math"
	"math/rand"
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
			fmt.Println("You chose to GetEntries")
		}
	case "2":
		{
			// Get more input
			username, email, password, err := GetMoreInput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "GetMoreInput failed: %v\n", err)
				os.Exit(1)
			}

			// Gen a unique ID from 0 - 9223372036854775807
			uniqueID := rand.Intn(math.MaxInt)

            //// Create entry 
            //entry := database.Entry {
                //ID: uniqueID,
                //Username: username,
                //Email: email,
                //Password: password,
            //}

			// Add entry to database
            err = database.AddEntry(uniqueID, username, email, password, connection)
			if err != nil {
				fmt.Fprintf(os.Stderr, "AddEntry failed: %v\n", err)
				os.Exit(1)
			}
		}
	case "3":
		{
            // Get ID
			fmt.Println("Enter ID:")
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
		}
	case "4":
		{
			fmt.Println("You chose to DeleteEntryByID")
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
