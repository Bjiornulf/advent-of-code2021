package main

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func main() {
	fmt.Println("day1")
	ints := utils.ReadInts("./input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Printf("Puzzle1: %v\n", countIncreasing(ints, 1))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Printf("Puzzle2: %v\n", countIncreasing(ints, 3))
	fmt.Println(time.Since(start))
}

func countIncreasing(values []int, groupSize int) int {
	var count int = 0
	for i := groupSize; i < len(values); i++ {
		if values[i]-values[i-groupSize] > 0 {
			count++
		}
	}
	return count
}
