package main

import (
    "fmt"
)

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

// func [func名]() 戻り値
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    nextInt2 := intSeq()
    fmt.Println(nextInt2())

    fmt.Println(adder()(2))
}
