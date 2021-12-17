package main

import (
	"aoc2021/utils"
	"testing"
)

var testInput []string = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

var input []string = utils.ReadLines("input")

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	const expected = 15
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	const expected = 1134
	if res != expected {
		t.Errorf("Puzzle2: expected %v but got %v\n", expected, res)
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle1(input)
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2(input)
	}
}
