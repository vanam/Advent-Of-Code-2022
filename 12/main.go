package main

import (
	"bufio"
	"bytes"
	"container/list"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var input []byte

type vertex struct {
	i   int
	j   int
	len int
}

var directions = [4][2]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	var grid []string
	var start vertex

	for scanner.Scan() {
		line := scanner.Text()
		for j, c := range line {
			if c == 'S' {
				start = vertex{i: len(grid), j: j}
				break
			}
		}
		grid = append(grid, line)
	}

	return shortestPath(grid, start)
}

func partTwo(scanner *bufio.Scanner) int {
	var grid []string
	var start []vertex

	for scanner.Scan() {
		line := scanner.Text()
		for j, c := range line {
			if c == 'S' || c == 'a' {
				start = append(start, vertex{i: len(grid), j: j})
			}
		}
		grid = append(grid, line)
	}

	minShortestPath := math.MaxInt32
	for _, s := range start {
		sp := shortestPath(grid, s)
		if sp != -1 && sp < minShortestPath {
			minShortestPath = sp
		}
	}

	return minShortestPath
}

func shortestPath(grid []string, start vertex) int {
	seen := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		seen[i] = make([]bool, len(grid[0]))
	}

	queue := list.New()
	queue.PushBack(start)
	seen[start.i][start.j] = true

	for queue.Len() > 0 {
		u := queue.Front().Value.(vertex)
		queue.Remove(queue.Front())

		uc := grid[u.i][u.j]

		if uc == 'E' {
			return u.len
		} else if uc == 'S' {
			uc = 'a'
		}

		for _, d := range directions {
			vi := u.i + d[0]
			vj := u.j + d[1]

			if vi < 0 || vj < 0 || vi >= len(grid) || vj >= len(grid[0]) || seen[vi][vj] {
				continue
			}

			vc := grid[vi][vj]
			if vc == 'E' {
				vc = 'z'
			}

			dc := rune(vc) - rune(uc)
			if dc <= 1 {
				queue.PushBack(vertex{vi, vj, u.len + 1})
				seen[vi][vj] = true
			}
		}
	}

	return -1 // no path found
}
