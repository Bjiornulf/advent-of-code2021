package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type line struct {
	x1, x2, y1, y2 int
}

func main() {
	fmt.Println("day5")
	input := utils.ReadLines("input")
	lines := importData(input)

	/* ----------- Puzzle 1 ---------- */
	// only keep vertical or horrizontal lines
	puzzle1Lines := filterLines(lines, func(l line) bool { return (l.x1 == l.x2 || l.y1 == l.y2) })
	xmax, ymax := maxCoords(puzzle1Lines)
	puzzle1Grid := initGrid(xmax+1, ymax+1)
	fillGrid(&puzzle1Grid, puzzle1Lines)
	fmt.Printf("Puzzle1: %v\n", countMatrix(puzzle1Grid, func(i int) bool { return i >= 2 }))

	/* ----------- Puzzle 2 ---------- */
	puzzle2Lines := filterLines(lines, func(l line) bool {
		// only keep vertical, horrizontal or 45 degree lines
		return (l.x1 == l.x2 || l.y1 == l.y2 || abs(l.x1-l.x2) == abs(l.y1-l.y2))
	})
	xmax2, ymax2 := maxCoords(puzzle2Lines)
	puzzle2Grid := initGrid(xmax2+1, ymax2+1)
	fillGrid(&puzzle2Grid, puzzle2Lines)
	fmt.Printf("Puzzle2: %v\n", countMatrix(puzzle2Grid, func(i int) bool { return i >= 2 }))
}

// make an [][]int grid of size width * height
func initGrid(width int, height int) [][]int {
	var grid [][]int = make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	return grid
}

// fills [][]int grid with lines in []line
// since grids can be big, does not return a new grid
// lines should be horrizontal, vertical or 45 deg. Otherwise unexpected behaviour
func fillGrid(grid *[][]int, lines []line) {
	for _, l := range lines {
		xRange, yRange := enumPositions(l)
		for i := range xRange {
			(*grid)[yRange[i]][xRange[i]]++
		}
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// return 1 if a > b; 0 if a == b; -1 if a < b
func intCmp(a int, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// creates an array with the points a line crosses
func enumPositions(l line) (x []int, y []int) {
	size := max(abs(l.x1-l.x2), abs(l.y1-l.y2)) // number of points
	x = make([]int, size+1)
	y = make([]int, size+1)
	xIncr, yIncr := intCmp(l.x2, l.x1), intCmp(l.y2, l.y1)
	for i := 0; i <= size; i++ {
		x[i] = l.x1 + xIncr*i
		y[i] = l.y1 + yIncr*i
	}
	return x, y
}

// counts the number of occurences in a matrix that match the filter function (filter(x) == true)
func countMatrix(matrix [][]int, filter func(int) bool) int {
	var count int
	for i := range matrix {
		for j := range matrix[0] {
			if filter(matrix[i][j]) {
				count++
			}
		}
	}
	return count
}

// import the data for Day 5
func importData(lines []string) []line {
	res := make([]line, len(lines))
	var x1, x2, y1, y2 int
	var fields, start, end []string
	for i, l := range lines {
		fields = strings.Fields(l) // x1,y1 -> x2,y2 = "x1,y1" "->" "x2,y2"
		start = strings.SplitN(fields[0], ",", 2)
		end = strings.SplitN(fields[2], ",", 2)
		x1, _ = strconv.Atoi(start[0])
		y1, _ = strconv.Atoi(start[1])
		x2, _ = strconv.Atoi(end[0])
		y2, _ = strconv.Atoi(end[1])
		res[i] = line{x1: x1, y1: y1, x2: x2, y2: y2}
	}
	return res
}

// filter an array of lines according to filter function; returns a new array
func filterLines(lines []line, filter func(line) bool) []line {
	res := make([]line, 0)
	for _, l := range lines {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}

// return the maximum x and y coordinates in a list of lines
func maxCoords(lines []line) (int, int) {
	var xmax, ymax int
	for _, l := range lines {
		xmax = max(xmax, max(l.x1, l.x2))
		ymax = max(ymax, max(l.y1, l.y2))
	}
	return xmax, ymax
}

// absolute function for integers
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
