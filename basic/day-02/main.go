package main

import (
	"fmt"

	greeting "github.com/dargoz/day01/exercise03"
	"github.com/dargoz/day02/pkg/controlstructure"
	"github.com/dargoz/day02/pkg/datastructure"
	// "github.com/dargoz/day01/internal/calculator" // internal package, not accessible outside its parent module
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run your Go application.
	// For example, you can print a message or perform some calculations.

	fmt.Printf("%s", greeting.GreetingText)
	// var add = calculator.Add(1, 2)
	// var substract = calculator.Subtract(5, 3)

	datastructure.ArrayExample()

	datastructure.SliceExample()

	fmt.Println("\n--- Map Example ---")
	datastructure.MapExample()
	controlstructure.IfElseExample()
}
