package main

import (
	"fmt"

	greeting "github.com/dargoz/day01/exercise03"
	"github.com/dargoz/day01/internal/calculator"
	"github.com/dargoz/day01/models"
	"github.com/dargoz/day01/pointer"
)

func main() {

	// This is a placeholder for the main function.
	// You can add your code here to run your Go application.
	// For example, you can print a message or perform some calculations.

	fmt.Printf("%s", greeting.GreetingText)
	var add = calculator.Add(1, 2)
	var substract = calculator.Subtract(5, 3)

	p, q := calculator.AddAndSubtract(1, 2)
	fmt.Printf("1 + 2 = %d\n", add)
	fmt.Printf("5 - 3 = %d\n", substract)
	fmt.Printf("1 + 2 = %d, 1 - 2 = %d\n", p, q)

	x := 10
	pointer.WithoutPointer(x)
	fmt.Printf("Value of x after without pointer manipulation: %d\n", x)

	pointer.WithPointer(&x)
	fmt.Printf("Value of x after pointer manipulation: %d\n", x)

	var p1 models.Person
	// p1.name = "John Doe" // This will not work because name is unexported
	p1.Age = 30
	p1.Birth = 1993
	p1.Address = "123 Main St"
	p1.Phone = "123-456-7890"
	fmt.Printf("Person: %s, Age: %d, Birth: %d, Address: %s, Phone: %s\n", p1.GetName(), p1.Age, p1.Birth, p1.Address, p1.Phone)
	fmt.Printf("Person: %+v\n", p1)

	p2 := models.Person{
		// Name:    "Jane Doe" // This will not work because name is unexported
		Age:     25,
		Birth:   1998,
		Address: "456 Elm St",
		Phone:   "987-654-3210",
	}
	fmt.Printf("Person: %+v\n", p2)

	p3 := models.NewPerson("Alice Smith", 28, 1995, "789 Oak St", "555-123-4567")

	fmt.Printf("Person: %+v\n", p3)
	p3.ChangeName("Alice Johnson")
	fmt.Printf("Person after name change: %+v\n", p3)

	p4 := &models.Person{
		// Name:    "Bob Brown" // This will not work because name is unexported
		Age:     35,
		Birth:   1988,
		Address: "321 Pine St",
		Phone:   "555-987-6543",
	}
	fmt.Printf("Person: %+v\n", p4)

}
