package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

type expedition struct {
	x int
	y int
}

type blizzard struct {
	x         int
	y         int
	direction [2]int
}

// x, y
var up = [2]int{0, -1}
var down = [2]int{0, 1}
var left = [2]int{-1, 0}
var right = [2]int{1, 0}
var stay = [2]int{0, 0}

var directions = [5][2]int{
	up, down, left, right, stay,
}

type void struct{}

var member void

//go:embed input.txt
var input []byte

var sizeX int
var sizeY int
var blizzards []blizzard // global state of current blizards

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	blizzards = make([]blizzard, 0)

	sizeX = 0
	sizeY = 0
	for scanner.Scan() {
		line := scanner.Text()
		sizeX = len(line)
		for x := 0; x < len(line); x++ {
			if line[x] == '>' {
				blizzards = append(blizzards, blizzard{x, sizeY, right})
			} else if line[x] == '<' {
				blizzards = append(blizzards, blizzard{x, sizeY, left})
			} else if line[x] == '^' {
				blizzards = append(blizzards, blizzard{x, sizeY, up})
			} else if line[x] == 'v' {
				blizzards = append(blizzards, blizzard{x, sizeY, down})
			}
		}
		sizeY++
	}

	start := [2]int{1, 0}
	end := [2]int{sizeX - 2, sizeY - 1}

	firstPass := simulate(start, end)
	fmt.Printf("part one: %v\n", firstPass)
	fmt.Printf("part two: %v\n", firstPass+simulate(end, start)+simulate(start, end))
}

func simulate(start, end [2]int) int {
	expeditions := make(map[expedition]void)
	expeditions[expedition{start[0], start[1]}] = member

	for s := 1; s <= 300; s++ { // 300 to make sure we don't have infinite cycle
		// move blizzards to new positions
		newBlizzards := make([]blizzard, 0)
		newBlizzardBuckets := make(map[int][]blizzard)
		for _, b := range blizzards {
			newB := blizzard{b.x + b.direction[0], b.y + b.direction[1], b.direction}
			if newB.x == 0 {
				newB.x = sizeX - 2
			} else if newB.x == sizeX-1 {
				newB.x = 1
			} else if newB.y == 0 {
				newB.y = sizeY - 2
			} else if newB.y == sizeY-1 {
				newB.y = 1
			}
			newBlizzards = append(newBlizzards, newB)
			newBlizzardBuckets[newB.x] = append(newBlizzardBuckets[newB.x], newB)
		}

		newExpeditions := make(map[expedition]void)
		for e := range expeditions {
			for _, d := range directions {
				newE := expedition{e.x + d[0], e.y + d[1]}

				// we can stay at the start as long as we like
				if newE.x == start[0] && newE.y == start[1] {
					newExpeditions[newE] = member
					continue
				}
				// we reached the end
				if newE.x == end[0] && newE.y == end[1] {
					blizzards = newBlizzards
					return s
				}

				// out of bounds
				if newE.x <= 0 || newE.y <= 0 || newE.x >= sizeX-1 || newE.y >= sizeY-1 {
					continue
				}

				survivesBlizzard := true
				for _, b := range newBlizzardBuckets[newE.x] {
					if b.x == newE.x && b.y == newE.y {
						survivesBlizzard = false
						break
					}
				}

				// keep the survivors
				if survivesBlizzard {
					newExpeditions[newE] = member
				}
			}

		}

		expeditions = newExpeditions
		blizzards = newBlizzards
	}

	return 0 // should not happen
}

func printBlizzards(blizzards []blizzard, sizeX, sizeY int) {
	M := make([][]rune, sizeY)
	for y := 0; y < sizeY; y++ {
		M[y] = make([]rune, sizeX)
	}

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if x == 0 || x == sizeX-1 || y == 0 || y == sizeY-1 {
				M[y][x] = '#'
			} else {
				M[y][x] = '.'
			}
		}
	}

	M[0][1] = '.'
	M[sizeY-1][sizeX-2] = '.'

	for _, b := range blizzards {
		switch M[b.y][b.x] {
		case '.':
			switch b.direction {
			case up:
				M[b.y][b.x] = '^'
			case down:
				M[b.y][b.x] = 'v'
			case left:
				M[b.y][b.x] = '<'
			case right:
				M[b.y][b.x] = '>'
			}
		case '^', 'v', '<', '>':
			M[b.y][b.x] = '2'
		case '2':
			M[b.y][b.x] = '3'
		case '3':
			M[b.y][b.x] = '4'
		}
	}

	for y := 0; y < sizeY; y++ {
		fmt.Println(string(M[y][:]))
	}
}
