package main

import (
	"aoc2021/utils"
	"testing"
)

var input = utils.ReadInts("./input")
var testInput = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestPuzzle1(t *testing.T) {
	res := countIncreasing(testInput, 1)
	const expected int = 7
	if expected != res {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := countIncreasing(testInput, 3)
	const expected int = 5
	if expected != res {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countIncreasing(input, 1)
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countIncreasing(input, 3)
	}
}
