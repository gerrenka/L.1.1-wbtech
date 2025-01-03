package main

import (
    "fmt"
    "sync"
)

func main() {
    var sharedMap sync.Map
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            sharedMap.Store(fmt.Sprintf("key_%d", i), i*i)
        }(i)
    }

    wg.Wait()

    sharedMap.Range(func(key, value any) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true
    })
}
