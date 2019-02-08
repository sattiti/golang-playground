package main

import (
    "fmt"
)

// for 文
func forStat() int {
    sum := 0

    // () 不要
    for i := 0; i < 10; i++ {
        sum += i
    }
    return sum
}

// while と同じような機能
func whileLike() int {
    sum := 1

    for sum < 100 {
        sum += sum
    }
    return sum
}

func rangeAge(ages []int) {
    for i, v := range ages {
        fmt.Printf("i == %d, v == %d\n", i, v)
    }
}

func rangeName(names []string) {
    // index や値を "_" へ代入することで破棄することができます。
    // for _, v := range names{}
    // for i, _ := range names{}
    for i, v := range names {
        fmt.Printf("i == %d, v == %s\n", i, v)
    }
}

func main() {
    fmt.Printf("forStat(): %d\n", forStat())

    fmt.Println("----")
    fmt.Printf("whileLike(): %d\n", whileLike())

    fmt.Println("----")
    rangeAge([]int{20, 12, 34, 32, 41, 53, 69})

    fmt.Println("----")
    rangeName([]string{"Apple", "Orange", "Banana", "Lemon"})

    for i := 0; i < 10; i++ {
        fmt.Println(i)

        // 9, 8, 7 .. 0
        defer fmt.Println(i)
    }
}
