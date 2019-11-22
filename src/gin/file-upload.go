package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/yosssi/ace"
  "net/http"
  "path/filepath"
)

func main() {
  r := gin.Default()

  // 8mb
  r.MaxMultipartMemory = 8 << 20

  r.GET("/", func(c *gin.Context) {
    tmplName := "views/file-upload"
    html, err := ace.Load(tmplName, "", nil)

    if err != nil {
      c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
      return
    }

    err = html.Execute(c.Writer, gin.H{})

    if err != nil {
      c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
      return
    }

    // set up template
    r.SetHTMLTemplate(html)
    c.HTML(http.StatusOK, tmplName, gin.H{})
  })

  r.POST("/upload", func(c *gin.Context) {
    name := c.PostForm("name")
    email := c.PostForm("email")

    form, err := c.MultipartForm()
    if err != nil {
      c.String(http.StatusBadRequest, fmt.Sprintf("get form error: %s", err.Error()))
      return
    }

    files := form.File["files"]

    for _, file := range files {
      filename := filepath.Base(file.Filename)
      c.String(http.StatusOK, fmt.Sprintf("%s\n", filename))
      // if err := c.SaveUploadedFile(file, filename); err != nil {
      //   c.String(http.StatusBadRequest, fmt.Sprintf("Upload file error: %s", err.Error()))
      //   return
      // }
    }

    c.String(http.StatusOK, fmt.Sprintf("name: %s\n", name))
    c.String(http.StatusOK, fmt.Sprintf("email: %s\n", email))
    c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files\n", len(files)))
  })

  r.Run(":5001")
}
