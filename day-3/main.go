package main

import (
	"fmt"
	"os"
	"strings"
)

const CAPITAL_LETTER_OFFSET int32 = 32 + 26

func unique(s string) string {
	set := make(map[string]bool)
	var dedup string
	for _, l := range strings.Split(s, "") {
		set[l] = true
	}
	for k := range set {
		dedup += k
	}
	return dedup
}

func sum_points(s string) int {
	var score int
	for _, o := range unique(s) {
		var current int32 = rune(o) - rune('a') + 1
		if 0 > current {
			current = current + CAPITAL_LETTER_OFFSET
		}
		score += int(current)
	}
	return score
}

func parse_parts(p1 string, p2 string) int {
	var overlaps string
	for _, letter_p1 := range p1 {
		for _, letter_p2 := range p2 {
			if letter_p1 == letter_p2 {
				overlaps += string(letter_p1)
			}
		}
	}
	return sum_points(overlaps)
}

func parse_parts_3(p1 string, p2 string, p3 string) int {
	var overlaps string
	for _, letter_p1 := range p1 {
		for _, letter_p2 := range p2 {
			for _, letter_p3 := range p3 {
				if letter_p1 == letter_p2 && letter_p2 == letter_p3 {
					overlaps += string(letter_p1)
				}
			}
		}
	}
	return sum_points(overlaps)
}

func main() {
	var total1 int
	var total2 int

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	for _, l := range lines {
		item1, item2 := l[:len(l)/2], l[len(l)/2:]
		total1 += parse_parts(item1, item2)
	}
	fmt.Println("part 1", total1)

	for i := 0; len(lines)/3 > i; i++ {
		cursor := i * 3
		bags := lines[cursor : cursor+3]
		total2 += parse_parts_3(bags[0], bags[1], bags[2])
	}

	fmt.Println("part 2", total2)
}
