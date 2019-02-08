package main

import (
  "fmt"
  "io"
  "io/ioutil"
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

func sjisDecode(ior io.Reader) ([]byte, error) {
  return ioutil.ReadAll(transform.NewReader(ior, japanese.ShiftJIS.NewDecoder()))
}

func sjisEncode(ior io.Reader) ([]byte, error) {
  return ioutil.ReadAll(transform.NewReader(ior, japanese.ShiftJIS.NewEncoder()))
}

func main() {
  run(func() {
    fpath, _ := filepath.Abs("./index.html")
    opath, _ := filepath.Abs("./out.html")

    f, err := os.OpenFile(fpath, os.O_RDONLY, 0644)
    throwError(err)
    defer f.Close()

    out, err := sjisDecode(f)
    throwError(err)

    fmt.Printf("%v\n", string(out))

    outsjis, err := sjisEncode(strings.NewReader(string(out)))
    throwError(err)

    err = ioutil.WriteFile(opath, outsjis, 0644)
    throwError(err)
  })
}
