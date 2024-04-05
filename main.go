package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())

  go func() {
    // Simulate a long-running task
    time.Sleep(3 * time.Second)
    cancel() // Cancel the context after 3 second
  }()

  select {
  case <-ctx.Done():
    fmt.Println("Task canceled")
  }

}
