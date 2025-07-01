package main

import (
	"fmt"
	"time"
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run your Go application.
	var firstName string
	var lastName = "Gozali"
	fmt.Print("Enter your name: ")
	fmt.Scanln(&firstName) // Wait for input to keep the console open
	fmt.Print("Enter your birth year: ")
	var birthYear int
	fmt.Scanln(&birthYear)
	currentYear := time.Now().Year()
	var age int = currentYear - birthYear

	var emoji rune = '😊' // Unicode for a smiley face emoji
	fmt.Printf("Hello %s %s %c, you're now %d years old\n", firstName, lastName, emoji, age)
}
