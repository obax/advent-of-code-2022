package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

func get_data() (harvest []int) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sep := strings.Split(string(file), "\n\n")
	for _, i := range sep {
		var curr int
		for _, j := range strings.Split(i, "\n") {
			val, _ := strconv.Atoi(j)
			curr += val
			harvest = append(harvest, curr)
		}
	}
	return harvest
}

func main() {
	var max int
	harvest := get_data()
	for _, i := range harvest {
		if i > max { max = i}
	}
	fmt.Println("top elf", max)
	sort.Sort(sort.Reverse(sort.IntSlice(harvest)))
	top3 := 0
	for _, i := range harvest[:3] {
		top3 += i
	}
	fmt.Println("top 3", top3)
}
