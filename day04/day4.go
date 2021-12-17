package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const drawn int = -1

type grid struct {
	values [5][5]int
	score  int
}

func main() {
	fmt.Println("day4")
	input := utils.ReadLines("input")
	numbers := utils.ReadIntsList(input[0], ",")
	bingoGrid := new(grid)
	grids := make([]grid, 0)

	// filling the bingo grids and creating the bingo array
	row := 0
	for i := 2; i < len(input); i++ {
		if input[i] == "" { // end of grid
			row = 0
			grids = append(grids, *bingoGrid)
			bingoGrid = new(grid)
		} else {
			fillGrid(input[i], row, bingoGrid)
			row++
		}
	}
	// last bingo grid is not followed by an empty line, thus we need to treat it separately
	grids = append(grids, *bingoGrid)

	orderedGrids := fillGrids(numbers, grids)

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", orderedGrids[0].score)

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", orderedGrids[len(orderedGrids)-1].score)
}

// fill the grids until every number has been drawn
// return an array of pointers to the grids in the order they have won
func fillGrids(draws []int, grids []grid) []*grid {
	res := make([]*grid, 0)
	for _, number := range draws {
		// very important to use indexes and access the elements of the array. Otherwise, we can't change their value!
		// range returns a copy of the element!
		for i := range grids {
			// if grid has not won, we fill it
			if grids[i].score == 0 {
				markGrid(number, &grids[i])
				if validateGrid(&grids[i]) {
					grids[i].score = number * gridValue(&grids[i])
					res = append(res, &grids[i])
				}
			}
		}
	}
	// return pointers to the grids in the order they won
	return res
}

// return the value of a grid (sum of the numbers that have not been marked)
func gridValue(grid *grid) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid.values[i][j] != drawn {
				sum += grid.values[i][j]
			}
		}
	}
	return sum
}

// mark the number on the grid (if it exists)
func markGrid(number int, grid *grid) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if grid.values[r][c] == number {
				grid.values[r][c] = drawn
				return // assuming every grid has a number only once
			}
		}
	}
}

// is a grid completed (i.e. is there a complete row or collumn)
func validateGrid(grid *grid) bool {
	for i := 0; i < 5; i++ {
		validateRow := true
		validateCollumn := true
		for j := 0; j < 5; j++ {
			validateRow = validateRow && grid.values[i][j] == drawn
			validateCollumn = validateCollumn && grid.values[j][i] == drawn
		}
		if validateRow || validateCollumn {
			return true
		}
	}
	return false
}

// fill a row of a grid with values
func fillGrid(str string, row int, grid *grid) {
	for pos, val := range strings.Fields(str) {
		num, _ := strconv.Atoi(val)
		grid.values[row][pos] = num
	}
}
