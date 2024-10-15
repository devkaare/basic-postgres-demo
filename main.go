package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/devkaare/basic-postgres-demo/database"
)

func main() {
	// Connect to Postgress DB
	database.Connect()

	//PrintAndListen()
}

// TODO: Add proper functionality to the different cases

// Print all available options and listen for input.
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
