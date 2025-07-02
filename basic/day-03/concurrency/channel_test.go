package concurrency

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
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

func TestRaceCondition(t *testing.T) {
	counter := 0

	increment := func() {
		counter++
	}

	for i := 0; i < 10000; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to finish

	fmt.Printf("Final counter value: %d\n", counter)
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)
}

func TestMutexRaceCondition(t *testing.T) {
	var mu sync.Mutex
	counter := 0

	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		counter++
	}

	for i := 0; i < 10000; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to finish

	fmt.Printf("Final counter value: %d\n", counter)
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)
}

func TestRWMutex(t *testing.T) {
	var mu sync.RWMutex
	counter := 0

	// Writer: increments the counter (exclusive lock)
	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		counter++
	}

	// Reader: reads the counter (shared lock)
	readCounter := func() int {
		mu.RLock()
		defer mu.RUnlock()
		return counter
	}

	// Simulate multiple writers
	for i := 0; i < 10000; i++ {
		go increment()
	}

	// Simulate multiple readers
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				val := readCounter()
				fmt.Printf("[Reader %d] Counter value: %d\n", id, val)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// Simulate that a writer blocks all readers and writers
	go func() {
		mu.Lock()
		fmt.Println("[Writer] Holding write lock for 200ms (no readers/writers allowed)")
		time.Sleep(200 * time.Millisecond)
		mu.Unlock()
		fmt.Println("[Writer] Released write lock")
	}()

	// Simulate that multiple readers can read at the same time
	go func() {
		mu.RLock()
		fmt.Println("[Reader] Holding read lock for 100ms (other readers allowed, writers blocked)")
		time.Sleep(100 * time.Millisecond)
		mu.RUnlock()
		fmt.Println("[Reader] Released read lock")
	}()

	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	fmt.Printf("Final counter value: %d\n", readCounter())
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)
}
