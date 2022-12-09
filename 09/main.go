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

type knot struct {
	x int
	y int
}

type void struct{}

var member void

var dir = map[string][2]int{
	"R": {1, 0},
	"L": {-1, 0},
	"U": {0, 1},
	"D": {0, -1},
}

func main() {
	fmt.Printf("part one: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), 2))
	fmt.Printf("part two: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), 10))
}

func solve(scanner *bufio.Scanner, knotCount int) int {
	set := make(map[knot]void)

	knots := make([]knot, knotCount)
	for i := 0; i < knotCount; i++ {
		knots[i] = knot{0, 0}
	}
	set[knots[knotCount-1]] = member

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		d := dir[lineParts[0]]
		steps, _ := strconv.Atoi(lineParts[1])

		for s := 0; s < steps; s++ {
			knots[0].x += d[0]
			knots[0].y += d[1]

			for i := 1; i < knotCount; i++ {
				head := knots[i-1]
				tail := &knots[i]

				dx := head.x - tail.x
				dy := head.y - tail.y

				if dx > 1 || dx < -1 || dy > 1 || dy < -1 {
					tail.x += sign(dx)
					tail.y += sign(dy)
				} else {
					break // knot doesn't move -> following knots don't move
				}

				if i == knotCount-1 {
					set[*tail] = member // keep track of tail
				}
			}
		}
	}

	return len(set)
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
