package main

import (
    "fmt"
    "time"
    "math/rand"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}

func main() {
    var source [50000000]int
    for i := 0; i < len(source); i++ {
        source[i] = rand.Int()
    }
    m := make(map[int]int)
    start := time.Now()
    for x := range source {
        m[x] = 0
    }
    fmt.Printf("Insertion: %s", time.Since(start))
    
    start = time.Now()
    acc := 0
    for x := range source {
        acc += m[x] % 10
    }
    fmt.Printf("Lookups: %s", time.Since(start))

} 