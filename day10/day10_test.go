package main

import (
	"testing"
)

var testInput = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestPuzzle1(t *testing.T) {
	res := puzzle1(testInput)
	expected := 26397
	if (res != expected) {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}

func TestPuzzle2(t *testing.T) {
	res := puzzle2(testInput)
	expected := 288957
	if (res != expected) {
		t.Errorf("Puzzle1: expected %v but got %v\n", expected, res)
	}
}
