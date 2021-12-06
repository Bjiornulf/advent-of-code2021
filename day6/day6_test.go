package main

import (
	"testing"
)

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(simulate(importData(), days2))
	}
}

var testData = []int{0, 1, 1, 2, 1, 0, 0, 0, 0}

func TestSim18(t *testing.T) {
	res := sum(simulate(testData, 18))
	const expected = 26
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}

func TestSim80(t *testing.T) {
	res := sum(simulate(testData, 80))
	const expected = 5934
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}

func TestSim256(t *testing.T) {
	res := sum(simulate(testData, 256))
	const expected = 26984457539
	if (res != expected) {
		t.Errorf("Simulating 18 days: expected %d but got %d\n", expected, res)
	}
}
