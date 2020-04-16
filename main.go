package main

import (
	"math/rand"
)

func buildDeck() []int {
	deck := []int{
		1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3,
		4, 4, 4, 4,
		5, 5, 5, 5,
		6, 6, 6, 6,
		7, 7, 7, 7,
		8, 8, 8, 8,
		9, 9, 9, 9,
		10, 10, 10, 10,
		10, 10, 10, 10,
		10, 10, 10, 10,
		10, 10, 10, 10,
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func main() {
	rand.Seed(981203981045791047)

}
