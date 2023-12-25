package main

import (
  "github.com/gin-gonic/gin"
  "strconv"
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

  r.GET("/cancel", func (c *gin.Context){
    c.HTML(200, "", gin.H{})
  })

  r.GET("/round/:round", func (c *gin.Context){
    round := c.Param("round")

    nextRound := 0
    i, err := strconv.Atoi(round)
    if err != nil {
      nextRound = 0
    }
    nextRound = i + 1

    c.HTML(200, "roundPlay.tmpl", gin.H{
      "description":"Hello :)",
      "contract": "Two Sets",
      "nextRound": nextRound,
    })
  })

  r.Run()
}
