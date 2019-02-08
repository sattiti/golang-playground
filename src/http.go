package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    _ "strings"
)

func main() {
    url := "https://www.apple.com/"
    respo, err := http.Get(url)

    if err != nil {
        panic(err)
    }
    defer respo.Body.Close()

    b, err := ioutil.ReadAll(respo.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(b))
}
