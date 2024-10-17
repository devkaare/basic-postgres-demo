package main

import (
	"fmt"
	//"math"
	//"math/rand"
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

	PrintWithFormat(entry)

	//GetInput()
}

// TODO: Add proper functionality to the different cases

// Print all available options and listen for input
func GetInput() {
	var input string
	fmt.Print(
		"What would you like to do?\n",
		"1: Get all entries\n",
		"2: Create new entry\n",
		"3: Get entry by ID\n",
		"4: Delete entry by ID\n",
		"Enter the number displayed beside your choice:\n",
	)
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Fprintf(os.Stderr, "GetInput failed: %v\n", err)
		os.Exit(1)
	}

	switch input {
	case "1":
		fmt.Println("You chose to GetEntries")
	case "2":
		fmt.Println("You chose to GetMoreInput")
		username, email, password, err := GetMoreInput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "GetMoreInput failed: %v\n", err)
			os.Exit(1)
		}
        
        // Add entry to database
        fmt.Println(username, email, password)
	case "3":
		fmt.Println("You chose to GetEntryByID")
	case "4":
		fmt.Println("You chose to DeleteEntryByID")
	}
}

// Generate a random ID from 0 - 9223372036854775807
//func GenID() int {
//return rand.Intn(math.MaxInt)
//}

// Print required fields and listen for input
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
func PrintWithFormat(entry *database.Entry) {
	//
	fmt.Printf("Found data for ID: %d successfully!\n", entry.ID)
	fmt.Printf("ID:             %d\n", entry.ID)
	fmt.Printf("Username:       %s\n", entry.Username)
	fmt.Printf("Email:          %s\n", entry.Email)
	fmt.Printf("Password:       %s\n", entry.Password)
}
