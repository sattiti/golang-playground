package main

//go:generate gopherjs build main.go ./libs/moment.inc.js ./libs/jquery.inc.js -o ./output/js/bundle.js -m

import (
  "github.com/gopherjs/gopherjs/js"
)

func now(el *js.Object) {
  period := 1000
  moment := js.Global.Get("moment")
  moment.Call("locale")

  js.Global.Call("setTimeout", func() {
    el.Call("text", moment.New().Call("format", "YYYY-MM-DDTHH:mm:ss (Z)"))
  }, period)

  now(el)
}

func main() {
  j := js.Global.Get("jquery")

  body := j.New("body")
  body.Call("append", "<p id=\"now\"></p>")

  p := j.New("p#now")
  now(p)
}
