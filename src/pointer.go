package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

type Point struct {
    X, Y float64
}

// func [pointer receiver] [Method Name] [return type]
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (p Point) Scale(f float64) {
    p.X = p.X * f
    p.Y = p.Y * f
}

func (p Point) Abs() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
    a := Vertex{2, 3}

    //q := &a
    var q *Vertex = &a
    q.X = 9
    fmt.Println(a.X)

    // pointer
    var v *Vertex = &Vertex{3, 4}
    v.Scale(5)
    fmt.Printf("Pointer あり:\n%#v, %f\n", v, v.Abs())

    // without pointer
    var p Point = Point{3, 4}
    p.Scale(5)
    fmt.Printf("Pointer なし:\n%#v, %f\n\n", p, p.Abs())

    x := 10
    var p2 *int = &x
    fmt.Printf("x pointer = %p\n", p2)
    fmt.Println("x value   =", *p2)
}
