package main

import (
  "fmt"
  "sync"
)

var count int
var mutex sync.Mutex

func main() {
  var wg sync.WaitGroup
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go increment(&wg)
  }
  wg.Wait()

  fmt.Println("Final count:", count)
}

func increment(wg *sync.WaitGroup) {
  defer wg.Done()
  mutex.Lock()
  defer mutex.Unlock()

  count++
}
