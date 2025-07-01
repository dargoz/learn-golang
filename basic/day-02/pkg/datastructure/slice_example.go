package datastructure

import "fmt"

func SliceExample() {
	prices := []float64{100.0, 200.0, 300.0, 400.0, 500.0}
	discounts(prices, 0.1)
	for i, price := range prices {
		fmt.Printf("Price after discount %v : %v", i+1, price)
	}

	productNames := [4]string{"Laptop"}
	productNames[2] = "E-Reader" // Update the third product name
	fmt.Printf("\nProduct Names: %v", productNames)

	// Slices are more flexible than arrays, they can grow and shrink
	// Slices are a reference type, they point to an underlying array
	sliceOfProducts := []string{"Laptop", "Smartphone", "Tablet"}
	sliceOfProducts = append(sliceOfProducts, "Smartwatch") // Add a new product
	fmt.Printf("\nSlice of Products: %v", sliceOfProducts)
	// Slices can be created from arrays
	sliceFromArray := productNames[1:3] // Create a slice from the array
	fmt.Printf("\nSlice from Array: %v", sliceFromArray)
	// Slices can be created from other slices
	sliceFromSlice := sliceOfProducts[1:3] // Create a slice from the
	fmt.Printf("\nSlice from Slice: %v", sliceFromSlice)
	// Slices can be created from a slice with a capacity
	sliceWithCapacity := make([]string, 2, 5) // Create a slice with a length of 2 and a capacity of 5
	sliceWithCapacity[0], sliceWithCapacity[1] = "Laptop", "Smartphone"
	fmt.Printf("\nSlice with Capacity: %v, Length: %d, Capacity: %d", sliceWithCapacity, len(sliceWithCapacity), cap(sliceWithCapacity))

}

func discounts(prices []float64, f float64) {
	for i, price := range prices {
		prices[i] = price * (1 - f)
	}
}
