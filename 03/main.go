package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

// https://linuxhint.com/golang-set/
type void struct{}

var member void

//go:embed input.txt
var input []byte

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		set := make(map[byte]void)

		// add first half to a set
		for i := 0; i < len(line)/2; i++ {
			set[line[i]] = member
		}

		// test second half for a duplicate char
		for i := len(line) / 2; i < len(line); i++ {
			if _, ok := set[line[i]]; ok {
				score += priority(rune(line[i]))
				break
			}
		}
	}

	return score
}

func partTwo(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		counter := make(map[rune]int)

		line1 := scanner.Text()
		for _, char := range line1 {
			counter[char] |= 1
		}

		scanner.Scan()
		line2 := scanner.Text()
		for _, char := range line2 {
			counter[char] |= 2
		}

		scanner.Scan()
		line3 := scanner.Text()
		for _, char := range line3 {
			counter[char] |= 4
		}

		for k, v := range counter {
			if v == 7 {
				score += priority(k)
				break
			}
		}
	}

	return score
}

func priority(char rune) int {
	if 'a' <= char && char <= 'z' {
		return int(char) - 'a' + 1
	} else {
		return int(char) - 'A' + 27
	}
}
