package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

var pi = 0
var jetPattern string

const chamberHeight = 35000
const chamberWidth = 7

var chamber [chamberHeight][chamberWidth]rune

var ri = 0
var rocks [][][]rune = [][][]rune{
	{
		{'#', '#', '#', '#'},
	},
	{
		{'.', '#', '.'},
		{'#', '#', '#'},
		{'.', '#', '.'},
	},
	{
		{'.', '.', '#'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	},
	{
		{'#'},
		{'#'},
		{'#'},
		{'#'},
	},
	{
		{'#', '#'},
		{'#', '#'},
	},
}

var series = make([]int, 0)
var heights = make([]int, 0)

func main() {
	heights = append(heights, 0)

	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Scan()
	jetPattern = scanner.Text()

	for i := 0; i < chamberHeight; i++ {
		for j := 0; j < chamberWidth; j++ {
			chamber[i][j] = '.'
		}
	}

	fmt.Printf("part one: %v\n", simulate(2022))
	fmt.Printf("part two: %v\n", partTwo())
}

func partTwo() int {
	// run simulation for next 20k rocks
	simulate(20000)

	// ugly way how to find a repeating sequence
	// TODO how to do it better?
	sequenceStart := 0
	sequenceLength := -1
	for i := 0; i < len(series)/2; i++ {
		for length := 5; length < (len(series)-i)/2; length++ {
			if Equal(series[i:i+length], series[i+length:i+2*length]) {
				if length > sequenceLength {
					sequenceStart = i
					sequenceLength = length
				}
			}
		}
	}

	// calculate height after 1.000.000.000.000 rocks have stopped
	rockCount := 1000000000000

	// height at the beginning until the repeating sequence starts
	beginningHeight := heights[sequenceStart]
	rockCount -= sequenceStart

	// height of N full repeating sequences
	sequencesHeight := rockCount / sequenceLength * (heights[sequenceStart+sequenceLength] - heights[sequenceStart])
	rockCount %= sequenceLength

	// remaining height
	remainingHeight := (heights[sequenceStart+rockCount] - heights[sequenceStart])

	return beginningHeight + sequencesHeight + remainingHeight
}

func simulate(rockCount int) int {
	height := heights[len(heights)-1]

	// simulute all desired rock count
	for i := 0; i < rockCount; i++ {
		rock := getRock()
		// top left corner coordinates
		x := 2
		y := height + len(rock) + 3

		for {
			// move to the side if we can
			move := getMove()
			switch move {
			case '>':
				if !intersect(x+1, y, rock) {
					x++
				}
			case '<':
				if !intersect(x-1, y, rock) {
					x--
				}
			}

			// stop if cannot move down
			if intersect(x, y-1, rock) {
				stopRock(&chamber, x, y, rock)
				height = max(height, y)

				series = append(series, encodeRow(height))
				heights = append(heights, height)
				break
			}

			// move down
			y--
		}
	}

	return height
}

func stopRock(chamber *[chamberHeight][chamberWidth]rune, x, y int, rock [][]rune) {
	for i := 0; i < len(rock); i++ {
		for j := 0; j < len(rock[0]); j++ {
			if rock[i][j] == '.' {
				continue
			}
			chamber[y-i][x+j] = '#'
		}
	}
}

// true if intersects with chamber content/walls
func intersect(x, y int, rock [][]rune) bool {
	// rock colides with the chamber
	if y-len(rock)+1 <= 0 || x < 0 || x+len(rock[0]) > chamberWidth {
		return true
	}

	// check if rock colide with the chamber content
	for i := 0; i < len(rock); i++ {
		for j := 0; j < len(rock[0]); j++ {
			if rock[i][j] == '.' {
				continue
			}
			if chamber[y-i][x+j] == '#' {
				return true
			}
		}
	}

	return false
}

func getMove() rune {
	move := rune(jetPattern[pi])
	pi = (pi + 1) % len(jetPattern)
	return move
}

func getRock() [][]rune {
	rock := rocks[ri]
	ri = (ri + 1) % len(rocks)
	return rock
}

// binary encode chamber row content
func encodeRow(row int) int {
	a := 0
	for i := 0; i < chamberWidth; i++ {
		if chamber[row][i] == '#' {
			a = a | (1 << i)
		}
	}
	return a
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printChamberWithFallingRock(height int, x, y int, rock [][]rune) {
	var duplicate [chamberHeight][chamberWidth]rune

	for i := 0; i < chamberHeight; i++ {
		for j := 0; j < chamberWidth; j++ {
			duplicate[i][j] = chamber[i][j]
		}
	}

	stopRock(&duplicate, x, y, rock)

	printChamber(duplicate, height)
}

func printChamber(chamber [chamberHeight][chamberWidth]rune, height int) {
	for i := height + 7; i > 0; i-- {
		fmt.Printf("|%s|\n", string(chamber[i][:]))
	}
	fmt.Println("+-------+")
}
