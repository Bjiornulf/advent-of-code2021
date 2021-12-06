package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
	"strconv"
)

const (
	days int = 80
	days2 	 = 256
)
func main() {
	fmt.Println("day6")
	/* ---------- Puzzle 1 ---------- */
	lantersFishes1 := importData();
	fmt.Println("Puzzle1:", sum(simulate(lantersFishes1, days)))

	/* ---------- Puzzle 2 ---------- */
	lantersFishes2 := importData();
	fmt.Println("Puzzle2:", sum(simulate(lantersFishes2, days2)))
}

// sum of int array
func sum(a []int) int {
	var sum int
	for _, i := range a {
		sum+=i
	}
	return sum
}

// import data as array containing the number of fishes in each state at the beginning
func importData() []int {
	line := utils.ReadLines("input")
	numbers := strings.Split(line[0], ",")
	fishes := make([]int, len(numbers))
	// fill array
	for i := range fishes {
		fish, err := strconv.Atoi(numbers[i])
		if (err != nil) {
			panic(err)
		}
		fishes[fish]++
	}
	return fishes
}


// given an array with the number of fishes in each state. ARRAY MUST BE LENGTH 9
// simulates the evolution of the population
// returns an array of the number of fishes at each state
func simulate(fishes []int, days int) []int {
	for i := 0; i < days; i++ {
		fishes0 := fishes[0]
		for j := 1; j < 9; j++ {
			fishes[j-1] = fishes[j]
		}
		fishes[8] = fishes0
		fishes[6] += fishes0
		
	}
	return fishes
}
