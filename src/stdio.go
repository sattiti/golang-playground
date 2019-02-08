package main

import (
    "fmt"
    "os"
)

func main() {
    // 標準入力取得
    // os.Args
    // os.Args[1]
    for i, v := range os.Args {
        fmt.Printf("%v: %v\n", i, v)
    }
}
