package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

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

func TestAntiDeadLockUsingWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Hello from Goroutine!"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		msg := <-ch
		fmt.Println(msg)
	}()

	wg.Wait() // Wait for all goroutines to finish

	// Check the number of goroutines
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of goroutines after processing: %d\n", numGoroutines)

	close(ch) // Close the channel after use

}
