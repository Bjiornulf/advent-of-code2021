package main

import (
	"testing"
)

var testInput []string = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

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
