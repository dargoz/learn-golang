package datastructure

import "fmt"

func increaseAmount(account map[string]int, s string, i int) {
	account[s] += i
}

func MapExample() {
	// Maps are unordered collections of key-value pairs
	// Maps are reference types, they point to an underlying data structure
	// Maps can be created using the make function or a map literal
	// Maps can be used to store data in a key-value format
	// Maps can be used to store data in a way that allows for fast lookups
	// Maps can be used to store data in a way that allows for fast updates

	var account = map[string]int{
		"123-ABC": 100000,
		"456-DEF": 125000,
		"789-GHI": 110000,
	}

	fmt.Printf("Account balances: %v\n", account)

	increaseAmount(account, "123-ABC", 5000)
	fmt.Println("Updated account balance for 123-ABC:", account["123-ABC"]) // Increase balance

	delete(account, "456-DEF") // Remove an account
	fmt.Println("Account 456-DEF deleted")
	fmt.Println("Remaining accounts:", account)

	emptyMap := make(map[string]float64) // Create an empty map
	fmt.Printf("Empty map: %v\n", emptyMap)
	emptyMap["USD"] = 15000.0 // Add a key-value pair
	emptyMap["EUR"] = 16000.0 // Add another key-value pair
	fmt.Printf("Updated empty map: %v\n", emptyMap)

	for key, value := range emptyMap {
		fmt.Printf("Currency: %s, Value: %.2f\n", key, value)
	}

	// Merge 2 Map
	anotherMap := map[string]float64{
		"GBP": 18000.0,
		"JPY": 200.0,
	}
	mergeMap(emptyMap, anotherMap)
}

// mergeMap merges all key-value pairs from src into dst.
func mergeMap(dst, src map[string]float64) {
	for k, v := range src {
		dst[k] = v
	}
}
