package main

import (
    "fmt"
)

func main() {
    // return するまで遅延させる。
    // hello friend world.
    defer fmt.Println("\n-- Defer --\nworld")
    fmt.Println("hello")
    fmt.Println("friend")
}
