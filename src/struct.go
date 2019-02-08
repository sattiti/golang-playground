package main

import (
    "fmt"
)

type Vertex struct {
    X int
    Y int
}

var (
    p = Vertex{1, 2}
    q = &Vertex{1, 2}
    r = Vertex{X: 1}
    s = Vertex{}
)

// struct tag
// フィールドの宣言の後に文字列でタグを指定することができる。
// タグは key:"value" の形式でスペース区切りで書く。通常、"を含むため`で囲む。
type s struct {
    microsec  uint64 `protobuf:"1"`
    serverIP6 uint64 `protobuf:"2"`
}

// struct nested
type T struct {
    Pilot []struct {
        Name   string
        Age    uint
        Gender string
    }
    Name    string
    Code    string
    Color   uint
    Wrapons []string
    Size    struct {
        Width  int
        Height int
    }
}

func main() {
    // fmt
    // %v  デフォルトフォーマットを適用した値
    // 構造体を出力する際、+フラグ(%+v)を加えるとフィールド名が表示される
    // %#v この値をGo言語の構文で表現する
    // %T  この値の型をGo言語の構文で表現する
    fmt.Printf("Vertex: %#v\n", Vertex{2, 3})

    // new(T) という表現は、ゼロ初期化した(zeroed) T の値をメモリに割り当て、そのポインタを返します。
    // var v *Vertex = new(Vertex)
    _ = new(Vertex)

    v := Vertex{4, 5}
    fmt.Printf("v.X: %d\n", v.X)
    fmt.Printf("v: %+#v\n", v)

    fmt.Println("\n----------")
    fmt.Printf("p -> Has type Vertex: %v\n", p)
    fmt.Printf("q -> Has type *Vertex: %v\n", q)
    fmt.Printf("r -> %#v\n", r)
    fmt.Printf("s -> %#v\n", s)

    // nameless struct
    peter := struct {
        Name        string
        Age         int
        MailAddress string
    }{
        Name:        "Peter",
        Age:         20,
        MailAddress: "peter@peter.com",
    }

    fmt.Println(peter.Name)
    fmt.Println(peter.Age)
    fmt.Println(peter.MailAddress)
    
    // init struct nested
    t := T{
        Pilot: []struct {
            Name   string
            Age    uint
            Gender string
        }{
            {Name: "kira yamato", Age: 18, Gender: "m"},
            {Name: "amuro", Age: 20, Gender: "m"},
        },
        Name:    "Gundam",
        Code:    "RX-78-2",
        Color:   0xffffff,
        Wrapons: []string{"laser beam", "light saver"},
        Size: struct {
            Width  int
            Height int
        }{
            Width:  10,
            Height: 20,
        },
    }

    fmt.Printf("%v\n", t)
}
