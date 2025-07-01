package main

import "fmt"

func main() {
	var pi float64 = 3.14
	var radius float64 = 5.0

	var area float64 = pi * radius * radius
	var circumference float64 = 2 * pi * radius

	fmt.Printf("from a radius of %v, the area of the circle is: %.2f\n", radius, area)
	fmt.Printf("from a radius of %v, the circumference of the circle is: %.2f\n", radius, circumference)
}
