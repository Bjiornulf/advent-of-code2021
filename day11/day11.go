package main

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

const (
	days1 = 100
)

func main() {
	fmt.Println("day11")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Println("Puzzle1:", puzzle1(input))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Println("Puzzl2:", puzzle2(input))
	fmt.Println(time.Since(start))
}

type Coord struct {
	i, j int
}

func puzzle1(input []string) int {
	res := 0
	data := importData(input)
	for i := 0; i < days1; i++ {
		res += simulateStep(data)
	}
	return res
}

func puzzle2(input []string) int {
	day := 0
	data := importData(input)
	nbOctopus := len(input) * len(input[0])
	for simulateStep(data) < nbOctopus {
		day++
	}
	day++ // the day they all flashed
	return day
}

func importData(input []string) [][]int {
	res := make([][]int, len(input))
	for i := range res {
		res[i] = make([]int, len(input[i]))
		for j, n := range input[i] {
			res[i][j] = int(n - '0')
		}
	}
	return res
}

// if point is inside the matrix (if it is not out of bounds)
func isInside(data [][]int, point Coord) bool {
	return point.i >= 0 && point.i < len(data) && point.j >= 0 && point.j < len(data[point.i])
}

func getNeighbours(data [][]int, point Coord) []Coord {
	res := make([]Coord, 9)
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if isInside(data, Coord{point.i + i, point.j + j}) {
				res[count] = Coord{point.i + i, point.j + j}
				count++
			}
		}
	}
	return res[:count]
}

func exploreNeighbours(data [][]int, start Coord) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			next := Coord{start.i + i, start.j + j}
			if isInside(data, next) && data[next.i][next.j] != 0 {
				recExplore(data, next)
			}
		}
	}
}

func recExplore(data [][]int, start Coord) {
	data[start.i][start.j]++
	if data[start.i][start.j] > 9 {
		data[start.i][start.j] = 0
		// neighbours := getNeighbours(data, start)
		// for i := range neighbours {
		// 	if data[neighbours[i].i][neighbours[i].j] != 0 {
		// 		recExplore(data, neighbours[i])
		// 	}
		// }
		exploreNeighbours(data, start)
	}
}

func iterativeFlash(data [][]int) int {
	nbFlash := 0
	stop := false
	for !stop {
		stop = true
		for i := range data {
			for j := range data[i] {
				if data[i][j] > 9 {
					nbFlash++
					data[i][j] = 0
					// for _, n := range getNeighbours(data, Coord{i, j}) {
					// 	if data[n.i][n.j] != 0 {
					// 		data[n.i][n.j]++
					// 		stop = false
					// 	}
					// }
					// if incrementNeighbours(data, Coord{i, j}) {
					// 	stop = false
					// }
					stop = !incrementNeighbours(data, Coord{i, j}) && stop // evaluation ordre is important!
				}
			}
		}
	}
	return nbFlash
}

func incrementNeighbours(data [][]int, point Coord) bool {
	incremented := false
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			next := Coord{point.i + i, point.j + j}
			if isInside(data, next) && data[next.i][next.j] != 0 {
				data[next.i][next.j]++
				incremented = true
			}
		}
	}
	return incremented
}

func flash(data [][]int) int {
	nbFlash := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] > 9 {
				recExplore(data, Coord{i, j})
			}
		}
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 0 {
				nbFlash += 1
			}
		}
	}
	return nbFlash
}

func simulateStep(data [][]int) int {
	for i := range data {
		for j := range data[i] {
			data[i][j]++
		}
	}
	return flash(data)
}
