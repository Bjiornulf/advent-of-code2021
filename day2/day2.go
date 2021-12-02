package main

import (
	"fmt"
	"strings"
	"strconv"
	"aoc2021/utils"
)

func main() {
	fmt.Println("day2")
	puzzle1();
	puzzle2();
}

func puzzle1() {
	commands := utils.ReadLines("./input");
	var depth, position int;
	for _, line := range commands {
		command := strings.SplitN(line, " ", 2);
		val, err := strconv.Atoi(command[1]);
		if (err != nil) {
			panic(err);
		}
		switch command[0] {
			case "forward":
				position+=val;
			case "down":
				depth+=val;
			case "up":
				depth-=val;
		}
	}
	fmt.Printf("puzzle1: %v\n", depth*position);
}

func puzzle2() {
	commands := utils.ReadLines("./input");
	var depth, position, aim int;
	for _, line := range commands {
		command := strings.SplitN(line, " ", 2);
		val, err := strconv.Atoi(command[1]);
		if (err != nil) {
			panic(err);
		}
		switch command[0] {
			case "forward":
				position+=val;
				depth += val * aim
			case "down":
				aim+=val;
			case "up":
				aim-=val;
		}
	}
	fmt.Printf("puzzle2: %v\n", depth*position);
}
