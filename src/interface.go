package main

import (
    "fmt"
)

type greeting interface {
    greet()
}
type ja struct{}
type en struct{}

func (u ja) greet() {
    fmt.Println("こんにちは。")
}

func (u en) greet() {
    fmt.Println("Good afternoon.")
}

type Gundam interface {
    GetName() string
}

// Pilot
type Pilot struct {
    Name string
}

// type を使って名前付けを行わなくとも定義できる
// func Func(v interface{GetName() string}){}

// implements Gundam interface
func (p *Pilot) GetName() string {
    return p.Name
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    pilot1 := Pilot{"kira yamato"}
    fmt.Printf("%#T\n", pilot1)
    fmt.Printf("%s\n", pilot1.GetName())

    fmt.Println("\n-- Interface --")

    g := []greeting{ja{}, en{}}
    for _, u := range g {
        u.greet()
    }

    // empty interface
    var a interface{}
    fmt.Println(a)
    describe(a)
}
