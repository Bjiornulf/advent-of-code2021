package main

import (
	"fmt"
	"aoc2021/utils"
)

func main() {
	fmt.Println("day7")
	input := utils.ReadLines("input")
	crabs := utils.ReadIntsList(input[0], ",")
	fuelFuction1 := func(c, t int) int {return utils.IntAbs(c-t)}
	fuelFuction2 := func(c, t int) int {
		n := utils.IntAbs(c-t)
		return n*(n+1) / 2
	}
	fmt.Println("Puzzle1:", findMinFuel(crabs, fuelFuction1))
	fmt.Println("Puzzle2:", findMinFuel(crabs, fuelFuction2))
}

func findMinFuel(crabs []int, fuelFuction func(int, int) int) int {
	xmax := utils.IntMax(crabs...)
	fuels := make([]int, xmax)
	for i := range fuels {
		fuels[i] = calcTotalFuelConsumption(crabs, i, fuelFuction)
	}
	return utils.IntMin(fuels...)
}

func calcTotalFuelConsumption(crabs []int, target int, fuelFuction func(int, int) int) int {
	fuelSum := 0
	for _, crab := range crabs {
		fuelSum += fuelFuction(crab, target)
	}
	return fuelSum
}
