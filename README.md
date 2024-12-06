# Go Race Condition in Counter Increment

This repository demonstrates a common race condition in Go programs that use goroutines to increment a shared counter.  The program may not produce the expected output of 1000 because multiple goroutines are concurrently accessing and modifying the `counter` variable without proper synchronization.

The `bug.go` file contains the buggy code. The `bugSolution.go` file provides a corrected version using a mutex to prevent race conditions.

## How to Reproduce

1. Clone this repository.
2. Run `go run bug.go` and observe the output; it might not be 1000.
3. Run `go run bugSolution.go` to see the corrected output (1000).

## Solution

The solution involves using a `sync.Mutex` to protect the `counter` variable. This ensures that only one goroutine can access and modify the counter at a time, preventing race conditions.