package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var beats = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
}

var loses = map[string]string{
	"Rock":     "Paper",
	"Paper":    "Scissors",
	"Scissors": "Rock",
}

var scores = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

func rule_1_get_score(p1 string, p2 string) int {
	var extra int
	var table = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	extra = scores[table[p2]]
	var round int
	if table[p1] == table[p2] {
		round = 3
	} else if beats[table[p1]] == table[p2] {
		round = 0
	} else {
		round = 6
	}
	return extra + round
}

func rule_2_get_score(c1 string, c2 string) int {
	var extra int
	var round int

	var c1_codes = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	var c2_codes = map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}
	var p2choice string
	switch c2_codes[c2] {
	case "Draw":
		p2choice = c1_codes[c1]
		round = 3
	case "Lose":
		p2choice = beats[c1_codes[c1]]
		round = 0
	case "Win":
		p2choice = loses[c1_codes[c1]]
		round = 6
	}

	extra = scores[p2choice]
	return extra + round
}

func main() {
	var scores1 int
	var scores2 int
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range strings.Split(string(input), "\n") {
		split := strings.Split(string(l), " ")
		if len(split) == 2 {
			scores1 += rule_1_get_score(split[0], split[1])
			scores2 += rule_2_get_score(split[0], split[1])
		}
	}
	fmt.Println("part 1 scores", scores1)
	fmt.Println("part 2 scores", scores2)
}
