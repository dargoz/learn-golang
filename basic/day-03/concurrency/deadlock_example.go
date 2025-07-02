package concurrency

import "fmt"

func DeadlockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	// This will cause a deadlock because both goroutines are waiting for each other to finish
	<-ch1
	<-ch2

	fmt.Println("This line will never be reached due to deadlock")
}
