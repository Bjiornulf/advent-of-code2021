package main

import (
	"aoc2021/utils"
	"testing"
)


func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(simulate(importData(utils.ReadLines("input")), days2))
	}
}

const testInput = "3,4,3,1,2"

func TestImport(t *testing.T) {
	res := importData([]string{testInput})
	expected := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}
	equal := true
	for i := range expected {
		if i >= len(res) {
			t.Errorf("Something went wrong importing data, too many values imported\n")
			return
		}
		equal = equal && expected[i] == res[i]
	}
	if !equal {
		t.Errorf("Importing test data: expcted %v but got %v\n", expected, res)
	}
}

func TestSim18(t *testing.T) {
	testData := importData([]string{testInput})
	res := sum(simulate(testData, 18))
	const expected = 26
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}

func TestSim80(t *testing.T) {
	testData := importData([]string{testInput})
	res := sum(simulate(testData, 80))
	const expected = 5934
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}

func TestSim256(t *testing.T) {
	testData := importData([]string{testInput})
	res := sum(simulate(testData, 256))
	const expected = 26984457539
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}
