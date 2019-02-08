package main

import (
  "bytes"
  "html/template"
  "io/ioutil"
  "path/filepath"
)

func templateParse(tplFile string, data interface{}) (*template.Template, bytes.Buffer, error) {
  var doc bytes.Buffer
  in, _ := ioutil.ReadFile(tplFile)

  t := template.Must(template.New(filepath.Base(tplFile)).Funcs(template.FuncMap{
    "raw": func(v string) template.HTML {
      return template.HTML(v)
    },
  }).Parse(string(in)))

  err := t.Execute(&doc, data)
  return t, doc, err
}

func main() {
  tplFile := "./index.html.tpl"
  outFile := "./index.html"
  data := map[string]string{
    "Title": "hello",
    "H1":    "world",
  }
  _, doc, err := templateParse(tplFile, &data)
  if err != nil {
    panic(err)
  }
  ioutil.WriteFile(outFile, doc.Bytes(), 0644)
}
