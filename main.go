package main

import (
	"contract_rummy/app/rounds"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	r.GET("/newPlayer", func(c *gin.Context) {
		// TODO: Create sqlite connection to create players
		c.HTML(200, "newPlayer.tmpl", gin.H{})
	})

	r.GET("/cancel", func(c *gin.Context) {
		c.HTML(200, "", gin.H{})
	})

	r.GET("/round/:round", func(c *gin.Context) {
		// TODO: Extract this to db
		round := c.Param("round")

		curr := rounds.GetRound(round)
		fmt.Println(curr)
		c.HTML(200, "roundPlay.tmpl", gin.H{
			"round":       round,
			"description": curr.Description,
			"contract":    curr.Contract,
			"nextRound":   curr.NextRound,
		})
	})

	r.Run(":8080")
}
