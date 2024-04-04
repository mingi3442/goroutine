package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  tasks := []string{"task1", "task2", "task3", "task4", "task5"}
  numWorkers := 3

  tasksCh := make(chan string)

  var wg sync.WaitGroup
  for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for task := range tasksCh {
        // Perform task logic here
        fmt.Println("Processing:", task)
        time.Sleep(1 * time.Second)
      }
    }()
  }

  for _, task := range tasks {
    tasksCh <- task
  }
  close(tasksCh)

  wg.Wait()
}
