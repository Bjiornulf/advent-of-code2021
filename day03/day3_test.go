package main

import (
	"testing"
)

var testInput []string = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	const expected = 198
	if res != expected {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	const expected = 230
	if res != expected {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}
