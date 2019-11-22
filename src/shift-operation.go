package main

import (
  "fmt"
  "math"
)

func main() {
  len := 5
  n1 := 8

  fmt.Println()
  fmt.Println("<<")
  fmt.Println("a = b * 2 の n倍")
  for i := 1; i < len; i++ {
    fmt.Printf("%d: %d * %d = %d\n", i, n1, int(math.Pow(2, float64(i))), n1<<i)
  }

  fmt.Println()
  fmt.Println(">>")
  fmt.Println("a = b * 0.5 の n倍")
  for i := 1; i < len; i++ {
    fmt.Printf("%d: %d * %f = %f\n", i, n1, 0.5/float64(i), float64(n1>>i))
  }
}
