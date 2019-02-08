package main

import(
  "fmt"
  "reflect"
  "math"
  "time"
)

// function {{{
// 複数の戻り値を返すことができる。
func sum(a int, b int)(int){
  return a + b
}

// closures
func adder() func(int) int{
  sum := 0
  return func(x int) int{
    sum += x
    return sum
  }
}

func say(s string){
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func fibonacci(c, quit chan int){
  x, y := 0, 1
  for{
    select {
    case c <- x:
      x, y = y, x + y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}
// }}}

// struct {{{
type player struct{
  name string
  gender string
  age uint
}
// }}}

// method {{{
// method は特別な receiver 引数を関数に取る。
// receiver は func キーワードとメソッド名の間に自身の引数リストで表現する。
type Vertex struct{
  X, Y float64
}
func (a Vertex) AbsMethod() float64{
  return math.Sqrt(a.X * a.X + a.Y * a.Y)
}

func AbsNormal(v Vertex) float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (a *Vertex) Scale(f float64) {
  a.X = a.X * f
  a.Y = a.Y * f
}

v := vertex{X: 10, Y: 20}
v.AbsNormal()
// }}}

// interface {{{
// 異なる構造体でも同じインタフェース型で処理ができる。

// Empty Interface
// 空のインターフェースは任意の型の値を保持できる。
// var i interface{}
// s := i.(string)
// n := i.(uint)

type greeting interface{
  greet()
}
type ja struct{}
type en struct{}

func (u ja) greet(){
  fmt.Println("こんにちは。")
}

func (u en) greet(){
  fmt.Println("Good afternoon.")
}
// }}}

func main(){
  // data types {{{
  fmt.Println("\n-- DataTypes --")

  // bool

  // string

  // nil

  // int
  // int8
  // int16
  // int32 (alias: rune. unicode のコードポイントを表す)
  // int64

  // uint
  // uint8 (alias: byte) 
  // uint16
  // uint32
  // uint64
  // uintptr

  // float32
  // float64

  // complex64
  // complex128

  // Array types
  // Slice types
  // Struct types
  // Pointer types
  // Function types
  // Interface types
  // Map types
  // Channel types

  // type cast
  var ci int = 40
  fmt.Printf("cast: %f\n", float64(ci))
  // }}}

  // var {{{
  fmt.Println("\n-- Var --")

  // _
  // 戻り値は使用しない場合、_ に代入する。

  // 小文字はパッケージ内のみ参照可。
  // 大文字はパッケージ外の参照可。(exported name)

  // 型指定
  var x int = 1

  // 型推測
  y := 2
  // }}}

  // const, enum {{{
  fmt.Println("\n-- Const --")

  const(
    A = 1
    B = 2
    C = 3
  )

  // enum
  const(
    _ = iota // blank identifier を使って、0 を破棄する。
    D        // 1
    E        // 2
    F        // 3
  )
  // }}}

  // Array {{{
  fmt.Println("\n-- Array  --")

  var ary[]int = []int{1, 2, 3, 4, 5}
  fmt.Printf("ary len = %d\n", len(ary))
  fmt.Printf("ary capacity = %d\n", cap(ary))
  // }}}

  // slices {{{
  fmt.Println("\n-- Slices --")

  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  fmt.Println(s)

  // make
  a := make([]int, 5)  // len(a)=5
  a = append(a, 2)
  fmt.Println("A =", a)

  // range
  for i, v := range a{
    fmt.Println("order:", i)
    fmt.Println("value:", v)
  }
  // }}}

  // map (hash) {{{
  fmt.Println("\n-- Map --")

  // m = make(map[string]float64)
  // m := map[keyType]valueType{"key": "value"}
  mydict := map[string]int{
    "john": 100, 
    "mary": 200,
  }
  fmt.Println(mydict)
  fmt.Println(mydict["john"])
  // }}}

  // pointer {{{
  fmt.Println("\n-- Pointer --")

  var p *int = &x
  fmt.Printf("x pointer = %p\n", p)
  fmt.Println("x value   =", *p)
  // }}}

  // switch {{{
  fmt.Println("\n-- Switch --")

  switch i := 0; i {
  case 0:
    fmt.Println(i)
  case 1, 3:
    fallthrough // 自動 break。次の case を実行したい場合、fallthrough を書く。
  default:
  }

  // }}}

  // for {{{
  fmt.Println("\n-- For --")

  for i := 0; i < 10; i++ {
    // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
    fmt.Println(i)

    // return の直前に 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 と逆に出力される
    defer fmt.Println(i)
  }

  i1 := 10
  for ; i1 < 20; {
    fmt.Println(i1)
    i1 += 1
  }

  i2 := 20
  for i2 < 30 {
    fmt.Println(i2)
    i2 += 1
  }

  // while(1)
  // for {
  //   ...
  // }
  /// }}}

  // if {{{
  fmt.Println("\n-- If --")

  if x > 1 {
    fmt.Printf("x = %d\n", x)
  } else if y > 2 {
    fmt.Printf("y = %d\n", y)
  } else {
    fmt.Println(A)
    fmt.Println(reflect.TypeOf(x))
    fmt.Printf("hello world.\n")
    fmt.Printf("sum func %d\n", sum(1, 3))
  }
  // }}}

  // defer {{{
  fmt.Println("\n-- Defer --")

  // return するまで遅延させる。
  // hello friend world.
  defer fmt.Println("\n-- Defer --\nworld")
  fmt.Println("hello")
  fmt.Println("friend")
  // }}}

  // closures {{{
  fmt.Println("\n-- Closures --")
  fmt.Println(adder()(2))
  // }}}

  // Method {{{
  fmt.Println("\n-- Method --")
  v := Vertex{3, 5}

  // Method を使わない場合の書き方。
  fmt.Println(AbsNormal(v))

  // Point method
  // (&v).Scale(10) として解釈する。
  v.Scale(10)

  // (*pp).Scale(20) として解釈する。
  pp := &v
  pp.Scale(20)

  // Method はオブジェクトっぽい書き方ができる。
  // v.Abs() (receiver).(func)
  fmt.Println(v.AbsMethod())
  // }}}

  // Interface {{{
  fmt.Println("\n-- Interface --")

  g := []greeting{ja{}, en{}}
  for _, u := range g {
    u.greet()
  }
  // }}}

  // Goroutines {{{
  // 並行処理。 runtime に管理される軽量なスレッド
  fmt.Println("\n-- Goroutines --")
  go say("gogo")
  say("nono")
  // }}}

  // Channels {{{
  // goroutines 間でデータのやりとりするための機能。
  // パイプのようなもの。
  fmt.Println("\n-- Channels --")

  ch := make(chan string)

  // 2. 2秒後に task 100 が処理されたから、task100 finished. が出力される。
  go func(ch chan string){
    fmt.Println(<-ch)
    time.Sleep(time.Millisecond * 100)
    fmt.Println("task100 finished.")
    ch <- "task100 result."
  }(ch)

  // 1. 並列処理により、task0-9 を処理
  go func(){
    var max int = 10
    for i := 0; i < max; i++{
      fmt.Printf("task%d finished.\n", i)
    }
    ch <- "task 0-9 result"
  }()

  // 3. task100 result を ch に渡し、最後の出力のとこでプリントアウト。
  time.Sleep(time.Millisecond * 2000)
  fmt.Println(<-ch)

  fmt.Println("\n-- Select --")
  fibCh   := make(chan int)
  fibQuit := make(chan int)

  go func(){
    for i := 0; i < 10; i++ {
      fmt.Println(<-fibCh)
    }
    fibQuit <- 0
  }()

  fibonacci(fibCh, fibQuit)

  // }}}
}
