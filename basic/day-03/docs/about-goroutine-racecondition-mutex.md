# Goroutines, Race Conditions, and Synchronization in Go

## What is a Goroutine?
A goroutine is a lightweight thread managed by the Go runtime. Goroutines allow you to run functions concurrently, making it easy to build concurrent and parallel programs in Go.

- **Start a goroutine:**
  ```go
  go myFunction()
  ```

## Race Conditions
A race condition occurs when two or more goroutines access shared data at the same time, and at least one of them modifies the data. This can lead to unpredictable results and bugs.

**Example (unsafe, race condition):**
```go
var counter int
for i := 0; i < 100; i++ {
    go func() {
        counter++ // not safe: race condition!
    }()
}
```

## Using Mutex to Prevent Race Conditions
A mutex (mutual exclusion lock) is a synchronization primitive that allows only one goroutine to access a critical section of code at a time. This is useful when multiple goroutines need to read or write shared data.

**Why use a mutex with goroutines?**
- Goroutines run concurrently and may access shared variables at the same time, causing race conditions.
- A mutex ensures that only one goroutine can access the shared data at a time, keeping the data consistent.
- Mutex allows you to keep the benefits of concurrency (speed, responsiveness) while protecting shared resources.

**Example:**
```go
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
```

**Partial Lock Example:**
```go
var mu sync.Mutex
counterX := 0
counterY := 0

increment := func() {
    mu.Lock()
    counterX++ // Only counterX is protected by the mutex
    mu.Unlock()
    counterY++ // counterY is NOT protected and may have race conditions
}

for i := 0; i < 10000; i++ {
    go increment()
}
```
In this example, only `counterX++` is protected by the mutex. `counterY++` is outside the lock and is not protected, so it can still have race conditions if accessed by multiple goroutines.

**When to use a mutex:**
- When multiple goroutines need to read/write the same variable or data structure.
- When you want to avoid race conditions but still want concurrent execution.

**Alternatives:**

## Read/Write Mutex (sync.RWMutex)
A `sync.RWMutex` is a special kind of mutex that allows multiple goroutines to read a shared resource at the same time, but only one goroutine to write. While a write lock is held, no other goroutine can read or write.

**When to use RWMutex:**
- When you have data that is read frequently but only occasionally written to.
- Allows higher concurrency for read-heavy workloads.

**How it works:**
- Use `RLock()`/`RUnlock()` for read-only access (multiple readers allowed).
- Use `Lock()`/`Unlock()` for write access (exclusive, blocks all readers and writers).

**Example scenario:**
Suppose you have a shared map that is read by many goroutines, but only occasionally updated.

```go
var rwmu sync.RWMutex
sharedMap := make(map[string]int)

// Reader goroutine
go func() {
    rwmu.RLock()
    value := sharedMap["foo"] // safe concurrent read
    rwmu.RUnlock()
    fmt.Println("Read value:", value)
}()

// Writer goroutine
go func() {
    rwmu.Lock()
    sharedMap["foo"] = 42 // safe concurrent write
    rwmu.Unlock()
    fmt.Println("Wrote value")
}()
```

## RWMutex Scenario Example
Below is a comprehensive example demonstrating the behavior of sync.RWMutex with multiple readers and writers, and how read and write locks interact:

```go
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
```

This example shows:
- Multiple writers (using Lock/Unlock) incrementing a counter.
- Multiple readers (using RLock/RUnlock) reading the counter concurrently.
- A writer holding the write lock, blocking all readers and writers.
- A reader holding the read lock, allowing other readers but blocking writers.

**Key points:**
- Multiple goroutines can hold the read lock at the same time.
- Only one goroutine can hold the write lock, and it blocks all readers and other writers.
- Always use `RLock`/`RUnlock` for read-only sections, and `Lock`/`Unlock` for write sections.

## Deadlock Scenario
A deadlock occurs when goroutines are waiting for each other and none can proceed. In Go, a common deadlock happens when you try to send or receive on a channel but there is no corresponding receiver or sender.

**Example (deadlock):**
```go
ch := make(chan int)
ch <- 1 // This will cause a deadlock! No goroutine is receiving.
```

**Example (deadlock on receive):**
```go
ch := make(chan int)
value := <-ch // Deadlock! No goroutine is sending.
```

**How to avoid deadlocks:**
- Always ensure there is a goroutine ready to receive when you send, and vice versa.
- Use buffered channels if you need to send without an immediate receiver (but be careful not to overflow the buffer).
- Close channels properly when done sending, especially when using range to receive.

## References
- [Go Blog: Goroutines](https://go.dev/blog/goroutines)
- [Go Blog: Concurrency is not parallelism](https://go.dev/blog/concurrency-is-not-parallelism)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
