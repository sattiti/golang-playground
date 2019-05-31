package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main() {
  ip := "192.168.0.1"
  bip := []string{}

  for _, v := range strings.Split(ip, ".") {
    n, err := strconv.ParseInt(v, 10, 32)
    if err != nil {
      panic(err)
    }
    bip = append(bip, strconv.FormatInt(n, 2))
  }
  fmt.Println(strings.Join(bip, "."))
}
