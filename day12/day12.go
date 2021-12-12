package main

import (
	"fmt"
	"aoc2021/utils"
	"strings"
)

const (
	maxPathLength = 1024 // maximum lengh of path. If path during exploration exceeds this number, the programm will fail
)

func main() {
	fmt.Println("day12")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", puzzle1(input))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", puzzle2(input))
}

type path []string
type graph map[string][]string

func importData(input []string) graph {
	res := make(graph)
	for _, line := range input {
		caves := strings.SplitN(line, "-", 2)
		_, exist := res[caves[1]]
		if !exist {
			res[caves[1]] = make([]string, 0)
		}
		res[caves[1]] = append(res[caves[1]], caves[0])
		if _, exist := res[caves[0]]; !exist {
			res[caves[0]] = make([]string, 0)
		}
		res[caves[0]] = append(res[caves[0]], caves[1])
	}
	return res
}

func countPaths(graph graph, allowed_repetition int) int {
	res := 0
	p := make(path, maxPathLength)
	p[0] = "start"
	visited := make(map[string]int)
	recCountPaths(graph, p, 1, visited, &res, allowed_repetition, false)

	return res
}

// count the number of paths possible that complete path to "end" node. Works with allowed_repetition in {1, 2} only.
// there is a line that can be changed if you want to use higher allowed_repetition
// only one lowercase node is allowed to have mulitple repetitions
func recCountPaths(graph graph, p path, next int, visited map[string]int, count *int, allowed_repetition int, smallExploredMultipleTimes bool) {
	accessible := graph[p[next-1]] // the nodes that are accessible as continuation of the path
	for i := range accessible {
		foundExploredMultipleTimes := false 
		if accessible[i] == "start" {
			continue
		}
		if accessible[i] == "end" {
			*count++
			continue
		}
		if isLower(accessible[i]) && visited[accessible[i]] > 0 { 
			// if smallExploredMultipleTimes && visited[accessible[i]] <= 1 || visited[accessible[i]] >= allowed_repetition { // can work if we can repeat more than 2 times lowercase nodes
			if smallExploredMultipleTimes || visited[accessible[i]] >= allowed_repetition { // works with allowed_repetition in {1, 2} only. But in those cases faster
				continue
			}
			foundExploredMultipleTimes = true // we have to set smallExploredMultipleTimes this way, so that we can reset it for the next iterator of the loop
		}
		p[next] = accessible[i] // append to path
		visited[accessible[i]]++ // keep track of how many times this node has been visited (only useful for lowercase nodes)
		recCountPaths(graph, p, next+1, visited, count, allowed_repetition, smallExploredMultipleTimes || foundExploredMultipleTimes)
		visited[accessible[i]]--
	}
}

func puzzle1(input []string) int {
	return countPaths(importData(input), 1)
}
func puzzle2(input []string) int {
	return countPaths(importData(input), 2)
}

func isLower(point string) bool {
	return strings.ToLower(point) == point
}
