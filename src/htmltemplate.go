package main

import (
    "html/template"
    "log"
    "os"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`(?ims)^[ \s\t\n]+`)
    tmp := re.ReplaceAllString(`<!DOCTYPE html>`, "")

    out := "./res.index.html"
    t, err := template.New("name").Parse(tmp)
    if err != nil {
        log.Fatal(err)
    }

    // openfile
    f, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR, 0775)
    if err != nil {
        log.Fatal(err)
    }

    err = t.Execute(f, struct {
        Title string
        H1    string
    }{
        Title: "Test",
        H1:    "hello world",
    })

    if err != nil {
        log.Fatal(err)
    }
}
