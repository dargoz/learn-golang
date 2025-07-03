package concurrency

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Before sending to channel\n")
		ch <- "Hello, Channel!"
		fmt.Printf("After sending to channel\n")
	}()

	msg := <-ch
	fmt.Println(msg)

	// Check the number of goroutines
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines: %d\n", numGoroutines)

	// Ensure the channel is closed after use
	close(ch)
}

func GiveMeData(channel chan string) {
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("Data %d", i)
		fmt.Printf("Sending: %s\n", data)
		channel <- data
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
	close(channel) // Close the channel when done
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)

	go GiveMeData(ch)

	data := <-ch

	fmt.Printf("Receiving data from channel: %v\n", data)

	// Check the number of goroutines
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)
}

func OnlySend(channel chan<- string) {
	time.Sleep(2 * time.Second)
	// data := <-channel // this statement will error because channel is send-only
	channel <- "Received data"
}

func OnlyReceive(channel <-chan string) {
	// channel <- "Received data" // this statement will error because channel is receive-only
	data := <-channel
	fmt.Printf("Data received: %s\n", data)
}

func TestChannelAsParamInOut(t *testing.T) {
	ch := make(chan string, 3) // Buffered channel to avoid blocking

	go OnlySend(ch)
	go OnlyReceive(ch)

	// Wait for goroutines to finish
	time.Sleep(3 * time.Second)

}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			// Simulate sending data to the channel
			ch <- "Send data to " + strconv.Itoa(i)
			time.Sleep(500 * time.Millisecond) // Simulate some work
		}
		close(ch) // Close the channel when done
	}()

	for data := range ch {
		fmt.Printf("Received: %s\n", data)
	}

	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)
}
