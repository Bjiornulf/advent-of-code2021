package main

import (
	"fmt"
	"aoc2021/utils"
	"strings"
	"strconv"
	"time"
)

func main() {
	fmt.Println("day13")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Println("Puzzle1:", puzzle1(input))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Println("Puzzle2:", puzzle2(input))
	fmt.Println(time.Since(start))
}

type foldInstructions struct {
	x, y int
}

func getMax(input []string) (int, int) {
	xMax, yMax := 0, 0
	for i := range input {
		if input[i] == "" {
			break
		}
		v := strings.SplitN(input[i], ",", 2)
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		xMax = utils.IntMax(x, xMax)
		yMax = utils.IntMax(y, yMax)
	}
	return xMax+1, yMax+1
}

func importData(input []string) ([][]bool, []int) {
	i := 0
	xMax, yMax := getMax(input)
	dots := make([][]bool, xMax)
	for j := range dots {
		dots[j] = make([]bool, yMax)
	}
	for i = range input {
		if input[i] == "" {
			break
		}
		v := strings.SplitN(input[i], ",", 2)
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		dots[x][y] = true
	}
	i++
	instructions := make([]int, len(input)-i)
	for j := 0; j < len(instructions); j++ {
		instructions[j] = parseInstruction(input[j+i])
	}
	return dots, instructions
}

func parseInstruction(line string) int {
	instruction := strings.SplitN(strings.SplitN(line, " ", 3)[2], "=", 2)
	fold, _ := strconv.Atoi(instruction[1])
	switch instruction[0] {
	case "y":
		return -fold
	case "x":
		return fold
	default:
		return 0
	}
}

func isInside(paper [][]bool, x, y int) bool {
	return x >= 0 && x < len(paper) && y >= 0 && y < len(paper[x])
}

func fold(paper [][]bool, fold int, xLen, yLen int) (int, int) {
	if fold < 0 {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				if paper[x][y] {
					if isInside(paper, x, -y-2*fold) {
						paper[x][-y-2*fold] = true
					}
				}
			}
		}
		return xLen, -fold
	}
	if fold > 0 {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				if paper[x][y] {
					if isInside(paper, -x+2*fold, y) {
						paper[-x+2*fold][y] = true
					}
				}
			}
		}
		return fold, yLen
	}
	return xLen, yLen
}

func puzzle1(input []string) int {
	paper, foldInstructions := importData(input)
	xLen, yLen := fold(paper, foldInstructions[0], len(paper), len(paper[0]))
	count := 0
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			if paper[x][y] {
				count++
			}
		}
	}
	return count
}

func puzzle2(input []string) string {
	paper, foldInstructions := importData(input)
	xLen, yLen := len(paper), len(paper[0])
	for _, i := range foldInstructions {
		xLen, yLen = fold(paper, i, xLen, yLen)
	}
	var res strings.Builder
	res.WriteString("\n")
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if paper[x][y] {
				res.WriteRune('\u2598')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()[:res.Len()-1]
}
