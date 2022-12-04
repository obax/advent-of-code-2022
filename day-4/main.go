package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_pairs(p1 []string, p2 []string) ([2]int, [2]int) {
	min1, _ := strconv.Atoi(p1[0])
	max1, _ := strconv.Atoi(p1[1])
	min2, _ := strconv.Atoi(p2[0])
	max2, _ := strconv.Atoi(p2[1])
	return [2]int{min1, max1}, [2]int{min2, max2}
}

func fully_contains(p1 [2]int, p2 [2]int) bool {
	min1, max1, min2, max2 := p1[0], p1[1], p2[0], p2[1]
	inc_col1 := min1 >= min2 && max2 >= max1
	inc_col2 := min2 >= min1 && max1 >= max2
	return inc_col1 || inc_col2
}

func has_overlaps(p1 [2]int, p2 [2]int) bool {
	min1, max1, min2, max2 := p1[0], p1[1], p2[0], p2[1]
	set := make(map[int]bool)
	for i := min1; max1 >= i; i++ {
		set[i] = true
	}
	for i := min2; max2 >= i; i++ {
		if set[i] == true {
			return true
		}
	}
	return false
}

func main() {
	var count_fully_contains int
	var count_have_overlaps int

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(input), ""), "\n")
	for _, l := range lines {
		pairs := strings.Split(l, ",")
		c1, c2 := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
		col1, col2 := parse_pairs(c1, c2)
		if fully_contains(col1, col2) {
			count_fully_contains++
		}
		if has_overlaps(col1, col2) {
			count_have_overlaps++
		}
	}
	fmt.Println("part 1", count_fully_contains)
	fmt.Println("part 2", count_have_overlaps) // 679
}
