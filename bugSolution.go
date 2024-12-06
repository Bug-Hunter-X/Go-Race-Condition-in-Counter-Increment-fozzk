```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        var mu sync.Mutex
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock()
                        counter++
                        mu.Unlock()
                }()
        }
        wg.Wait()
        fmt.Println(counter)
}
```