package main

// import は以下のように書くこともできる。
// import "fmt"
// import "math"
import (
    "fmt"
    "math"
)

// 初期化
func init() {
}

func main() {
    fmt.Println("hello world.")
    fmt.Printf("%s\n", "printf hello world")

    // 最初の文字が [大文字] で始まる場合は外部へ公開される。
    // math.Pi 外部公開される
    fmt.Println(math.Pi)

    // math.pi 外部公開されない
    //fmt.Println(math.pi)
}
