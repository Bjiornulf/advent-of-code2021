package main

import (
	"aoc2021/utils"
	"fmt"
)

func main() {
	fmt.Println("day7")
	input := utils.ReadLines("input")
	crabs := utils.ReadIntsList(input[0], ",")
	// cost to target when each step costs 1
	fuelFuction1 := func(c, t int) int { return utils.IntAbs(c - t) }
	// cost to target when each step increasingly costs 1 more : 1, 2, 3...
	fuelFuction2 := func(c, t int) int {
		n := utils.IntAbs(c - t)
		return n * (n + 1) / 2
	}

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", findMinFuel(crabs, fuelFuction1))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", findMinFuel(crabs, fuelFuction2))
}

// return the miniumum total fuel required so that every crab ends on the same position
func findMinFuel(crabs []int, fuelFuction func(int, int) int) int {
	xmax := utils.IntMax(crabs...)
	fuels := make([]int, xmax)
	for i := range fuels {
		fuels[i] = calcTotalFuelConsumption(crabs, i, fuelFuction)
	}
	return utils.IntMin(fuels...)
}

// calculate the total fuel cost to move every crab to the target
func calcTotalFuelConsumption(crabs []int, target int, fuelFuction func(int, int) int) int {
	fuelSum := 0
	for _, crab := range crabs {
		fuelSum += fuelFuction(crab, target)
	}
	return fuelSum
}
