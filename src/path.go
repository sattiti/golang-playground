package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {
    paths := []string{
        "a/c",
        "a//c",
        "a/c/.",
        "a/c/b/..",
        "/../a/b/../././/c",
        "",
        "/ab/c/de/fg/h/i/index.html",
    }

    for _, v := range paths {
        fmt.Printf("= [%s]\n", v)

        abs, _ := filepath.Abs(v)
        fmt.Printf("Abs:        %s\n", abs)
        rel, _ := filepath.Rel(v, "/")
        fmt.Printf("Rel:        %s\n", rel)
        fmt.Printf("Base:       %s\n", filepath.Base(v))
        fmt.Printf("Clean:      %s\n", filepath.Clean(v))
        fmt.Printf("Dir:        %s\n", filepath.Dir(v))
        fmt.Printf("Ext:        %s\n", filepath.Ext(v))
        fmt.Printf("VolumeName: %s\n", filepath.VolumeName(v))
        fmt.Printf("FromSlash:  %s\n", filepath.FromSlash(v))
        fmt.Printf("ToSlash:    %s\n", filepath.ToSlash(v))
        fmt.Printf("IsAbs:      %v\n", filepath.IsAbs(v))
        fmt.Printf("Join:       %s\n", filepath.Join(v, "z/x/c/", "w/h/g"))

        fmt.Println("----\n")
    }

    // Split
    fmt.Println("Split")
    p1 := "/k/u/t/g/c/index.html"
    dir, file := filepath.Split(p1)
    fmt.Printf("split dir:  %s\n", dir)
    fmt.Printf("split file: %s\n", file)
    fmt.Println("----\n")

    fmt.Println("SplitList")
    fmt.Printf("SplitLsit:    %v\n", filepath.SplitList(p1))
    fmt.Printf("SplitLsit len: %d\n", len(filepath.SplitList(p1)))
    fmt.Printf("SplitLsit cap: %d\n", cap(filepath.SplitList(p1)))
    fmt.Println("----\n")

    // match
    // Match(pattern, name string) (matched bool, err error)
    // * matches any sequence of non-/character
    // ? matches any single non-/character
    // [] character range
    fmt.Println("Match")
    pat := "**/*"
    matched, err := filepath.Match(pat, p1)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Matched: %v\n", matched)
    fmt.Println("----\n")
}
