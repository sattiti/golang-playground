package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // get
  r.GET("/get", func(c *gin.Context) {})

  // post
  r.POST("/post", func(c *gin.Context) {})

  // put
  r.PUT("/put", func(c *gin.Context) {})

  // delete
  r.DELETE("/delete", func(c *gin.Context) {})

  // patch
  r.PATCH("/patch", func(c *gin.Context) {})

  // head
  r.HEAD("/head", func(c *gin.Context) {})

  // options
  r.OPTIONS("/options", func(c *gin.Context) {})

  r.Run(":5001")
}
