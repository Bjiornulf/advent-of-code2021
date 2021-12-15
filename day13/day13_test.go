package main

import (
	"aoc2021/utils"
	"testing"
)

var input []string = utils.ReadLines("input")
var testInput = []string{
	"6,10",
	"0,14",
	"9,10",
	"0,3",
	"10,4",
	"4,11",
	"6,0",
	"6,12",
	"4,1",
	"0,13",
	"10,12",
	"3,4",
	"3,0",
	"8,4",
	"1,10",
	"2,14",
	"8,10",
	"9,0",
	"",
	"fold along y=7",
	"fold along x=5",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	expected := 17
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func BenchmarkImportData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		importData(input)
	}
}

func BenchmarkGetMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMax(input)
	}
}
