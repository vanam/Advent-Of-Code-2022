package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

type monkey struct {
	items                []int
	operation            string
	op                   int
	divTest              int
	testTrueMonkeyIndex  int
	testFalseMonkeyIndex int
	inspectCount         int
}

func main() {
	fmt.Printf("part one: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), 20, 3))
	fmt.Printf("part two: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), 10000, 1))
}

func solve(scanner *bufio.Scanner, rounds int, divisor int) int {
	var monkeys []monkey
	modulo := 1

	for scanner.Scan() {
		m := readMonkey(scanner)
		monkeys = append(monkeys, m)
		modulo *= m.divTest
	}

	for r := 0; r < rounds; r++ {
		for mi := 0; mi < len(monkeys); mi++ {
			m := &monkeys[mi]
			for _, wl := range m.items {
				if m.operation == "**" {
					wl *= wl
				} else if m.operation == "+" {
					wl += m.op
				} else if m.operation == "*" {
					wl *= m.op
				}

				if divisor != 1 {
					wl /= 3 // round down
				} else {
					wl %= modulo
				}

				if wl%m.divTest == 0 {
					monkeys[m.testTrueMonkeyIndex].items = append(monkeys[m.testTrueMonkeyIndex].items, wl)
				} else {
					monkeys[m.testFalseMonkeyIndex].items = append(monkeys[m.testFalseMonkeyIndex].items, wl)
				}
				m.inspectCount++
			}

			m.items = make([]int, 0)
		}
	}

	inspectionCounts := make([]int, len(monkeys))
	for i, m := range monkeys {
		// fmt.Println("Monkey", i, "inspected items", m.inspectCount, "times.")
		inspectionCounts[i] = m.inspectCount
	}
	sort.Ints(inspectionCounts)

	return inspectionCounts[len(monkeys)-1] * inspectionCounts[len(monkeys)-2]
}

func readMonkey(scanner *bufio.Scanner) monkey {
	// Monkey 0:
	scanner.Text()
	// 	Starting items: 97, 81, 57, 57, 91, 61
	scanner.Scan()
	itemLine := scanner.Text()
	itemLineParts := strings.Split(itemLine, " ")[4:]
	items := make([]int, len(itemLineParts))
	for i := 0; i < len(items); i++ {
		itemLen := len(itemLineParts[i])
		if itemLineParts[i][itemLen-1] == ',' {
			itemLen--
		}
		item, _ := strconv.Atoi(itemLineParts[i][:itemLen])
		items[i] = item
	}

	// 	Operation: new = old * 7
	scanner.Scan()
	opLine := scanner.Text()
	opLineParts := strings.Split(opLine, " ")
	op := opLineParts[6]
	opVal, err := strconv.Atoi(opLineParts[7])
	if err != nil {
		op = "**"
	}

	// 	Test: divisible by 11
	scanner.Scan()
	testLine := scanner.Text()
	testLineParts := strings.Split(testLine, " ")
	modulo, _ := strconv.Atoi(testLineParts[5])
	// 		If true: throw to monkey 5
	scanner.Scan()
	trueTestLine := scanner.Text()
	trueTestLineParts := strings.Split(trueTestLine, " ")
	trueTestLineMonkeyIndex, _ := strconv.Atoi(trueTestLineParts[9])
	// 		If false: throw to monkey 6
	scanner.Scan()
	falseTestLine := scanner.Text()
	falseTestLineParts := strings.Split(falseTestLine, " ")
	falseTestLineMonkeyIndex, _ := strconv.Atoi(falseTestLineParts[9])
	// \n
	scanner.Scan()

	return monkey{
		items:                items,
		operation:            op,
		op:                   opVal,
		divTest:              modulo,
		testTrueMonkeyIndex:  trueTestLineMonkeyIndex,
		testFalseMonkeyIndex: falseTestLineMonkeyIndex,
		inspectCount:         0,
	}
}
