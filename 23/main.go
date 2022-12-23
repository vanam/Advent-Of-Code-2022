package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var input []byte

type elf struct {
	x    int
	y    int
	next *elf
}

var di = 0
var dir = [][3][2]int{
	{{0, -1}, {1, -1}, {-1, -1}}, // N, NE, NW
	{{0, 1}, {1, 1}, {-1, 1}},    // S, SE, SW
	{{-1, 0}, {-1, -1}, {-1, 1}}, // W, NW, SW
	{{1, 0}, {1, -1}, {1, 1}},    // E, NE, SE
}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	elves := make([]elf, 0)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if line[x] != '#' {
				continue
			}
			elves = append(elves, elf{x, y, nil})
		}
		y++
	}

	for r := 1; r <= 2500; r++ {
		xBuckets := make(map[int][]elf)
		for _, e := range elves {
			xBuckets[e.x] = append(xBuckets[e.x], e)
		}

		proposedPositions := make(map[elf]int)

		for ie := 0; ie < len(elves); ie++ {
			e := &elves[ie]
			// find neighbours using simple space division
			possibleNeighbours := concatMultipleSlices([][]elf{xBuckets[e.x], xBuckets[e.x+1], xBuckets[e.x-1]})
			neighbours := make([]elf, 0)

			for _, pn := range possibleNeighbours {
				// do not consider self
				if pn.x == e.x && pn.y == e.y {
					continue
				}

				// not a neighbour
				if abs(pn.x-e.x) > 1 || abs(pn.y-e.y) > 1 {
					continue
				}
				neighbours = append(neighbours, pn)
			}

			// no neighbours?
			if len(neighbours) == 0 {
				// do nothing
				e.next = nil
				continue
			}
			var next *elf = nil
			for i := 0; i < len(dir); i++ {
				edi := (di + i) % len(dir)

				canMoveInD := true
				for _, d := range dir[edi] {
					if !canMoveInD {
						break
					}
					for _, n := range neighbours {
						if n.x == e.x+d[0] && n.y == e.y+d[1] {
							canMoveInD = false
							break
						}
					}
				}

				if canMoveInD {
					next = &elf{e.x + dir[edi][0][0], e.y + dir[edi][0][1], nil}
					proposedPositions[*next] = proposedPositions[*next] + 1
					e.next = next
					break
				}
			}

			e.next = next
		}

		// find who can move
		newElves := make([]elf, 0)
		elfMoved := false
		for _, e := range elves {
			if e.next != nil && proposedPositions[*e.next] == 1 {
				newElves = append(newElves, *e.next)
				elfMoved = true
			} else {
				newElves = append(newElves, e)
			}
		}

		// simulate until it stops
		if !elfMoved {
			if r <= 10 {
				partOne(elves)
			}
			fmt.Printf("part two: %v\n", r)
			break
		}

		elves = newElves

		di = (di + 1) % len(dir)

		if r == 10 {
			partOne(elves)
		}
	}
}

func partOne(elves []elf) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for _, e := range elves {
		minX = min(minX, e.x)
		minY = min(minY, e.y)
		maxX = max(maxX, e.x)
		maxY = max(maxY, e.y)
	}

	fmt.Printf("part one: %v\n", (maxX-minX+1)*(maxY-minY+1)-len(elves))
}

// https://freshman.tech/snippets/go/concatenate-slices/
func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
