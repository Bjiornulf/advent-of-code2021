package main

import (
	"aoc2021/utils"
	"fmt"
	"sort"
	"strings"
)

var corruptionScores map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScores map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var corresponding map[rune]rune = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}

func main() {
	fmt.Println("day10")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", puzzle1(input))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", puzzle2(input))
}

// return the characters that have not been matched and the first corrupted character
// if there is a corrupted character, the unmatched characters are not returned
// if there is not corrupted character, 0 is returned instead
func processLine(s string) ([]rune, rune) {
	open := []rune{}
	for _, l := range s {
		if strings.ContainsRune("({[<", l) {
			open = append(open, l)
			continue
		}
		if corresponding[l] != open[len(open)-1] {
			return []rune{}, l
		}
		open = open[:len(open)-1]
	}
	return open, 0
}

// first illegal character else 0
func findFirstIllegal(s string) rune {
	_, res := processLine(s)
	return res
}

// characters that have not been matched, else empty []rune
func getToComplete(s string) []rune {
	open, _ := processLine(s)
	return open
}

func puzzle1(input []string) int {
	score := 0
	for _, s := range input {
		r := findFirstIllegal(s)
		if r != 0 {
			score += corruptionScores[r]
		}
	}
	return score
}

func puzzle2(input []string) int {
	res := []int{}
	for _, s := range input {
		score := 0
		toComplete := getToComplete(s)
		for i := range toComplete {
			score *= 5
			score += completionScores[toComplete[len(toComplete)-1-i]] // completing last first
		}
		if score != 0 {
			res = append(res, score)
		}
	}
	// we want the median. There are more efficient ways to obtain it, but this is the fastest to implement
	sort.Ints(res)
	return res[len(res)/2]
}
