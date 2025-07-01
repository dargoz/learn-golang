package main

import "fmt"

func main() {
	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"Reading", "Cycling", "Cooking"}
	// Output (print) that array in the command line.
	fmt.Println("Hobbies:", hobbies)
	// 2) Also output more data about that array:
	// - The first element (standalone)
	fmt.Println(hobbies[0]) // First hobby
	// - The second and third element combined as a new list
	fmt.Println("Hobbies 2 and 3:", hobbies[1:3]) // Second and third hobbies
	// 3) Create a slice based on the first element that contains the first and second elements.
	// Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2]      // First way: using slicing syntax
	slice2 := make([]string, 2) // Second way: using make function
	slice2[0] = hobbies[0]
	slice2[1] = hobbies[1]
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	// 4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
	slice1 = hobbies[1:3] // Reslicing to contain second and last elements
	fmt.Println("Resliced Slice 1:", slice1)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Build a cli application"}
	fmt.Println("Course Goals:", courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Contribute to open source")
	fmt.Println("Updated Course Goals:", courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products).
	// Then add a third product to the existing list of products.
	type Product struct {
		Title string
		ID    int
		Price float64
	}
	products := []Product{
		{Title: "Laptop", ID: 1, Price: 999.99},
		{Title: "Mouse", ID: 2, Price: 25.50},
	}
	products = append(products, Product{Title: "Keyboard", ID: 3, Price: 45.00})
	fmt.Println("Products:", products)
}
