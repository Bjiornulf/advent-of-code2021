package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
	"time"
)

type pixel struct {
	curr, next bool
}
type image [][]pixel

func main() {
	fmt.Println("day20")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Println("Puzzle1:", puzzle1(input))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Println("Puzzle2:", puzzle2(input))
	fmt.Println(time.Since(start))
}

func importData(input []string, padding int) (image, string) {
	enhancements := input[0]
	image := make(image, len(input) - 2 + padding*2)
	for i := range image {
		image[i] = make([]pixel, len(input[2]) + padding*2)
	}
	for i := 2; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			image[i-2+padding][j+padding].curr = input[i][j]=='#'
		}
	}
	return image, enhancements
}

func getIndex(image image, i, j int) int {
	res := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			res *= 2
			if image[i+x][j+y].curr {
				res++
			}
		}
	}
	return res
}

func enhanceImage(image image, enhancements string, count int) {
	for i := 0; i < count; i++ {
		// calculate new image
		for i := 1; i < len(image)-1; i++ {
			for j := 1; j < len(image[i])-1; j++ {
				image[i][j].next = enhancements[getIndex(image, i, j)] == '#'
			}
		}
		// put result as current image
		for i := range image {
			for j := range image[i] {
				image[i][j].curr = image[i][j].next
			}
		}
	}
}

func countLit(image image, yStart, xStart int) int {
	count := 0
	for i := yStart; i < len(image)-yStart; i++ {
		for j := xStart; j < len(image[i])-xStart; j++ {
			if image[i][j].curr {
				count++
			}
		}
	}
	return count
}

func (image image)String() string {
	var buf strings.Builder
	for i := range image {
		for j := range image[i] {
			switch (image[i][j].curr) {
			case true:
				buf.WriteRune('#')
			case false:
				buf.WriteRune('.')
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func puzzle1(input []string) int {
	padding := 4
	image, enhancements := importData(input, padding)
	enhanceImage(image, enhancements, 2)
	return countLit(image, 2, 2)

}
func puzzle2(input []string) int {
	padding := 100
	image, enhancements := importData(input, padding)
	enhanceImage(image, enhancements, 50)
	return countLit(image, padding-50, padding-50)

}
