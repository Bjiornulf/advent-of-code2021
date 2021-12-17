package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("day2")
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

func puzzle1(commands []string) int {
	var depth, position int
	for _, line := range commands {
		command := strings.SplitN(line, " ", 2)
		val, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}
		switch command[0] {
		case "forward":
			position += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return depth * position
}

func puzzle2(commands []string) int {
	var depth, position, aim int
	for _, line := range commands {
		command := strings.SplitN(line, " ", 2)
		val, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}
		switch command[0] {
		case "forward":
			position += val
			depth += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return depth * position
}
