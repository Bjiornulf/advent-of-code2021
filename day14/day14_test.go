package main

import (
	"aoc2021/utils"
	"testing"
)

var input = utils.ReadLines("input")

var testInput = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	expected := 1588
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	expected := 2188189693529
	if res != expected {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
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
