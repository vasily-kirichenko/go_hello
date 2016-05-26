package main

import (
    "fmt"
    "time"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}


func main() {
    m := make(map[int]int)

    start := time.Now()
    for i := 0; i < 50000000; i++ {
        m[i] = 0
    }
    fmt.Printf("Insertion: %s", time.Since(start))
    
    start = time.Now()
    acc := 0
    for i := 0; i < 50000000; i++ {
        acc += m[i] % 10
    }
    fmt.Printf("Lookups: %s", time.Since(start))

} 