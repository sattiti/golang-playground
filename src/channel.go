package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func say(msg string) {
    fmt.Println(msg)
}

func main() {
    // Goroutines
    // 並行処理。 runtime に管理される軽量なスレッド
    fmt.Println("\n-- Goroutines --")
    go say("gogo")
    say("nono")

    // Channels
    // goroutines 間でデータのやりとりするための機能。
    // パイプのようなもの。
    fmt.Println("\n-- Channels --")

    ch := make(chan string)

    // 2. 2秒後に task 100 が処理されたから、task100 finished. が出力される。
    go func(ch chan string) {
        fmt.Println(<-ch)
        time.Sleep(time.Millisecond * 100)
        fmt.Println("task100 finished.")
        ch <- "task100 result."
    }(ch)

    // 1. 並列処理により、task0-9 を処理
    go func() {
        var max int = 10
        for i := 0; i < max; i++ {
            fmt.Printf("task%d finished.\n", i)
        }
        ch <- "task 0-9 result"
    }()

    // 3. task100 result を ch に渡し、最後の出力のとこでプリントアウト。
    time.Sleep(time.Millisecond * 2000)
    fmt.Println(<-ch)

    fmt.Println("\n-- Select --")
    fibCh := make(chan int)
    fibQuit := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-fibCh)
        }
        fibQuit <- 0
    }()

    fibonacci(fibCh, fibQuit)
}
