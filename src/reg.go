package main

import (
  "fmt"
  _ "log"
  "regexp"
)

func main() {
  re := regexp.MustCompile(`(?ims)(hello .*?)<`)

  doc := `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width">
<title>hello world</title>
</head>
<body>
<h1>Hello World</h1>
</body>
</html>
`
  fmt.Println(re.MatchString(doc))
  fmt.Printf("%#v\n", re.FindAllString(doc, -1))

  // print matched group
  groups := re.FindAllStringSubmatch(doc, -1)
  for _, m := range groups {
    fmt.Println(m[1])
  }

  // replace matched string
  fmt.Printf("%#v\n", re.ReplaceAllString(doc, `Aloha world`))
}
