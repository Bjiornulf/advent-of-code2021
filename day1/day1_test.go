package main

import (
	"testing"
)

var testData = []int{
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
	res := countIncreasing(testData, 1)
	const expected int = 7
	if expected != res {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := countIncreasing(testData, 3)
	const expected int = 5
	if expected != res {
		t.Errorf("Expected %v but got %v\n", expected, res)
	}
}
