package main

import (
	"aoc2021/utils"
	"testing"
)

var testInput1 = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

var testInput2 = []string{
	"dc-end",
	"HN-start",
	"start-kj",
	"dc-start",
	"dc-HN",
	"LN-dc",
	"HN-end",
	"kj-sa",
	"kj-HN",
	"kj-dc",
}

var testInput3 = []string{
	"fs-end",
	"he-DX",
	"fs-he",
	"start-DX",
	"pj-DX",
	"end-zg",
	"zg-sl",
	"zg-pj",
	"pj-he",
	"RW-he",
	"fs-DX",
	"pj-RW",
	"zg-RW",
	"start-pj",
	"he-WI",
	"zg-he",
	"pj-fs",
	"start-RW",
}

var input = utils.ReadLines("input")

func TestPuzzle1_1(t *testing.T) {
	res := puzzle1(testInput1)
	expected := 10
	if res != expected {
		t.Errorf("Puzzle1 small input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle1_2(t *testing.T) {
	res := puzzle1(testInput2)
	expected := 19
	if res != expected {
		t.Errorf("Puzzle1 medium input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle1_3(t *testing.T) {
	res := puzzle1(testInput3)
	expected := 226
	if res != expected {
		t.Errorf("Puzzle1 big input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2_1(t *testing.T) {
	res := puzzle2(testInput1)
	expected := 36
	if res != expected {
		t.Errorf("Puzzle2 small input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2_2(t *testing.T) {
	res := puzzle2(testInput2)
	expected := 103
	if res != expected {
		t.Errorf("Puzzle2 medium input: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2_3(t *testing.T) {
	res := puzzle2(testInput3)
	expected := 3509
	if res != expected {
		t.Errorf("Puzzle2 big input: expected %v but got %v\n", expected, res)
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
