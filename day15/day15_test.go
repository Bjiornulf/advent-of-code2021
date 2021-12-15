package main

import (
	"aoc2021/utils"
	"testing"
)

var input = utils.ReadLines("input")
var testInput = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	expected := 40
	if res != expected {
		t.Errorf("Puzzle1: exptected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	expected := 315
	if res != expected {
		t.Errorf("Puzzle2: exptected %v but got %v\n", expected, res)
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
