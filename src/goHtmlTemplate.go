package main

import (
    "fmt"
    "html/template"
    "os"
    "path/filepath"
)

func run(f func()) {
    fmt.Printf("----\n")
    f()
    fmt.Printf("----\n\n")
}

type Li struct {
    Value string
    Href  string
    Attr  map[string]string
}

func main() {
    run(func() {
        tplfile := "./index.html.tpl"

        data := map[string]interface{}{
            "Title": "hello",
            "H1":    "hello world",
            "Li": []*Li{
                &Li{
                    Value: "Home",
                    Href:  "/",
                    Attr:  map[string]string{"class": "hli", "id": "hli"},
                },
                &Li{
                    Value: "About",
                    Href:  "/about/",
                    Attr:  map[string]string{"class": "hli", "id": "hli"},
                },
            },
        }

        tpl := template.Must(template.New(filepath.Base(tplfile)).Funcs(template.FuncMap{
            "safeAttr": func(s string) template.HTMLAttr {
                return template.HTMLAttr(s)
            },
        }).ParseFiles(tplfile))

        err := tpl.Execute(os.Stdout, data)
        if err != nil {
            panic(err)
        }
    })
}
