package main

import (
	"bufio"
	"bytes"
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

const N = 3

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	mostCalories := 0
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mostCalories = Max(mostCalories, currentCalories)
			currentCalories = 0
		}
		calories, _ := strconv.Atoi(line)
		currentCalories += calories
	}

	mostCalories = Max(mostCalories, currentCalories)

	return mostCalories
}

func partTwo(scanner *bufio.Scanner) int {
	mostCalories := list.New()
	mostCalories.PushBack(0)

	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			keepNMax(mostCalories, currentCalories)
			currentCalories = 0
		}
		calories, _ := strconv.Atoi(line)
		currentCalories += calories
	}

	keepNMax(mostCalories, currentCalories)

	return Sum(mostCalories)
}

func keepNMax(l *list.List, y int) {
	for e := l.Front(); e != nil; e = e.Next() {
		if y > e.Value.(int) {
			l.InsertBefore(y, e)
			break
		}
	}

	if l.Len() > N {
		l.Remove(l.Back())
	}

}

func Sum(l *list.List) int {
	result := 0
	for e := l.Front(); e != nil; e = e.Next() {
		result += e.Value.(int)
	}
	return result
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
