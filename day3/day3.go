package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("day3")
	lines := utils.ReadLines("input");
	comm := make([]int, len(lines[0]))
	for i := range lines[0] {
		comm[i] = countOnesAtPos(lines, i);
	}
	gamma, epsilon := "", "";
	for _, val := range comm {
		if 2*val >= len(lines) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}	
	}
	fmt.Println(comm);
	g, _ := strconv.ParseInt(gamma, 2, 0);
	e, _ := strconv.ParseInt(epsilon, 2, 0);
	fmt.Printf("puzzle1: %v\n", e *g );
	oxy := lines;
	co := lines;
	for i1 := 0; i1 < len(lines[0]) && len(oxy) > 1; i1++ {
		ones := countOnesAtPos(oxy, i1);
		if (ones >= len(oxy) - ones) {
			oxy = utils.Filter(oxy, func(val string) bool {
				return val[i1] == '1'})
		} else {
			oxy = utils.Filter(oxy, func(val string) bool {
				return val[i1] == '0'})
		}	
	}	
	for i1 := 0; i1 < len(lines[0]) && len(co) > 1; i1++ {
		ones := countOnesAtPos(co, i1)
		if (ones >= len(co) - ones) {
			co = utils.Filter(co, func(val string) bool {
				return val[i1] == '0'})
		} else {
			co = utils.Filter(co, func(val string) bool {
				return val[i1] == '1'})
		}
	}	
	o, _ := strconv.ParseInt(oxy[0], 2, 0);
	c, _ := strconv.ParseInt(co[0], 2, 0);
	fmt.Printf("puzzle2: %v\n", o *c );
}

func countOnesAtPos(values []string, position int) int {
	count := 0;
	for _, val := range values {
		if val[position] == '1' {
			count++
		}
	}
	return count
}
