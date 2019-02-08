package main

import (
    "fmt"
)

type Player struct{
  name string
  age uint
  gender string
}

func main() {
    fmt.Println("\n-- Array  --")

    ary := []string{"apple", "banana", "orange", "lemon", "melon", "cherry", "blue berry"}
    // var ary [6]string = {"apple", "banana", "orange", "lemon", "cherry", "blue berry"}
    // make(type, length, capacity)
    // ary := make([]string, 6, 10)

    // length 長さ
    fmt.Printf("ary len      = %d\n", len(ary))

    // capacity 容量
    fmt.Printf("ary capacity = %d\n", cap(ary))

    for i, val := range ary {
        fmt.Printf("%d: %s\n", i, val)
    }

    // slice
    fmt.Println("\n-- Slices --")
    var s []string = ary[1:4]
    fmt.Println(s)

    s = append(s, "grape")
    fmt.Println(s)
    
    // pointer inside
    arr := []*Player{
        &Player{name: "aaa", age: 20, gender: "m"},
        &Player{name: "bbb", age: 20, gender: "m"},
        &Player{name: "ccc", age: 20, gender: "m"},
        &Player{name: "ddf", age: 20, gender: "m"},
    }
    
    // pointer array
    var *p *[]*Player = &arr
    
    for _, o := range *p{
        fmt.Printf("%#v\n", o)
    }
}
