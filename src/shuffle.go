package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    ary := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    fmt.Printf("before: %v\n", ary)

    for i := len(ary) - 1; i > 0; i-- {
        r := rand.Intn(i + 1)
        tmp := ary[i]
        ary[i] = ary[r]
        ary[r] = tmp
    }

    fmt.Printf("after:  %v\n", ary)
}
