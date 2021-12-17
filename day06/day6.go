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

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", puzzle1(input))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", puzzle2(input))
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
	for i := 0; i < days; i++ {
		state = append(state[1:], state[0]) // rotate the array left
		state[6] += state[8]
	}
	return state
}

func puzzle1(input []string) int {
	state := importData(input)
	state = simulate(state, days)
	return sum(state)
}

func puzzle2(input []string) int {
	state := importData(input)
	state = simulate(state, days2)
	return sum(state)
}
