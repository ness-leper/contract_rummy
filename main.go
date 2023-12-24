package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.GET("/", func(c *gin.Context){
    c.HTML(200, "index.tmpl", gin.H{
        "title":"Main Website",
    })
  })

  r.GET("/newPlayer", func (c *gin.Context){
    c.HTML(200, "newPlayer.tmpl", gin.H{})
  })

  r.Run()
}
