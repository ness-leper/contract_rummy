package main

import (
	"strconv"
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

		var rounds = []struct {
			description string
			contract    string
		}{
			{"Deal 10 cards", "2 Sets"},
			{"Deal 10 cards", "1 Set and 1 Sequence"},
			{"Deal 10 cards", "2 Sequence's"},
			{"Deal 10 cards", "3 Sets"},
			{"Deal 12 cards", "2 Sets and 1 Sequence"},
			{"Deal 12 cards", "1 Set and 2 Sequence's"},
			{"Deal 12 cards", "Three Sequence's"},
		}
		round := c.Param("round")

		nextRound := 0
		i, err := strconv.Atoi(round)
		if err != nil {
			nextRound = 0
		}
		nextRound = i + 1

		if nextRound > 7 {
			nextRound = 1
		}

		c.HTML(200, "roundPlay.tmpl", gin.H{
			"round":       testFunction(),
			"description": rounds[i-1].description,
			"contract":    rounds[i-1].contract,
			"nextRound":   nextRound,
		})
	})

	r.Run()
}
