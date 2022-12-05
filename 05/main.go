package main

import (
	"bufio"
	"bytes"
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) string {
	stacks := readStacks(scanner)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		move, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		from--
		to, _ := strconv.Atoi(parts[5])
		to--

		for i := 0; i < move; i++ {
			crate := stacks[from].Front()
			stacks[from].Remove(crate)
			stacks[to].PushFront(crate.Value)
		}
	}

	return getOutput(stacks)
}

func partTwo(scanner *bufio.Scanner) string {
	stacks := readStacks(scanner)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		move, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		from--
		to, _ := strconv.Atoi(parts[5])
		to--

		var crates []rune

		for i := 0; i < move; i++ {
			crate := stacks[from].Front()
			stacks[from].Remove(crate)
			crates = append(crates, crate.Value.(rune))
		}

		for i := move - 1; i >= 0; i-- {
			stacks[to].PushFront(crates[i])
		}
	}

	return getOutput(stacks)
}

func readStacks(scanner *bufio.Scanner) []*list.List {
	var stacks []*list.List

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		for i := 0; i < len(line); i += 4 {
			crate := rune(line[i+1])

			// make sure stack i is created
			if i/4+1 > len(stacks) {
				stacks = append(stacks, list.New())
			}

			// skip empty crates
			if crate == ' ' {
				continue
			}

			stacks[i/4].PushBack(crate)
		}
	}

	return stacks
}

func getOutput(stacks []*list.List) string {
	var result []rune
	for _, s := range stacks {
		result = append(result, s.Front().Value.(rune))
	}

	return string(result)
}
