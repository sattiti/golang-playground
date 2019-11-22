package main

import (
  "github.com/gin-gonic/gin"
  "github.com/yosssi/ace"
  "net/http"
)

func main() {
  r := gin.Default()

  // get
  r.GET("/", func(c *gin.Context) {
    tmplName := "views/index"

    // Use ace template
    html, err := ace.Load(tmplName, "", nil)

    if err != nil {
      http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
      return
    }

    // render ace template
    err = html.Execute(c.Writer, gin.H{
      "H1":   "hello",
      "Html": html,
    })

    if err != nil {
      http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
      return
    }

    // set up template
    r.SetHTMLTemplate(html)
    c.HTML(http.StatusOK, tmplName, gin.H{})
  })

  r.Run(":5001")
}
