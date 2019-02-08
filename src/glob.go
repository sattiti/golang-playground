package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    // glob
    pat := "./**/*.html"
    files, err := filepath.Glob(pat)

    if err != nil {
        panic(err)
    }

    for _, f := range files {
        fmt.Println(f)
    }
}
