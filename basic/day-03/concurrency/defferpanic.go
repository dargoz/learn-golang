package concurrency

import "fmt"

func DeferPanicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting DeferPanicExample")
	panic("This is a panic example")
	fmt.Println("This line will not be executed")
}
