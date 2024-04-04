package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  tasks := []string{"task1", "task2", "task3", "task4", "task5"}

  concurrencyLimit := 2
  sem := make(chan struct{}, concurrencyLimit)

  var wg sync.WaitGroup
  for _, task := range tasks {
    wg.Add(1)
    go func(task string) {
      defer wg.Done()
      sem <- struct{}{}
      defer func() { <-sem }()

      // Perform task logic here
      fmt.Println("Processing: ", task)
      time.Sleep(1 * time.Second)
    }(task)
  }

  wg.Wait()
}
