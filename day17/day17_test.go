package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
	"math/rand"
)

const input = "target area: x=179..201, y=-109..-63"

func BenchmarkSscanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x1, x2, y1, y2 int
		fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	}
}

func BenchmarkFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x1, x2, y1, y2 int
		vals := strings.FieldsFunc(input, func (c rune) bool {
			return !unicode.IsNumber(c) && c != '-'
		})
		x1, _ = strconv.Atoi(vals[0])
		x2, _ = strconv.Atoi(vals[1])
		y1, _ = strconv.Atoi(vals[2])
		y2, _ = strconv.Atoi(vals[3])
		_, _, _, _ = x1, x2, y1, y2
	}
}

func BenchmarkUtilsMin (b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IntMin(rand.Int(), rand.Int())
	}
}

func BenchmarkMin (b *testing.B) {
	min := func (a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < b.N; i++ {
		min(rand.Int(), rand.Int())
	}
}
