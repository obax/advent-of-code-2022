package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board [][]string
type Instructions [][3]int

func parse_board(b string) Board {
	var board Board
	with_empty_boxes := regexp.MustCompile("\\s{4}").ReplaceAllString(b, "[ ]")
	without_boxes := regexp.MustCompile("[\\[\\]]").ReplaceAllString(with_empty_boxes, "")
	lines := strings.Split(string(without_boxes), "\n")
	reversed_lines := make([]string, len(lines))

	for idx, l := range lines {
		if idx == len(lines)-1 {
			continue
		}
		reversed_lines = append([]string{l}, reversed_lines...)
	}

	for _, rl := range reversed_lines {
		boxes := strings.Split(rl, " ")
		for idx_l, letter := range boxes {
			if idx_l >= len(board) {
				board = append(board, []string{})
			}
			if letter == "" {
				continue
			}
			board[idx_l] = append(board[idx_l], letter)
		}
	}
	return board
}

func parse_instructions(i string) Instructions {
	var inst Instructions
	re := regexp.MustCompile("move (\\d*)* from (\\d*) to (\\d*)")
	for _, l := range strings.Split(strings.Trim(i, "\n\n"), "\n") {
		matches := re.FindStringSubmatch(l)
		move, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		inst = append(inst, [3]int{move, from, to})
	}
	return inst
}

func get_last(b Board) string {
	solution := make([]string, len(b))
	for _, i := range b {
		solution = append(solution, i[len(i)-1])
	}
	return strings.Join(solution, "")
}

func sol1(board Board, instructions Instructions) string {
	b1 := make(Board, len(board))
	copy(b1, board)
	for _, instr := range instructions {
		num_move, from, to := instr[0], instr[1]-1, instr[2]-1
		end_from := b1[from][len(b1[from])-num_move:]
		b1[from] = b1[from][:len(b1[from])-num_move]

		for i := len(end_from) - 1; i >= 0; i-- {
			b1[to] = append(b1[to], end_from[i])
		}
	}
	return get_last(b1)
}

func sol2(board Board, instructions Instructions) string {
	b2 := make(Board, len(board))
	copy(b2, board)
	for _, instr := range instructions {
		num_move, from, to := instr[0], instr[1]-1, instr[2]-1
		end_from := b2[from][len(b2[from])-num_move:]
		b2[from] = b2[from][:len(b2[from])-num_move]
		b2[to] = append(b2[to], end_from...)
	}
	return get_last(b2)
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	sections := strings.Split(strings.Trim(string(input), ""), "\n\n")
	board_section, instruction_section := sections[0], sections[1]
	board := parse_board(board_section)
	instructions := parse_instructions(instruction_section)
  
	fmt.Println("solution 1", sol1(board, instructions))
	fmt.Println("solution 2", sol2(board, instructions))
}
