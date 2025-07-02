# Parallelism vs Concurrency, Process vs Thread in Go

## Concurrency vs Parallelism

| Aspect        | Concurrency                                      | Parallelism                                 |
|---------------|--------------------------------------------------|---------------------------------------------|
| Definition    | Many tasks make progress at the same time        | Many tasks run at the exact same time       |
| How           | Tasks are interleaved (may not run simultaneously)| Tasks run simultaneously on multiple cores  |
| Go Example    | Goroutines scheduled by Go runtime               | Goroutines running on multiple CPU cores    |
| Hardware Need | Works on single or multi-core                    | Requires multi-core CPU                     |
| Analogy       | Chef cooking several dishes, switching between them | Several chefs cooking several dishes at once |

- **Concurrency** is about structuring a program to handle many tasks that can be interleaved (mixed together in time), even if only one is running at any instant.
- **Parallelism** is about doing many things at the exact same time, using multiple CPU cores.

## Interleaving
- Interleaving means the CPU switches between tasks, so their steps are mixed together in time. This gives the appearance of simultaneous progress, even if only one task is running at a time (on a single core).


## Concurrency vs Sequential Execution

**Sequential execution:** Tasks are performed one after another—each must finish before the next starts. The program cannot make progress on more than one task at a time.

**Concurrency (interleaved execution):** Multiple tasks can make progress independently, and their steps are mixed together in time. The program can switch between tasks, so if one is waiting (e.g., for I/O), another can run. This allows overall progress on several tasks, even on a single CPU core.

**Example:**

- Sequential:  Task A runs to completion → then Task B runs to completion.
- Concurrent:  Task A runs a bit → Task B runs a bit → Task A continues → Task B continues, etc.

Concurrency allows a program to remain responsive and handle multiple activities at once, even on a single CPU core, by switching between tasks. Sequential code cannot do this—it blocks on each task until it finishes.

---
## Process vs Thread

| Aspect        | Process                                         | Thread                                      |
|---------------|-------------------------------------------------|---------------------------------------------|
| Definition    | Independent program with its own memory space   | Lightweight unit of execution within a process |
| Memory        | Separate address space                          | Shared address space with other threads      |
| Communication | Inter-process communication (IPC) needed        | Can communicate via shared memory           |
| Overhead      | Higher (context switch, memory)                 | Lower (faster context switch)               |
| Failure       | Crash is isolated                               | Crash can affect other threads              |

- **Process:** An independent running program. In Go, each go run or go build creates a new process.
- **Thread:** A unit of execution within a process. Go's goroutines are managed by the Go runtime and multiplexed onto threads.

## Go and Goroutines
- Goroutines are lightweight, managed by the Go runtime, and can be scheduled concurrently or in parallel.
- By default, Go uses all available CPU cores (since Go 1.5), so goroutines can run in parallel if hardware allows.

## Summary Table
| Concept      | Go Example/Behavior                              |
|--------------|--------------------------------------------------|
| Concurrency  | Many goroutines, may be interleaved              |
| Parallelism  | Many goroutines, run on multiple CPU cores       |
| Process      | OS process, separate memory                      |
| Thread       | OS thread, shared memory; goroutines multiplexed |

## Best Practices
- Use goroutines for concurrency; Go will handle parallelism if possible.
- Understand the difference: concurrency is about structure, parallelism is about execution.
- Use channels for safe communication between goroutines.

---

**References:**
- [Go Blog: Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
- [Go Blog: Goroutines](https://blog.golang.org/goroutines)
- [Effective Go: Goroutines](https://golang.org/doc/effective_go#goroutines)
