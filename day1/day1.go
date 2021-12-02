package main

import (
	"fmt"
	"aoc2021/utils"
)

func main() {
	fmt.Println("day1")
	ints := utils.ReadInts("./input");
	fmt.Printf("Puzzle1: %v\n", countIncreasing(ints));
	groupedInts := make([]int, len(ints) - 2);
	for i := 2; i < len(ints); i++ {
		groupedInts[i-2] = ints[i-2] + ints[i-1] + ints[i];
	}
	fmt.Printf("Puzzle2: %v\n", countIncreasing(groupedInts));
}

func countIncreasing(values []int) int {
	var count int = 0;
	for i := 1; i < len(values); i++ {
		if (values[i] - values[i-1] > 0) {
			count++;
		}
	}
	return count;
}
