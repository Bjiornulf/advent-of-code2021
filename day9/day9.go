package main

import (
	"aoc2021/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := utils.ReadLines("input")
	fmt.Println("day9")

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", puzzle1(input))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", puzzle2(input))
}

func puzzle1(input []string) int {
	lows, _ := lowLevels(importData(input))
	sum := 0
	for _, l := range lows {
		sum += l + 1
	}
	return sum
}

func puzzle2(input []string) int {
	bassins := bassinsSizes(importData(input))
	sort.Ints(bassins)
	end := len(bassins) - 1
	return bassins[end] * bassins[end-1] * bassins[end-2]
}

/* ---------- Structures and their methods ---------- */
type Coord struct {
	i, j int
}

/* ---------- Functions ---------- */
// import the data for the day
func importData(input []string) [][]int {
	res := make([][]int, len(input))
	for i := range input {
		res[i] = make([]int, len(input[i]))
		for j := range input[i] {
			level, _ := strconv.Atoi(string(input[i][j]))
			res[i][j] = level
		}
	}
	return res
}

// return all points that are low level (their level and their coordinates)
func lowLevels(mat [][]int) ([]int, []Coord) {
	res := make([]int, 0)
	coords := make([]Coord, 0)
	for i := range mat {
		for j := range mat[i] {
			isLow := true
			if i == 0 {
				isLow = isLow && (mat[i+1][j] > mat[i][j])
			} else if i == len(mat)-1 {
				isLow = isLow && (mat[i-1][j] > mat[i][j])
			} else {
				isLow = isLow && (mat[i-1][j] > mat[i][j]) && (mat[i+1][j] > mat[i][j])
			}
			if j == 0 {
				isLow = isLow && (mat[i][j+1] > mat[i][j])
			} else if j == len(mat[i])-1 {
				isLow = isLow && (mat[i][j-1] > mat[i][j])
			} else {
				isLow = isLow && (mat[i][j-1] > mat[i][j]) && (mat[i][j+1] > mat[i][j])
			}
			if isLow {
				res = append(res, mat[i][j])
				coords = append(coords, Coord{i: i, j: j})
			}
		}
	}
	return res, coords
}

func bassinsSizes(mat [][]int) []int {
	lows, coords := lowLevels(mat)
	explored := make([][]bool, len(mat))
	for i := range explored {
		explored[i] = make([]bool, len(mat[i]))
	}
	res := make([]int, 0)
	for i := range lows {
		size := 0
		recExplore(mat, coords[i], explored, &size)
		res = append(res, size)
	}
	return res
}

func recExplore(mat [][]int, start Coord, explored [][]bool, size *int) {
	// An imperative alternative has been tested, but is ~4 times slower
	i, j := start.i, start.j
	if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[i]) {
		return // this is not explorable (outisde of matrix bounds)
	}
	if explored[i][j] {
		return // this has been explored
	}
	if mat[i][j] == 9 {
		return // 9 is never in a basin
	}

	// Every point that can be reached by exploring the neighbours only is in the same bassin

	/* ---------- Recursive exploration of the neightbours ---------- */
	explored[i][j] = true
	(*size)++
	recExplore(mat, Coord{i: i - 1, j: j}, explored, size)
	recExplore(mat, Coord{i: i + 1, j: j}, explored, size)
	recExplore(mat, Coord{i: i, j: j + 1}, explored, size)
	recExplore(mat, Coord{i: i, j: j - 1}, explored, size)
}
