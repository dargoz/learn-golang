package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func DisplayNumber(n int) {
	fmt.Printf("Displaying number: %d\n", n)
}

func TestManyGoroutines(t *testing.T) {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	// For testing purposes, we can limit to 1 core to avoid overwhelming the output
	runtime.GOMAXPROCS(1)
	start := time.Now()
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	totalTime := time.Since(start)
	time.Sleep(5 * time.Second) // Wait for goroutines to finish
	fmt.Printf("Total time taken: %s\n", totalTime)
}
