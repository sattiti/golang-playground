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
    num := 5
    min := num
    max := num * 2

    // int
    // range in 5...9
    for i := 0; i < max; i++ {
        fmt.Println(rand.Intn(max-min) + min)
    }

    // float64
    // range in 5...9
    for i := 0; i < max; i++ {
        fmt.Println(rand.Float64()*float64(max-min) + float64(min))
    }
}
