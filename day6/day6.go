package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	days  int = 80
	days2     = 256
)

func main() {
	fmt.Println("day6")
	input := utils.ReadLines("input")
	lantersFishes := importData(input)
	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", sum(simulate(lantersFishes, days)))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", sum(simulate(lantersFishes, days2)))
}

// sum of int array
func sum(a []int) int {
	var sum int
	for _, i := range a {
		sum += i
	}
	return sum
}

// import data as array containing the number of fishes in each state at the beginning
func importData(input []string) []int {
	numbers := strings.Split(input[0], ",")
	state := make([]int, 9)
	// fill array
	for i := range numbers {
		fish, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic(err)
		}
		state[fish]++
	}
	return state
}

// given an array with the number of fishes in each state. ARRAY MUST BE LENGTH 9
// simulates the evolution of the population
// returns an array of the number of fishes at each state, does not change the first array
func simulate(state []int, days int) []int {
	res := make([]int, 9)
	copy(res, state) // avoid changing initial array
	for i := 0; i < days; i++ {
		day0 := res[0]
		for j := 1; j < 9; j++ {
			res[j-1] = res[j]
		}
		res[8] = day0
		res[6] += day0

	}
	return res
}
