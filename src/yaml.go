package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "log"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
    A string
    B struct {
        RenamedC int   `yaml:"c"`
        D        []int `yaml:",flow"`
    }
}

func run(f func()) {
    fmt.Printf("----\n")
    f()
    fmt.Printf("----\n\n")
}

func main() {
    run(func() {
        t := T{}
        if err := yaml.Unmarshal([]byte(data), &t); err != nil {
            log.Fatalf("error: %v", err)
        }

        fmt.Printf("%v\n\n", t)
    })
}
