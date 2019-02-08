package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func run(f func()) {
    fmt.Printf("----\n")
    f()
    fmt.Printf("----\n\n")
}

func main() {
    run(func() {
        out, err := exec.Command("ls", "-la").Output()
        log.Printf("cmd is Running...\n")

        if err != nil {
            log.Fatal(err)
        }

        log.Printf("Finished: %v\n", string(out))
        os.Exit(0)
    })
}
