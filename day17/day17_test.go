package main

import (
	"aoc2021/utils"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

const testInput = "target area: x=20..30, y=-10..-5"
const input = "target area: x=179..201, y=-109..-63"

func TestPuzzle1_1(t *testing.T) {
	res := puzzle1([]string{testInput})
	expected := 45
	if expected != res {
		t.Errorf("Puzzle 1 - test input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle1_2(t *testing.T) {
	res := puzzle1([]string{input})
	expected := 5886
	if expected != res {
		t.Errorf("Puzzle 1 - my input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2_1(t *testing.T) {
	res := puzzle2([]string{testInput})
	expected := 112
	if expected != res {
		t.Errorf("Puzzle 2 - test input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2_2(t *testing.T) {
	res := puzzle2([]string{input})
	expected := 1806
	if expected != res {
		t.Errorf("Puzzle 2 - my input: expected %v but got %v\n", expected, res)
	}
}

func BenchmarkSscanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x1, x2, y1, y2 int
		fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	}
}

func BenchmarkFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x1, x2, y1, y2 int
		vals := strings.FieldsFunc(input, func(c rune) bool {
			return !unicode.IsNumber(c) && c != '-'
		})
		x1, _ = strconv.Atoi(vals[0])
		x2, _ = strconv.Atoi(vals[1])
		y1, _ = strconv.Atoi(vals[2])
		y2, _ = strconv.Atoi(vals[3])
		_, _, _, _ = x1, x2, y1, y2
	}
}

func BenchmarkUtilsMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IntMin(rand.Int(), rand.Int())
	}
}

func BenchmarkMin(b *testing.B) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < b.N; i++ {
		min(rand.Int(), rand.Int())
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle1([]string{input})
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2([]string{input})
	}
}
