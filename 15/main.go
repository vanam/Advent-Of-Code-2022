package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

type void struct{}

var member void

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var data [][4]int

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.FieldsFunc(line, Split)
		var c [4]int
		for i := 0; i < len(lineParts); i++ {
			parts := strings.Split(lineParts[i], "=")
			v, _ := strconv.Atoi(parts[len(parts)-1])
			c[i] = v
		}
		data = append(data, c)
	}

	// input-small.txt
	// fmt.Printf("part one: %v\n", partOne(data, 10))
	// fmt.Printf("part two: %v\n", partTwo(data, 20))

	// input.txt
	fmt.Printf("part one: %v\n", partOne(data, 2000000))
	fmt.Printf("part two: %v\n", partTwo(data, 4000000))
}

func partOne(data [][4]int, testY int) int {
	var segments [][2]int
	beaconsAtY := make(map[int]void)

	for _, d := range data {
		sx := d[0]
		sy := d[1]
		bx := d[2]
		by := d[3]

		dm := dist(sx, sy, bx, by)

		// no coverage by this sensor
		if testY < sy-dm || sy+dm < testY {
			continue
		}
		dx := dm - Abs(sy-testY)
		segments = append(segments, [2]int{sx - dx, sx + dx})

		// Keep track existing beacon at Y
		if by == testY {
			beaconsAtY[bx] = member
		}
	}
	sortSegments(segments)

	positionsWithoutBeacon := 0
	largestX := math.MinInt32
	for _, s := range segments {
		// Prevent from recording duplicate parts of segments
		if largestX >= s[0] {
			s[0] = largestX + 1
		}

		// Whole segment was already recorded
		if s[0] > s[1] {
			continue
		}

		// Record lenght of a segment
		positionsWithoutBeacon += s[1] - s[0] + 1
		largestX = s[1] // Move pointer to the end of last processed segment
	}

	return positionsWithoutBeacon - len(beaconsAtY)
}

func partTwo(data [][4]int, testBounds int) int {
	// Just try every line if there is gap for beacon
	for i := 0; i <= testBounds; i++ {
		bx := testLine(data, i, 0, testBounds)

		if bx != -1 {
			return bx*4000000 + i
		}
	}

	return -1 // Should not happen
}

func testLine(data [][4]int, testY, minC, maxC int) int {
	var segments [][2]int

	for _, d := range data {
		sx := d[0]
		sy := d[1]
		bx := d[2]
		by := d[3]

		dm := dist(sx, sy, bx, by)

		if testY < sy-dm || sy+dm < testY {
			continue
		}
		dx := dm - Abs(sy-testY)
		segments = append(segments, [2]int{sx - dx, sx + dx})
	}
	sortSegments(segments)

	largestX := minC
	for _, s := range segments {
		// Prevent from recording duplicate parts of segments
		if largestX >= s[0] {
			s[0] = largestX + 1
		}

		// Whole segment was already recorded
		if s[0] > s[1] {
			continue
		}

		// Look for a gap
		if largestX+1 != s[0] {
			return largestX + 1
		}

		largestX = s[1]
	}

	return -1 // no gap for beacon found
}

// Sort by X coordinate ascending, by Y coordinate descending
func sortSegments(segments [][2]int) {
	sort.Slice(segments, func(i, j int) bool {
		if segments[i][0] == segments[j][0] {
			return segments[i][1] > segments[j][1]
		}
		return segments[i][0] < segments[j][0]
	})
}

// Manhattan distance
func dist(ax, ay, bx, by int) int {
	return Abs(ax-bx) + Abs(ay-by)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Split(r rune) bool {
	return r == ':' || r == ','
}
