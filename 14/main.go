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

const sizeX = 1000
const sizeY = 1000

var minX = sizeX
var minY = sizeY
var maxX = 0
var maxY = 0

const nextSandX = 500
const nextSandY = 0

var sandMoves = [3][2]int{{0, 1}, {-1, 1}, {1, 1}}

var grid [sizeX][sizeY]rune

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	readGrid(scanner)

	sandUnits := 1
	for ; ; sandUnits++ {
		if !nextSand() {
			break
		}
	}
	// printGrid()

	return sandUnits - 1 // last sand falls into the abyss
}

func nextSand() bool {
	x := nextSandX
	y := nextSandY

	for {
		var nx int
		var ny int
		hasValidMove := false
		for _, move := range sandMoves {
			nx = x + move[0]
			ny = y + move[1]

			if grid[ny][nx] == '.' {
				hasValidMove = true
				break
			}
		}

		if !hasValidMove {
			grid[y][x] = 'o'
			return true
		}

		// Is the sand falling to the abyss?
		if nx < minX || nx > maxX || ny > maxY {
			return false
		}

		x = nx
		y = ny
	}
}

func partTwo(scanner *bufio.Scanner) int {
	readGrid(scanner)

	sandUnits := 1
	for ; ; sandUnits++ {
		if !nextSand2() {
			break
		}
	}
	// printGrid()

	return sandUnits
}

func nextSand2() bool {
	x := nextSandX
	y := nextSandY

	rockBotttomY := 2 + maxY

	for {
		var nx int
		var ny int
		hasValidMove := false
		for _, move := range sandMoves {
			nx = x + move[0]
			ny = y + move[1]

			// Check if we hit the rock bottom
			if grid[ny][nx] == '.' && ny < rockBotttomY {
				hasValidMove = true
				break
			}
		}

		if !hasValidMove {
			grid[y][x] = 'o'

			// Check if there is no more room for sand
			if x == nextSandX && y == nextSandY {
				return false
			}
			return true
		}

		x = nx
		y = ny
	}
}

func readGrid(scanner *bufio.Scanner) {
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			grid[y][x] = '.'
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " -> ")

		for i := 1; i < len(lineParts); i++ {
			fromParts := strings.Split(lineParts[i-1], ",")
			fx, _ := strconv.Atoi(fromParts[0])
			fy, _ := strconv.Atoi(fromParts[1])

			toParts := strings.Split(lineParts[i], ",")
			tx, _ := strconv.Atoi(toParts[0])
			ty, _ := strconv.Atoi(toParts[1])

			drawLine(fx, fy, tx, ty)

			minX = min(minX, fx)
			minX = min(minX, tx)
			maxX = max(maxX, fx)
			maxX = max(maxX, tx)

			minY = min(minY, fy)
			minY = min(minY, ty)
			maxY = max(maxY, fy)
			maxY = max(maxY, ty)
		}
	}
}

func drawLine(fx, fy, tx, ty int) {
	for x := min(fx, tx); x <= max(fx, tx); x++ {
		for y := min(fy, ty); y <= max(fy, ty); y++ {
			grid[y][x] = '#'
		}
	}
}

func printGrid() {
	for y := max(0, minY-50); y < min(maxY+5, sizeY); y++ {
		for x := max(0, minX-50); x < min(maxX+50, sizeX); x++ {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
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
