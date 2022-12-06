package main

import (
	"fmt"
	"os"
)

func all_unique(s string) bool {
  set := make(map[rune]bool)
  for _, r := range s {
    if _, pres := set[r]; pres {
      return false
    }
    set[r] = true
  }
  return true
}

func identify_distinct(input string, number_distinct int) int {
  for i:=number_distinct; len(input) > i; i++ {
    prev_distinct := input[i-number_distinct:i]
    if all_unique(prev_distinct) {
      return i
    }
  }
  return 0
}


func main(){
  input, err := os.ReadFile("./input.txt"); if err != nil {
    panic(err)
  }

  fmt.Println("solution 1", identify_distinct(string(input), 4))
  fmt.Println("solution 2", identify_distinct(string(input), 14))
}