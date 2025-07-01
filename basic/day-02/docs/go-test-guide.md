# Go Unit Test and Benchmark Test Guide

## Unit Testing in Go
- Unit tests are written in files ending with `_test.go` and functions starting with `Test`.
- Example:
  ```go
  func TestAdd(t *testing.T) {
      result := Add(2, 3)
      if result != 5 {
          t.Errorf("Expected 5, got %d", result)
      }
  }
  ```

### Running Unit Tests
- **PowerShell / Command Prompt / Unix Shell:**
  ```sh
  go test
  go test -v           # Verbose output
  go test ./...        # Run tests in all subdirectories
  ```

## Benchmark Testing in Go
- Benchmark tests are written in functions starting with `Benchmark` and take `*testing.B` as an argument.
- Example:
  ```go
  func BenchmarkAdd(b *testing.B) {
      for i := 0; i < b.N; i++ {
          Add(2, 3)
      }
  }
  ```

### Running Benchmark Tests
- **PowerShell:**
  ```powershell
  go test -bench .
  go test -bench . -benchmem
  go test -bench . -benchtime 10x   # Run each benchmark 10 times
  ```
- **Unix Shell (bash/zsh) or Command Prompt:**
  ```sh
  go test -bench=. 
  go test -bench=. -benchmem
  go test -bench=. -benchtime=10x
  ```

## Coverage and Reports
- **Code coverage:**
  ```sh
  go test -cover
  go test -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```

## Notes
- Always run `go test` in the package directory (not on individual files) for correct context.
- Both unit and benchmark tests can be in the same `_test.go` file.
- Use `-v` for verbose output, `-bench` for benchmarks, and `-benchmem` for memory stats.

---

**References:**
- [Go Testing Package](https://pkg.go.dev/testing)
- [Go Blog: Testing Techniques](https://go.dev/blog/cover)
