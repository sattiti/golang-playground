package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  r := gin.Default()

  // Querystring
  r.GET("/qs", func(c *gin.Context) {

    // firstname キーがなければ、`Guest` というデフォルト値を代入する。
    firstname := c.DefaultQuery("firstname", "Guest")

    // shortcut for `c.Request.URL.Query().Get("lastname")`
    lastname := c.Query("lastname")

    // ids[a]=111&ids[b]=222
    // のように取得することができる。
    ids := c.QueryMap("ids")

    c.String(http.StatusOK, "Hello %s %s\n%#v", firstname, lastname, ids)
  })
  r.Run(":5001")
}
