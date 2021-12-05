package main

import (
	"fmt"
	"aoc2021/utils"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("day4")
	puzzle1()
}

type cell struct {
	val int
	drawn bool
}

type grid [5][5]cell

func puzzle1() {
	lines := utils.ReadLines("input")
	var inputs []int = []int{}
	var bingoGrid *grid = new(grid)
	var bingo []grid = make([]grid, 0)
	// getting the numbers in the order they will be drawn
	for _, input := range strings.Split(lines[0], ",") {
		in, err := strconv.Atoi(input)
		if (err != nil) {
			panic(err)
		}
		inputs = append(inputs, in)
	}
	row := 0
	for i := 2; i < len(lines); i++ {
		if (lines[i] == "") {
			row = 0
			bingo = append(bingo, *bingoGrid)
			bingoGrid = new(grid)
		} else {
			fill_grid(lines[i], row, bingoGrid)
			row++
		}
	}
	bingo = append(bingo, *bingoGrid)
	// fmt.Println(bingo)
	foundWinner := false
	var calledNum, lastWinner, lastWinnningCalledNum, winnersFound, lastWinnerValue int
	lastWinnerValue = 0
	_ = lastWinner
	var won []bool = make([]bool, len(bingo))
	for _, calledNum = range inputs {
		max := 0
		for i := range bingo {
			mark_grid(calledNum, &bingo[i])
			if !won[i] && validate_grid(&bingo[i]) {
				winnersFound++
				foundWinner = true
				won[i] = true
				lastWinner = i
				lastWinnerValue = grid_value(&bingo[i])
				lastWinnningCalledNum = calledNum
				if (max < calledNum * grid_value(&bingo[i])) {
					max = calledNum * grid_value(&bingo[i])
				}
			}
		}
		if foundWinner && winnersFound == 1 {
			fmt.Printf("Puzzle1: %v\n", max)
		}
	}
	fmt.Printf("Puzzle2: %v\n", lastWinnerValue * lastWinnningCalledNum)
}

func grid_value(grid *grid) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !(*grid)[i][j].drawn {
				sum += (*grid)[i][j].val
			}
		}
	}
	return sum
}

func mark_grid(number int, grid *grid) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if (*grid)[r][c].val == number {
				(*grid)[r][c].drawn = true
			}
		}
	}
}

func validate_grid(grid *grid) bool {
	for r := 0; r < 5; r++ {
		validateRow := true;
		for c := 0; c < 5; c++ {
			validateRow = validateRow && (*grid)[r][c].drawn
		}
		if validateRow {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		validateRow := true;
		for r := 0; r < 5; r++ {
			validateRow = validateRow && (*grid)[r][c].drawn
		}
		if validateRow {
			return true
		}
	}
	return false
}

func fill_grid(str string, row int, grid *grid) {
	for pos, val := range strings.Fields(str) {
		num, err := strconv.Atoi(val)
		if (err != nil) {
			panic(err)
		}
		(*grid)[row][pos].val = num;
		(*grid)[row][pos].drawn = false;
	}
}
