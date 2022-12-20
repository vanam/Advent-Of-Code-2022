package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

const size = 100

const empty = 0
const cube = 1
const visited = 2

var grid [size][size][size]int

var dirs = [6][3]int{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo())
}

func partOne(scanner *bufio.Scanner) int {
	surface := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ",")
		x, _ := strconv.Atoi(lineParts[0])
		y, _ := strconv.Atoi(lineParts[1])
		z, _ := strconv.Atoi(lineParts[2])

		// center cube in space
		x += size / 2
		y += size / 2
		z += size / 2

		grid[x][y][z] = cube
		surface += 6 // add all 6 sides

		// substact 2 sides for every neighbouring cube
		for _, d := range dirs {
			if grid[x+d[0]][y+d[1]][z+d[2]] == cube {
				surface -= 2
			}
		}
	}

	return surface
}

func partTwo() int {
	surface := 0
	// run 3D floodfill
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{0, 0, 0})

	for len(queue) > 0 {
		x := queue[0][0]
		y := queue[0][1]
		z := queue[0][2]
		queue = queue[1:]

		if grid[x][y][z] == visited {
			continue
		}

		grid[x][y][z] = visited

		for _, d := range dirs {
			nx := x + d[0]
			ny := y + d[1]
			nz := z + d[2]

			// don't flood outside of boundaries
			if nx < 0 || ny < 0 || nz < 0 || nx >= size || ny >= size || nz >= size {
				continue
			}

			// did we meet a cube?
			if grid[nx][ny][nz] == cube {
				surface++
			} else if grid[nx][ny][nz] == empty {
				// continue flooding
				queue = append(queue, [3]int{nx, ny, nz})
			}
		}
	}

	return surface
}
