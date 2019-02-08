package main

import (
    "fmt"
)

type Vertex struct {
    Lat, Long float64
}

func main() {
    m := map[string]Vertex{
        // 型の省略可能
        // "home": {40.68433, -74.39967},
        // "google": {37.42202, -122.08408},
        "home":   Vertex{40.68433, -74.39967},
        "google": Vertex{37.42202, -122.08408},
    }

    // map[key の型] value の型
    d := make(map[string]Vertex)
    d["home"] = Vertex{40.68433, -74.39967}

    // キーが存在するかどうか
    // ok が true であればある
    // elem, ok := m[key]
    fmt.Println(m)
    fmt.Println(d["home"])

    // 要素の削除は:
    delete(m, "home")
    fmt.Println(m)
}
