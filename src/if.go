package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := 0
    b := 1
    if a > b {
        fmt.Println(a)
    } else {
        fmt.Println(b)
    }

    x := 10
    y := 20
    if x > 1 {
        fmt.Printf("x = %d\n", x)
    } else if y > 2 {
        fmt.Printf("y = %d\n", y)
    } else {
        fmt.Println(reflect.TypeOf(x))
        fmt.Printf("hello world.\n")
    }

}
