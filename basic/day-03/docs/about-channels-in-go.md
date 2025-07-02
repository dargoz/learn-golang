# Channels in Go

> For goroutine, race condition, mutex, and deadlock explanations, see `about-goroutine-racecondition-mutex.md` in this folder.

## What is a Channel?
A channel in Go is a built-in type that enables communication between goroutines. It allows one goroutine to send data to another goroutine in a safe and synchronized way. Channels are strongly typed, meaning you specify the type of data the channel will carry.

- **Create a channel:**
  ```go
  ch := make(chan int) // channel of int
  ```
- **Send data:**
  ```go
  ch <- 42
  ```
- **Receive data:**
  ```go
  value := <-ch
  ```

## How Channels Work
Channels provide a way for goroutines to synchronize execution and communicate data. When you send data on a channel, the sending goroutine is blocked until another goroutine receives from that channel (unless the channel is buffered and not full).

## Example: Basic Channel Usage
```go
func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello, Channel!"
    }()
    msg := <-ch
    fmt.Println(msg) // Output: Hello, Channel!
}
```


## Synchronous (Unbuffered) vs Buffered Channels
- **Synchronous (unbuffered):** The sender waits until the receiver is ready to receive. This is the default channel type in Go.
  ```go
  ch := make(chan int) // unbuffered channel
  ```
- **Buffered:** You can specify a buffer size. The sender only blocks when the buffer is full, and the receiver only blocks when the buffer is empty.
  ```go
  ch := make(chan int, 2) // buffered channel with capacity 2
  ch <- 1
  ch <- 2 // does not block yet
  // ch <- 3 // would block until a value is received
  ```

## Channel Direction: OnlySend and OnlyReceive as Parameters
You can restrict a function parameter to only send or only receive on a channel:

- **Send-only channel:**
  ```go
  func OnlySend(ch chan<- string) {
      ch <- "data" // can only send
  }
  ```
- **Receive-only channel:**
  ```go
  func OnlyReceive(ch <-chan string) {
      data := <-ch // can only receive
      fmt.Println(data)
  }
  ```
This helps clarify intent and prevents misuse inside the function.

## Iterating with Range on a Channel
You can use `range` to receive values from a channel until it is closed:
```go
ch := make(chan string)
go func() {
    for i := 0; i < 5; i++ {
        ch <- fmt.Sprintf("Data %d", i)
    }
    close(ch)
}()
for data := range ch {
    fmt.Println("Received:", data)
}
```

A deadlock occurs when goroutines are waiting for each other and none can proceed. In Go, a common deadlock happens when you try to send or receive on a channel but there is no corresponding receiver or sender.

## Deadlock Scenario
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

A race condition can occur if multiple goroutines access and modify shared data without proper synchronization. Channels can help avoid race conditions by providing a safe way to communicate between goroutines, but if you share variables directly, you may still have race conditions.

**Example (safe with channel):**
```go
ch := make(chan string)
go func() {
    for i := 0; i < 100; i++ {
        ch <- fmt.Sprintf("Data %d", i)
    }
    close(ch)
}()
for data := range ch {
    fmt.Println("Received:", data)
}
```

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
```

**When to use a mutex:**
- When multiple goroutines need to read/write the same variable or data structure.
- When you want to avoid race conditions but still want concurrent execution.

**Alternatives:**
- Use channels to coordinate access to shared data (channel as a lock or to pass ownership of data).
- Make the code sequential (no concurrency), but this removes the benefits of goroutines.

## Closing a Channel
When no more values will be sent, close the channel:
```go
close(ch)
```
Receivers can check if a channel is closed using the second value:
```go
value, ok := <-ch
if !ok {
    // channel is closed
}
```

## Channel vs PubSub
- **Channel:** Point-to-point communication. Each value sent is received by only one receiver. Built-in to Go.
- **PubSub (Publish/Subscribe):** Many-to-many communication. Multiple subscribers can receive the same message. Not built-in to Go, but can be implemented using channels and goroutines.

## Example: Channel is Not PubSub
```go
ch := make(chan string)
go func() { ch <- "data" }()
msg1 := <-ch // Only one receiver gets the message
msg2 := <-ch // This will block unless another value is sent
```

## When to Use Channels
- To coordinate work between goroutines
- To safely share data between goroutines
- To signal completion or pass results

## References
- [Go Blog: Channels](https://go.dev/blog/pipelines)
- [Effective Go: Channels](https://go.dev/doc/effective_go#channels)
- [Go by Example: Channels](https://gobyexample.com/channels)
