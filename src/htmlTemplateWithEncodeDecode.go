package main

import (
  "bytes"
  "html/template"
  "io"
  "io/util"
  "os"
  "path/filepath"
  "strings"

  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)

func run(f func()) {
  fmt.Printf("----\n")
  f()
  fmt.Printf("----\n\n")
}

func throwError(e error) {
  if e != nil {
    panic(e)
  }
}

func iso2022jpEncode(ior io.Reader) ([]bytes, error) {
  return ioutil.ReadAll(transform.NewReader(ior, japanese.ISO2022JP.NewEncoder()))
}

func iso2022jpDecode(ior io.Reader) ([]bytes, error) {
  return ioutil.ReadAll(transform.NewReader(ior, japanese.ISO2022JP.NewDecoder()))
}

func main() {
  run(func() {
    inFile, _ := filepath.Abs("./index.html.tpl")
    outFile, _ := filepath.Abs("./out.html")

    data := map[string]string{
      "Title": "hello world",
    }

    in, _ := os.OpenFile(inFile, os.O_RDONLY, 0644)
    defer in.Close()

    inBytes, _ := iso2022jpDecode(in)

    var doc bytes.Buffer
    t := template.Must(template.New(filepath.Base(inFile)).Parse(string(inBytes)))
    if err := t.Execute(&doc, data); err != nil {
      panic(err)
    }

    outBytes, _ := iso2022jpEncode(strings.NewReader(doc.String()))
    if err := ioutil.WriteFile(outFile, outBytes, 0644); err != nil {
      panic(err)
    }
  })
}
