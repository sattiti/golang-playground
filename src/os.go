package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func run(f func()) {
    fmt.Println("----")
    f()
    fmt.Println("----\n")
}

func main() {
    // env
    run(func() {
        for _, val := range os.Environ() {
            env := strings.Split(val, "=")
            fmt.Printf("[%s]\n%s\n\n", env[0], env[1])
        }
    })

    // Getenv
    run(func() {
        fmt.Println(1)
    })

    // cd
    run(func() {
        pwd, err := os.Getwd()
        if err != nil {
            panic(err)
        }
        fmt.Printf("pwd: %s\n", pwd)

        err = os.Chdir("../")
        if err != nil {
            panic(err)
        }
        fmt.Println("cd ..")

        pwd, err = os.Getwd()
        if err != nil {
            panic(err)
        }
        fmt.Printf("pwd: %s\n", pwd)
    })

    // expandEnv
    run(func() {
        fmt.Println(os.ExpandEnv("${USER} lives in ${HOME}."))
    })

    // expand
    run(func() {
        mapper := func(ph string) string {
            switch ph {
            case "MORNING":
                return "MILK"
            case "AFTERNOON":
                return "COLA"
            case "EVENING":
                return "BEER"
            }
            return ""
        }
        fmt.Println(os.Expand("DRINK ${MORNING} in the morning.", mapper))
        fmt.Println(os.Expand("DRINK ${AFTERNOON} in the afternoon.", mapper))
        fmt.Println(os.Expand("DRINK ${EVENING} in the evening.", mapper))
    })

    run(func() {
        // user id
        fmt.Printf("uid:       %d\n", os.Geteuid())

        // hostname
        hostname, err := os.Hostname()
        if err != nil {
            panic(err)
        }
        fmt.Printf("hostname:  %s\n", hostname)

        // return the process id.
        fmt.Printf("pid:       %d\n", os.Getpid())

        // returns the process id of the caller's parent.
        fmt.Printf("ppid:      %d\n", os.Getppid())

        // returns the underlying system's memory page size.
        fmt.Printf("page size: %d\n", os.Getpagesize())

        // group id
        fmt.Printf("gid:       %d\n", os.Getegid())
        // groups
        groups, err := os.Getgroups()
        if err != nil {
            panic(err)
        }
        fmt.Printf("groups:    %v\n", groups)

    })

    // exists
    run(func() {
        filename, err := filepath.Abs("./index.html")
        if err != nil {
            panic(err)
        }
        c, err := os.Stat(filename)
        fmt.Println(c)
        fmt.Printf("IsExist:      %v\n", os.IsExist(err))
        fmt.Printf("IsNotExist:   %v\n", os.IsNotExist(err))
        fmt.Printf("IsPermission: %v\n", os.IsPermission(err))
    })

    // FileInfo
    run(func() {
        filename := "./index.html"
        f, err := os.Create(filename)
        if err != nil {
            panic(err)
        }

        stat, err := f.Stat()
        if err != nil {
            panic(err)
        }
        fmt.Printf("Name:    %v\n", stat.Name())
        fmt.Printf("Size:    %v\n", stat.Size())
        fmt.Printf("Mode:    %v\n", stat.Mode())
        fmt.Printf("ModTime: %v\n", stat.ModTime())
        fmt.Printf("IsDir:   %v\n", stat.IsDir())
        fmt.Printf("Sys:     %t\n", stat.Sys())

        fmt.Printf("rm: %s\n", filename)
        os.Remove(filename)
    })

    // mkdir
    run(func() {
        dir1 := "./public_html/"
        dir2 := "./secure_html/about/"
        var mode os.FileMode = 0775

        // mkdir
        fmt.Printf("mkdir: %v\n", dir1)
        if err := os.Mkdir(dir1, mode); err != nil {
            panic(err)
        }

        // mkdir -p
        fmt.Printf("mkdir -p: %v\n", dir2)
        if err := os.MkdirAll(dir2, mode); err != nil {
            panic(err)
        }

        // rm
        fmt.Printf("rm: %v\n", dir1)
        if err := os.Remove(dir1); err != nil {
            panic(err)
        }

        // rm -fr
        fmt.Printf("rm -rf: %v\n", dir2)
        if err := os.RemoveAll(dir2); err != nil {
            panic(err)
        }
    })

    // exit
    os.Exit(0)
}
