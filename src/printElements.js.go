package main

import (
  "github.com/gopherjs/gopherjs/js"
)

func main() {
  d := js.Global.Get("document")
  els := d.Call("querySelectorAll", "body *").Interface().(map[string]interface{})
  for _, el := range els {
    println(el)
  }
  // .Call("addEventListener", "click", func() {
  //   println("hello world")
  // })
}
