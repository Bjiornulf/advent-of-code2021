package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"time"
)

func main() {
	fmt.Println("day17")
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

type Vector struct {
	x, y int
}

var start = Vector{0, 0}
func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}

func importData(input []string) (Vector, Vector) {
	var x1, y1, x2, y2 int
	vals := strings.FieldsFunc(input[0], func (c rune) bool {
		return !unicode.IsNumber(c) && c != '-'
	})
	x1, _ = strconv.Atoi(vals[0])
	x2, _ = strconv.Atoi(vals[1])
	y1, _ = strconv.Atoi(vals[2])
	y2, _ = strconv.Atoi(vals[3])
	return Vector{min(x1, x2), min(y1, y2)}, Vector{max(x2, x1), max(y2, y1)}
}

func isInside(point, targetStart, targetEnd Vector) bool {
	return (
		point.x >= targetStart.x &&
		point.y >= targetStart.y &&
		point.x <= targetEnd.x &&
		point.y <= targetEnd.y )
}

func canLand(point, velocity, targetStart, targetEnd Vector) bool {
	return !(
		point.x < targetStart.x && velocity.x <= 0 ||
		point.y < targetStart.y && velocity.y <= 0 ||
		point.x > targetEnd.x && velocity.x >= 0 )
}

func simulate(start, velocity, targetStart, targetEnd Vector) (int, bool) {
	maxY := start.y
	for !isInside(start, targetStart, targetEnd) && canLand(start, velocity, targetStart, targetEnd) {
		start.x += velocity.x
		start.y += velocity.y
		maxY = max(start.y, maxY)
		if velocity.x > 0 {
			velocity.x--
		} else if velocity.x < 0 {
			velocity.x++
		}
		velocity.y--
	}
	return maxY, isInside(start, targetStart, targetEnd)
}

func puzzle1(input []string) int {
	// supposing the target is always under the start position
	targetStart, _ := importData(input)
	return -targetStart.y * (-targetStart.y - 1) / 2
}

func puzzle2(input []string) int {
	targetStart, targetEnd := importData(input)
	xMin := min(start.x, targetStart.x)
	xMax := max(start.x, targetEnd.x)
	yMin := min(start.y, targetStart.y)
	yMax := max(start.y, utils.IntAbs(targetStart.y))
	res := 0
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y < yMax; y++ {
			_, success := simulate(start, Vector{x, y}, targetStart, targetEnd)
			if success {
				res++
			}
		}
	}
	return res
}
