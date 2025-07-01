# Go Heap, Garbage Collection, and Debugging GC

## How the Heap Works in Go
- The heap is a region of memory used for dynamic (runtime) allocation.
- In Go, variables are allocated on the heap if they need to outlive the function call or if their address is taken and used elsewhere (e.g., pointers stored in arrays or structs).
- The Go runtime automatically manages heap memory using garbage collection (GC).

### Example from `array_example.go`
```go
saldo1, saldo2 := 1000.0, 2000.0
var saldoArray = [...]*float64{&saldo1, &saldo2}
```
- Here, `saldoArray` is an array of pointers to float64 values. The float64 variables (`saldo1`, `saldo2`) may be allocated on the heap if their addresses are used outside their declaring scope.
- The Go compiler decides whether to allocate on the stack or heap using escape analysis.

## How Garbage Collection (GC) Works
- Go uses a concurrent, tri-color mark-and-sweep garbage collector.
- The GC automatically frees memory that is no longer referenced by any part of the program.
- Heap-allocated objects are tracked by the GC, while stack-allocated objects are not.


## Debugging and Observing GC and Heap Allocation in Go
### 4. Debugging Escape Analysis (Heap Allocation)
You can use the Go CLI with `-gcflags` to see which variables escape to the heap (i.e., are not stack-allocated):
```sh
go build -gcflags="-m" array_example.go
```
This will print messages like `moved to heap` for variables that are allocated on the heap due to escape analysis.


### 1. Enable GC Tracing
You can enable detailed GC logs by running your program with the `GODEBUG` environment variable:
```sh
$env:GODEBUG = "gctrace=1"
go run array_example.go
```
This will print GC events, pause times, and heap sizes to the console.

### 2. Forcing a GC Cycle
You can manually trigger a garbage collection cycle in code:
```go
import "runtime"
...
runtime.GC()
```

### 3. Checking Heap Stats
You can inspect heap and GC stats at runtime:
```go
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("HeapAlloc = %v bytes\n", m.HeapAlloc)
fmt.Printf("NumGC = %v\n", m.NumGC)
```

## Summary Table
| Concept         | Description                                      |
|-----------------|--------------------------------------------------|
| Heap            | Memory for dynamic allocation, managed by GC      |
| Stack           | Memory for local variables, managed automatically |
| GC              | Automatic memory management in Go                 |
| GODEBUG         | Env var to enable GC/heap debug output            |
| runtime.GC()    | Manually trigger a GC cycle                       |
| runtime.MemStats| Inspect heap and GC stats at runtime              |
| -gcflags="-m"   | Show escape analysis and heap allocation info     |

---

**References:**
- [Go Blog: Memory Management](https://blog.golang.org/memory)
- [Go Blog: Garbage Collection](https://blog.golang.org/garbage-collection)
- [Go Runtime Package Docs](https://pkg.go.dev/runtime)
