package main

import (
	"fmt"
	"aoc2021/utils"
	"strings"
)

// converting the digits to bitmaps will allow us to compare them in constant time
// this is hard otherwise, since the letters are not always in the same order
var compareLookup map[rune]uint8 = map[rune]uint8 {
	'a': 0b00000001,
	'b': 0b00000010,
	'c': 0b00000100,
	'd': 0b00001000,
	'e': 0b00010000,
	'f': 0b00100000,
	'g': 0b01000000,
}

const inputFile = "input"
func main() {
	fmt.Println("day8")
	input := utils.ReadLines(inputFile)

	/* ---------- Puzzle 1 ---------- */
	fmt.Println("Puzzle1:", puzzle1(input))

	/* ---------- Puzzle 2 ---------- */
	fmt.Println("Puzzle2:", puzzle2(input))
}

func importData(input []string) ([][]string, [][]string){
	data := make([][]string, len(input))
	reads := make([][]string, len(input))
	for i := range input {
		line := strings.Split(input[i], " ")
		data[i] = line[:10]
		reads[i] = line[11:]
	}
	return data, reads
}

// deduce the number encoded on a line
func deduceNumber(data, read []string) int {
	decode := make(map[uint8]int)
	var one, four, seven uint8
	// solve the trivial cases
	for _, d := range data {
		if len(d) == 2 {
			one = byteRepr(d)
			decode[byteRepr(d)] = 1
		} else if len(d) == 3 {
			seven = byteRepr(d)
			decode[byteRepr(d)] = 7
		} else if len(d) == 4 {
			four = byteRepr(d)
			decode[byteRepr(d)] = 4
		} else if len(d) == 7 {
			decode[byteRepr(d)] = 8
		}
	}

	// this is the part of the 4 that is not in common with the 1
	sideFour := four & (^one) // ^ = bitwise not : ^0b0101 = 0b1010

	for _, d := range data {
		if len(d) == 5 {
			if containsAllByte(d, sideFour) {
				decode[byteRepr(d)] = 5	// 5 is the only digit that contains a part of 4 and has 5 segments
			} else if containsAllByte(d, seven) {
				decode[byteRepr(d)] = 3 // 3 is the only digit that contains 7 and has 5 segments
			} else {
				decode[byteRepr(d)] = 2 // 2 is the only other digit that has 5 segments
			}
		} else if len(d) == 6 {
			if containsAllByte(d, four) {
				decode[byteRepr(d)] = 9 // 9 is the only digit that contains 4 and has 6 segments
			} else if containsAllByte(d, seven) {
				decode[byteRepr(d)] = 0 // 0 is the only digit that does not contain 4 but contains 7 and has 6 segments
			} else {
				decode[byteRepr(d)] = 6 // 6 is the only other digit that has 6 segments
			}
		}
	}

	// using hoerner method to calculate a base10 representation of the numbers at the end
	res := 0
	for _, r := range read {
		res *= 10
		res += decode[byteRepr(r)]
	}

	return res
}

// Solution to puzzle1 : count the number of digits we are certain of
func puzzle1(input []string) int {
	_, reads := importData(input)
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

func puzzle2(input []string) int {
	data, reads := importData(input)
	_ = data
	res := 0
	for i := range reads {
		res += deduceNumber(data[i], reads[i])
	}
	return res
}

// Convert string representing a digit into a byte for easier comparaison
func byteRepr(s string) uint8 {
	res := uint8(0)
	for _, l := range s {
		res |= compareLookup[l]
	}
	return res
}

func containsAllByte(s string, b uint8) bool {
	return (byteRepr(s) & b == b)
}
