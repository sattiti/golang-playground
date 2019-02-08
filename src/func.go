package main

import (
    "fmt"
    "time"
)

// func [func 名]([引数] [引数の型]) ([戻り値の型])
// func add(x, y int)int{
func add(x int, y int) int {
    z := 3
    return x + y + z
}

// 複数の戻り値
func swap(x, y string) (string, string) {
    // return だけで構わない。
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var say = func(s string) {
    fmt.Println(s)
}

func sayAndWait(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    // nameless func
    defer (func() {
        sayAndWait("exec an Nameless func.")
    })()

    fmt.Println(add(1, 2))
    fmt.Println(swap("a", "b"))
    fmt.Println(split(10))

    fmt.Printf("%#v\n", say)
    say("hello")
}
