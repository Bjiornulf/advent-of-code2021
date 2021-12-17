package main

import (
	"aoc2021/utils"
	"container/heap"
	"fmt"
	"time"
)

func main() {
	fmt.Println("day15")
	input := utils.ReadLines("input")

	/* ---------- Puzzle 1 ---------- */
	start := time.Now()
	fmt.Println("Puzzle1:", puzzle1(input))
	fmt.Println(time.Since(start))

	/* ---------- Puzzle 2 ---------- */
	start = time.Now()
	fmt.Println("Puzzle1:", puzzle2(input))
	fmt.Println(time.Since(start))
}

type Coord struct {
	i, j int
}

type Item struct {
	coord    *Coord
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq *PriorityQueue) Push(a interface{}) {
	n := len(*pq)
	item := a.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, coord *Coord, priority int) {
	item.coord = coord
	item.priority = priority
	heap.Fix(pq, item.index)
}

func importData(input []string) [][]uint8 {
	res := make([][]uint8, len(input))
	for i := range input {
		res[i] = make([]uint8, len(input[i]))
		for j := range input[i] {
			res[i][j] = input[i][j] - '0'
		}
	}
	return res
}

func importData2(input []string) [][]uint8 {
	res := make([][]uint8, len(input)*5)
	for i := range res {
		res[i] = make([]uint8, len(input[0])*5)
	}
	val := func(v int) int {
		if v == 0 {
			return 9
		}
		return v
	}
	// extend the initial matrix 5 times in each direction following given algorithm
	for i := range input {
		for j := range input[i] {
			for n := 0; n < 5; n++ {
				for m := 0; m < 5; m++ {
					res[i+len(input)*m][j+len(input[i])*n] = uint8(val((int(input[i][j]-'0') + m + n) % 9))
				}

			}
		}
	}
	return res
}

func isInside(graph [][]uint8, point Coord) bool {
	return point.i >= 0 && point.i < len(graph) && point.j >= 0 && point.j < len(graph[point.i])
}

func updateNeighbours(graph [][]uint8, visited [][]bool, start *Coord, lengths [][]int, pq *PriorityQueue) {
	visited[start.i][start.j] = true
	// generating neighbours and updating the length of the path since start to them
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			n := Coord{start.i + i, start.j + j}
			if isInside(graph, n) && !visited[n.i][n.j] && !(i != 0 && j != 0) { // !(i != 0 && j != 0) ignore diagonal links
				if lengths[start.i][start.j]+int(graph[n.i][n.j]) < lengths[n.i][n.j] || lengths[n.i][n.j] == -1 {
					// found a shorter path
					lengths[n.i][n.j] = lengths[start.i][start.j] + int(graph[n.i][n.j])
					// if the neighbour was updated, we add it to the priority queue to be later explored
					heap.Push(pq, &Item{
						coord:    &n,
						priority: lengths[n.i][n.j],
					})
				}
			}
		}
	}
}

// find the shortest path in a matrix graph using the Dijstra algorithm.
func shortestPath(graph [][]uint8, start Coord, end Coord) int {
	length := make([][]int, len(graph))
	// init length
	for i := range length {
		length[i] = make([]int, len(graph[i]))
		for j := range length[i] {
			length[i][j] = -1
		}
	}
	length[start.i][start.j] = 0
	visited := make([][]bool, len(graph))
	for i := range visited {
		visited[i] = make([]bool, len(graph[i]))
	}
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		&start,
		0,
		0,
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)
		toVisit := p.coord
		updateNeighbours(graph, visited, toVisit, length, &pq)
		if visited[end.i][end.j] {
			break
		}
	}
	return length[end.i][end.j]

}

func puzzle1(input []string) int {
	graph := importData(input)
	return shortestPath(graph, Coord{0, 0}, Coord{len(graph) - 1, len(graph[0]) - 1})
}

func puzzle2(input []string) int {
	graph := importData2(input)
	return shortestPath(graph, Coord{0, 0}, Coord{len(graph) - 1, len(graph[0]) - 1})
}
