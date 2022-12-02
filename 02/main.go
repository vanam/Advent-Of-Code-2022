package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input-small.txt
var input []byte

const rock = 1
const paper = 2
const scissors = 3

const lose = 0
const draw = 3
const win = 6

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	choice2score := map[string]int{
		"A X": rock + draw,
		"B X": rock + lose,
		"C X": rock + win,

		"A Y": paper + win,
		"B Y": paper + draw,
		"C Y": paper + lose,

		"A Z": scissors + lose,
		"B Z": scissors + win,
		"C Z": scissors + draw,
	}

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += choice2score[line]
	}

	return score
}

func partTwo(scanner *bufio.Scanner) int {
	choice2score := map[string]int{
		"A X": scissors + lose,
		"B X": rock + lose,
		"C X": paper + lose,

		"A Y": rock + draw,
		"B Y": paper + draw,
		"C Y": scissors + draw,

		"A Z": paper + win,
		"B Z": scissors + win,
		"C Z": rock + win,
	}

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += choice2score[line]
	}

	return score
}
