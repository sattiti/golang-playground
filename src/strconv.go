package main

import (
    "fmt"
    "log"
    "strconv"
)

func main() {
    // parseInt
    // BASE 2, 8, 10, 16 進数
    // BITSIZE 0-64 指定可能
    // strconv.ParseInt(STRING, BASE, BITSIZE)
    var num string = "20"
    fmt.Printf("num: %v\n\n", num)
    fmt.Printf("type:  %T\n", num)
    fmt.Printf("value: %v\n\n", num)

    num10, err := strconv.ParseInt(num, 10, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("type:  %T\n", num10)
    fmt.Printf("value: %v\n\n", num10)

    num8, err := strconv.ParseInt(num, 8, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("type:  %T\n", num8)
    fmt.Printf("value: %v\n\n", num8)

    num16, err := strconv.ParseInt(num, 16, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("type:  %T\n", num16)
    fmt.Printf("value: %v\n\n", num16)
}
