package main

import (
    "github.com/eknkc/amber"
    "os"
)

func main() {
    file := "./index.amber"
    out := "./index.html"

    tpl, err := amber.CompileFile(file, amber.DefaultOptions)

    if err != nil {
        panic(err)
    }

    f, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0755)

    if err != nil {
        panic(err)
    }

    data := map[string]string{
        "Title": "title",
        "H1":    "hello world",
    }

    tpl.Execute(f, data)
}
