package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/devkaare/basic-postgres-demo/database"
)

func main() {
	// Connect to database and get a connection
	connection, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connecting to database failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	// Pass database connection
	entry, err := database.GetEntryByID(1, connection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetEntryByID failed: %v\n", err)
		os.Exit(1)
	}

	PrintEntryInFormat(entry)

	//PrintAndListen()
}

// TODO: Add proper functionality to the different cases

// Print all available options and listen for input
func PrintAndListen() {
	var input string
	fmt.Print("What would you like to do?\n1: Fetch all entries\n2: Fetch entry by ID\n3: Delete user by ID\n4: Create NEW user\nEnter the number associated with the action you want to preform:\n")
	fmt.Scanln(&input)
	switch input {
	case "1":
		// FetchEntries
		fmt.Println("You chose to FetchEntries")
	case "2":
		// FetchEntryByID
		fmt.Println("You chose to FetchEntryByID")
	case "3":
		// DeleteEntryByID
		fmt.Println("You chose to DeleteEntryByID")
	case "4":
		// InsertEntry
		fmt.Println("You chose to InsertEntry")
	}
}

// Generate a random ID from 0 - 9223372036854775807
func GenID() int {
	return rand.Intn(math.MaxInt)
}

// Print entry in a custom format
func PrintEntryInFormat(entry *database.Entry) {
    fmt.Printf("Found data for ID: %d successfully!\n", entry.ID)
	fmt.Printf("ID:             %d\n", entry.ID)
	fmt.Printf("Username:       %s\n", entry.Username)
	fmt.Printf("Email:          %s\n", entry.Email)
    fmt.Printf("Password:       %s\n", entry.Password)
}
