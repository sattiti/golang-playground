package main

import (
    "encoding/xml"
    "fmt"
)

func run(f func()) {
    fmt.Printf("----\n")
    f()
    fmt.Printf("----\n\n")
}

type Result struct {
    XMLName xml.Name `xml:"root"`
    Person  []Person `xml:"person"`
}

type Person struct {
    Fullname string   `xml:"fullname"`
    Emails   []Email  `xml:"email"`
    Groups   []string `xml:"group>item"`
}

type Email struct {
    Mailaddress Mailaddress `xml:"mailaddress"`
}

type Mailaddress struct {
    Addr  string `xml:",chardata"`
    Where string `xml:"where,attr"`
}

func main() {
    data := `<?xml version="1.0" encoding="utf-8"?>
  <root>
    <person>
      <fullname>bbb</fullname>
      <email>
        <mailaddress where="home">bbb@home.com</mailaddress>
        <mailaddress where="work">bbb@work.com</mailaddress>
      </email>
      <group>
        <item>football</item>
        <item>friends</item>
      </group>
    </person>
    <person>
      <fullname>aaa</fullname>
      <email>
        <mailaddress where="home">aaa@home.com</mailaddress>
        <mailaddress where="work">aaa@work.com</mailaddress>
      </email>
      <group>
        <item>friends</item>
        <item>friends2</item>
      </group>
    </person>
  </root>
  `
    run(func() {
        v := &Result{}
        err := xml.Unmarshal([]byte(data), &v)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%#v\n", v)
    })
}
