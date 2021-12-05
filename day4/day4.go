package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day4")
	puzzle1()
}

type cell struct {
	val   int
	drawn bool
}

type grid [5][5]cell

func puzzle1() {
	input := utils.ReadLines("input")
	var numbers []int = make([]int, len(strings.Split(input[0], ",")))
	var bingoGrid *grid = new(grid)
	var bingo []grid = make([]grid, 0)
	// getting the numbers in the order they will be drawn
	for i, number := range strings.Split(input[0], ",") {
		intNumber, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		numbers[i] = intNumber
	}
	row := 0
	// filling the bingo grids and creating the bingo array
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			row = 0
			bingo = append(bingo, *bingoGrid)
			bingoGrid = new(grid)
		} else {
			fill_grid(input[i], row, bingoGrid)
			row++
		}
	}
	// last bingo grid is not followed by an empty line, thus we need to treat it separately
	bingo = append(bingo, *bingoGrid)

	// finding the winners of the first and second round
	firstWinnerFound := false
	var calledNum, lastWinnningCalledNum, winnersFound, lastWinnerValue int
	var won []bool = make([]bool, len(bingo)) // each grid only wins once, thus we need to keep track of all the grids which have previously won
	for _, calledNum = range numbers {
		for i := range bingo {
			mark_grid(calledNum, &bingo[i])
			if !won[i] && validate_grid(&bingo[i]) {
				winnersFound++
				won[i] = true
				lastWinnerValue = grid_value(&bingo[i])
				lastWinnningCalledNum = calledNum
				if !firstWinnerFound {
					firstWinnerFound = true
					fmt.Println("Puzzle1:", lastWinnerValue*lastWinnningCalledNum)
				}
			}
		}
	}
	fmt.Println("Puzzle2:", lastWinnerValue*lastWinnningCalledNum)
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
		validateRow := true
		for c := 0; c < 5; c++ {
			validateRow = validateRow && (*grid)[r][c].drawn
		}
		if validateRow {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		validateRow := true
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
		if err != nil {
			panic(err)
		}
		(*grid)[row][pos].val = num
		(*grid)[row][pos].drawn = false
	}
}
