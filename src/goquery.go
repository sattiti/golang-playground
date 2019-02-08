package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/parnurzeal/gorequest"
    "strings"
)

func main() {
    url := "http://www.apple.com/"
    request := gorequest.New()
    resp, body, errs := request.Get(url).End()

    if len(errs) > 0 || resp.StatusCode != 200 {
        panic(errs)
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        panic(err)
    }

    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        href, exists := s.Attr("href")
        if exists {
            fmt.Println(href)
        }
    })
}
