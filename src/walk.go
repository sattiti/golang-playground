package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    dist := "./"
    err := filepath.Walk(dist, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            panic(err)
        }

        fmt.Printf("%s\n", path)
        return nil
    })

    if err != nil {
        panic(err)
    }
}
