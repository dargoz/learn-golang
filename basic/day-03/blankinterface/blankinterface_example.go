package blankinterface

import "fmt"

func BlankInterfaceExample() {
	// Blank interface can hold any type
	var i interface{}

	// Assigning an integer to the blank interface
	i = 42
	println("Integer value:", i.(int))

	// Assigning a string to the blank interface
	i = "Hello, World!"
	println("String value:", i.(string))

	// Assigning a float to the blank interface
	i = 3.14
	println("Float value:", i.(float64))
	// Assigning a boolean to the blank interface
	i = true
	println("Boolean value:", i.(bool))
	// Assigning a struct to the blank interface
	type Person struct {
		Name string
		Age  int
	}
	i = Person{Name: "Alice", Age: 30}
	fmt.Printf("Struct value: %+v\n", i.(Person))
	// Check blank interface type
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println("Type is int with value:", v)
	case string:
		fmt.Println("Type is string with value:", v)
	case float64:
		fmt.Println("Type is float64 with value:", v)
	case bool:
		fmt.Println("Type is bool with value:", v)
	default:
		fmt.Println("Unknown type")
	}
}
