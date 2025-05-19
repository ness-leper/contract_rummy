package rounds

import (
	"fmt"
	"strconv"
)

type Round struct {
	Description string
	Contract    string
	NextRound   int
}

func GetRound(round string) Round {
	var rounds = []Round{
		{"Deal 10 cards", "2 Sets", 2},
		{"Deal 10 cards", "1 Set and 1 Sequence", 3},
		{"Deal 10 cards", "2 Sequence's", 4},
		{"Deal 10 cards", "3 Sets", 5},
		{"Deal 12 cards", "2 Sets and 1 Sequence", 6},
		{"Deal 12 cards", "1 Set and 2 Sequence's", 7},
		{"Deal 12 cards", "Three Sequence's", 1},
	}

	nextRound := 0
	i, err := strconv.Atoi(round)
	if err != nil {
		nextRound = 0
	}
	nextRound = i + 1

	if nextRound > 7 {
		nextRound = 1
	}

	fmt.Print("\n\n\n")
	fmt.Println("The Next round is", rounds[i-1].NextRound)
	fmt.Println("The current round is", i)
	fmt.Print("\n\n\n")
	return rounds[i-1]
}
