package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("day3")
	puzzle1()
	puzzle2()
}

func puzzle1() {
	var lines []string = utils.ReadLines("input")
	var gamma, epsilon int
	for i := range lines[0] {
		// calculating binary representation value on the fly
		// using a0a1a2a3... = (...(((a0 * b) + a1 * b) + a2 * b)...) where b is the base of a0a1...
		gamma *= 2
		epsilon *= 2
		if 2*countOnesAtPos(lines, i) >= len(lines) { // majority of 1s at position i
			gamma++
		} else {
			epsilon++
		}
	}
	fmt.Printf("puzzle1: %v\n", gamma * epsilon)
}

func puzzle2() {
	var lines = utils.ReadLines("input")
	var oxygen, co2 []string = lines, lines // copying lines; oxygen, co2 and lines are independant
	for i := 0; i < len(lines[0]) && len(oxygen) > 1; i++ {
		ones := countOnesAtPos(oxygen, i)
		// filtering the strings. If in the list, most strings have a "1" at position i,
		// we only want to keep the strings that have "1" at position i
		if 2*ones >= len(oxygen) {
			oxygen = utils.StrFilter(oxygen, func(val string) bool {
				return val[i] == '1'
			})
		} else {
			oxygen = utils.StrFilter(oxygen, func(val string) bool {
				return val[i] == '0'
			})
		}
	}
	for i := 0; i < len(lines[0]) && len(co2) > 1; i++ {
		ones := countOnesAtPos(co2, i)
		// filtering the strings. If in the list, most strings have a "1" at position i,
		// we only want to keep the strings that have "0" at position i
		if 2*ones >= len(co2) {
			co2 = utils.StrFilter(co2, func(val string) bool {
				return val[i] == '0'
			})
		} else {
			co2 = utils.StrFilter(co2, func(val string) bool {
				return val[i] == '1'
			})
		}
	}
	oxygenValue, _ := strconv.ParseInt(oxygen[0], 2, 0)
	co2Value, _ := strconv.ParseInt(co2[0], 2, 0)
	fmt.Printf("puzzle2: %v\n", oxygenValue * co2Value)
}

// Counts the number of "1" at position pos in the array of strings
func countOnesAtPos(values []string, pos int) (count int) {
	for _, val := range values {
		if pos < len(val) && val[pos] == '1' {
			count++
		}
	}
	return
}
