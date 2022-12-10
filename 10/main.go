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

const strengthReadIncrement = 40
const crtWidth = 40
const crtHeight = 6

var instructionLenght = map[string]int{
	"noop": 1,
	"addx": 2,
}

func main() {
	// 0=value during the cycle, 1=value at the end of the cycle, 2=value in the next cycle
	registeryLength := 3
	R := make([]int, registeryLength)
	for i := 0; i < registeryLength; i++ {
		R[i] = 1
	}

	CRT := make([][]rune, crtHeight)
	for i := range CRT {
		CRT[i] = make([]rune, crtWidth)
	}
	for i := 0; i < crtHeight; i++ {
		for j := 0; j < crtWidth; j++ {
			CRT[i][j] = '.'
		}
	}

	strength := 0   // signal strength
	cStrength := 20 // cycle when we add signal strength

	c := 0
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		c++
		line := scanner.Text()
		lineParts := strings.Split(line, " ")

		if lineParts[0] == "addx" {
			val, _ := strconv.Atoi(lineParts[1])
			R[registeryLength-1] += val
		}

		for i := 0; i < instructionLenght[lineParts[0]]; i++ {
			if c == cStrength {
				strength += R[0] * c
				cStrength += strengthReadIncrement
			}

			p := c - 1
			dp := p%crtWidth - R[0]
			if dp == -1 || dp == 0 || dp == 1 {
				CRT[p/crtWidth][p%crtWidth] = '#'
			}

			for j := 1; j < registeryLength; j++ {
				R[j-1] = R[j]
			}
			c++
		}
		c--
	}

	fmt.Printf("part one: %v\n", strength)
	fmt.Println("part two:")
	for _, l := range CRT {
		for _, c := range l {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}
