package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("Part 1")
	solve(bufio.NewScanner(bytes.NewReader(input)), 4)
	fmt.Println("Part 2")
	solve(bufio.NewScanner(bytes.NewReader(input)), 14)
}

func solve(scanner *bufio.Scanner, n int) int {
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		counter := make(map[byte]int)

		for i := 0; i < len(line); i++ {
			if i < n-1 {
				counter[line[i]]++
				continue
			}
			counter[line[i]]++

			if isMarker(counter) {
				fmt.Printf("%5d\n", i+1)
				break
			}
			counter[line[i-n+1]]--
		}
	}

	return result
}

func isMarker(m map[byte]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}

	return true
}
