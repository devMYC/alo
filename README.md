```go
package main

import (
        "fmt"
        "sync"

        "github.com/devMYC/alo"
)

func main() {
        var wg sync.WaitGroup
        var mu alo.AtomicLock

        counter := 0

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        defer mu.Unlock()
                        mu.Lock()
                        counter++
                }()
        }

        wg.Wait()
        fmt.Println("Counter =", counter)  // Output: Counter = 10
}
```
