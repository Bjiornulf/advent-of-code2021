package main

import (
	"testing"
)

var testInput = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	expected := 1656
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	expected := 195
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle1(testInput)
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2(testInput)
	}
}
