package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  r := gin.Default()

  // Parameters
  // This handler will match
  //   - `/user/john`
  // but will not match
  //   - `/user/`
  //   - `/user`
  r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // name with action
  // This handler will match
  //   - `/user/john/`
  //   - `/user/john/send`
  r.GET("user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, "%s\n", message)
    c.String(http.StatusOK, "%s\n", c.FullPath())
  })

  r.Run(":5001")
}
