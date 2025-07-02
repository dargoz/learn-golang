package concurrency

import (
	"fmt"
	"time"
)

func RunHello() {
	fmt.Println("Hello, from Goroutine!")
}

func ConcurrencyExample() {
	fmt.Println("Entering Concurrency Example")
	go RunHello() // Start a new goroutine
	fmt.Println("Leaving Concurrency Example")
	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}
