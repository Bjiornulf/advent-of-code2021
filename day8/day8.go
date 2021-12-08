package main

import (
	"fmt"
	"aoc2021/utils"
	"strings"
)
const inputFile = "input"
func main() {
	fmt.Println("day8")
	data, reads := importData()
	_ = data
	fmt.Println("Puzzle1:", countUniques(reads))
	puz2 := 0
	for i := range reads {
		puz2 += deduceNumber(data[i], reads[i])
		fmt.Println(reads[i])
		fmt.Println(deduceNumber(data[i], reads[i]))
	}
	fmt.Println(puz2)
}

func importData() ([][]string, [][]string){
	input := utils.ReadLines(inputFile)	
	data := make([][]string, len(input))
	reads := make([][]string, len(input))
	for i := range input {
		line := strings.Split(input[i], " ")
		data[i] = line[:10]
		reads[i] = line[11:]
	}
	return data, reads
}

func deduceNumber(data, read []string) int {
	decode := make(map[string]int)
	var one, three, four, seven, eight string
	_, _ = seven, eight
	for _, d := range data {
		if len(d) == 2 {
			one = d
			decode[d] = 1
		} else if len(d) == 3 {
			seven = d
			decode[d] = 7
		} else if len(d) == 4 {
			four = d
			decode[d] = 4
		} else if len(d) == 7 {
			eight = d
			decode[d] = 8
		}
	}
	// var top string
	// for _, l7 := range seven {
	// 	found := false
	// 	for _, l1 := range one {
	// 		if l1 == l7 {
	// 			found = true
	// 		}
	// 	}
	// 	if !found {
	// 		top = fmt.Sprintf("%c", l7)
	// 	}
	// }

	var sideFour string
	letters := []rune{}
	for _, l4 := range four {
		found := false
		for _, l1 := range one {
			if l1 == l4 {
				found = true
			}
		}
		if !found {
			letters = append(letters, l4)
		}
	}
	sideFour = fmt.Sprintf("%c%c", letters[0], letters[1])
	for _, d := range data {
		if len(d) == 6 && !containsAll(d, one) {
			decode[d] = 6
		} else if len(d) == 5 && containsAll(d, sideFour) {
			decode[d] = 5
		} else if len(d) == 5 && containsAll(d, one) {
			decode[d] = 3
			three = d
		} else if len(d) == 5 {
			decode[d] = 2
		}
	}
	for _, d := range data {
		if len(d) == 6 && containsAll(d, three) {
			decode[d] = 9
		} else {
			_, ok := decode[d]
			if len(d) == 6 && !ok {
				decode[d] = 0
			}
		}
	}
	res := 0
	for _, r := range read {
		for k, v := range decode {
			if (equalByRune(k, r)) {
				res *= 10
				res += v
			}
		}
	}

	return res
}

func countUniques(reads [][]string) int {
	var count int
	for i := range reads {
		for j := range reads[i] {
			if len(reads[i][j]) >= 2 && len(reads[i][j]) <= 4 || len(reads[i][j]) == 7 {
				count++
			}
		}
	}
	return count
}

func containsAll(s, sub string) bool {
	res := true
	for _, r := range sub {
		res = res && strings.ContainsRune(s, r)
	}
	return res
}

func equalByRune(s, cmp string) bool {
	if len(s) != len(cmp) {
		return false
	}
	found := 0
	for _, c := range cmp {
		if strings.ContainsRune(s, c) {
			found++
		}
	}
	return found == len(s)
}
