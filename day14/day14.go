package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("day14")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Println("Puzzle1", puzzle1(input))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Println("Puzzle2", puzzle2(input))
	fmt.Println(time.Since(start))
}

type Counter struct {
	val int
	acc int
}
type Steps map[string]string
type Polymer struct {
	repr map[string]*Counter
	steps Steps
	count map[string]int
}

func (c *Counter)incr(i int) {
	c.acc += i
}

func (c *Counter)reset() {
	c.val = c.acc
	c.acc = 0
}

func importData(input []string) *Polymer {
	polymer := input[0]
	steps := make(Steps)
	for _, line := range input[2:] {
		thing := strings.SplitN(line, " ", 3)
		steps[thing[0]] = thing[2]
	}
	return newPolymer(polymer, steps)
}

func newPolymer(poly string, steps Steps) *Polymer {
	res := new(Polymer)
	res.repr = make(map[string]*Counter)
	for i := 0; i < len(poly)-1; i++ {
		res.incr(poly[i:i+2], 1)
	}
	res.steps = steps
	res.count = make(map[string]int)
	for _, c := range poly {
		res.count[string(c)]++
	}
	for _, c := range res.repr {
		c.reset()
	}
	return res
}

func (p *Polymer)incr(literal string, i int) {
	if _, ok := p.repr[literal]; !ok {
		p.repr[literal] = new(Counter)
	}
	p.repr[literal].incr(i)
}

func (p *Polymer)get(literal string) int {
	return p.repr[literal].val
}

func (p *Polymer)step() {
	keys := make([]string, 0)
	for k := range p.repr {
		keys = append(keys, k)
	}
	for _, k := range keys {
		l1 := k[:1]+p.steps[k]
		l2 := p.steps[k]+k[1:]
		v := p.get(k)

		p.incr(l1, v)
		p.incr(l2, v)
		p.count[p.steps[k]] += v
	}
	for _, c := range p.repr {
		c.reset()
	}
}

func (p *Polymer)getMinMax() (int, int) {
	max := 0
	for _, v := range p.count {
		max = utils.IntMax(max, v)
	}
	min := max
	for _, v := range p.count {
		min = utils.IntMin(min, v)
	}
	return min, max
}

func puzzle2(input []string) int {
	poly := importData(input)
	for i := 0; i < 40; i++ {
		poly.step()
	}
	min, max := poly.getMinMax()
	return max - min
}

func puzzle1(input []string) int {
	poly := importData(input)
	for i := 0; i < 10; i++ {
		poly.step()
	}
	min, max := poly.getMinMax()
	return max - min
}
