package main

import (
	"testing"
)

var testInput []string = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	const expected = 150
	if res != expected {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	const expected = 900
	if res != expected {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}
