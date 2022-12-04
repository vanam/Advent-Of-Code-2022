package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Assignment struct {
	from int
	to   int
}

//go:embed input.txt
var input []byte

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")

		a1 := makeAssignment(assignments[0])
		a2 := makeAssignment(assignments[1])

		if assignmentsContain(a1, a2) {
			count++
		}
	}

	return count
}

func partTwo(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")

		a1 := makeAssignment(assignments[0])
		a2 := makeAssignment(assignments[1])

		if assignmentsOverlap(a1, a2) {
			count++
		}
	}

	return count
}

func assignmentsContain(a1 Assignment, a2 Assignment) bool {
	minFrom := min(a1.from, a2.from)
	maxTo := max(a1.to, a2.to)

	if a1.from == minFrom && a1.to == maxTo {
		return true
	}

	if a2.from == minFrom && a2.to == maxTo {
		return true
	}

	return false
}

func assignmentsOverlap(a1 Assignment, a2 Assignment) bool {
	maxFrom := max(a1.from, a2.from)
	minTo := min(a1.to, a2.to)

	return maxFrom <= minTo
}

func makeAssignment(assignmentStr string) Assignment {
	parts := strings.Split(assignmentStr, "-")

	var partInts []int
	for _, idStr := range parts {
		id, _ := strconv.Atoi(idStr)
		partInts = append(partInts, id)
	}

	return Assignment{partInts[0], partInts[1]}
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
