package main

import (
  _ "fmt"
  "github.com/gin-gonic/gin"
  _ "html/template"
  "net/http"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    urls := map[string]string{
      "index": "/",
      "form":  "/form",
    }
    r.LoadHTMLFiles("./views/index.tmpl")
    c.HTML(http.StatusOK, "index.tmpl", gin.H{"UrlList": urls})
  })

  r.GET("form", func(c *gin.Context) {
    r.LoadHTMLFiles("./views/post-method.tmpl")
    c.HTML(http.StatusOK, "post-method.tmpl", gin.H{})
  })

  r.POST("/submit", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")
    name := c.PostFormMap("name")

    c.JSON(http.StatusOK, gin.H{
      "status":     "posted",
      "message":    message,
      "nick":       nick,
      "first-name": name["first"],
      "last-name":  name["last"],
    })
  })

  r.Run(":5001")
}
