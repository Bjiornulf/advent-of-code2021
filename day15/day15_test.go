package main

import (
	"testing"
	"aoc2021/utils"
)

var input = utils.ReadLines("input")

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
