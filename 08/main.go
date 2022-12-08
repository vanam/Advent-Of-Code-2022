package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

var trees [100][100]int
var size = 0

var directions = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		size = len(line)
		for j, b := range line {
			val, _ := strconv.Atoi(string(b))
			trees[i][j] = val
		}
		i++
	}

	fmt.Printf("part one: %v\n", partOne())
	fmt.Printf("part two: %v\n", partTwo())
}

func partOne() int {
	count := 4*size - 4
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			for _, d := range directions {
				if isVisible(i, j, size, d) {
					count++
					break
				}
			}
		}
	}
	return count
}

func isVisible(i, j, size int, d [2]int) bool {
	height := trees[i][j]
	for i != 0 && j != 0 && i < size && j < size {
		i += d[0]
		j += d[1]
		if trees[i][j] >= height {
			return false
		}
	}
	return true
}

func partTwo() int {
	maxScore := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			score := getScore(i, j, size)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func getScore(i, j, size int) int {
	score := 1
	for _, d := range directions {
		score *= getViewDistance(i, j, size, d)
	}
	return score
}

func getViewDistance(i, j, size int, d [2]int) int {
	height := trees[i][j]
	viewDist := 0
	for i != 0 && j != 0 && i < size-1 && j < size-1 {
		i += d[0]
		j += d[1]
		viewDist++
		if trees[i][j] >= height {
			break
		}
	}
	return viewDist
}
