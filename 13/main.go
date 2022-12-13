package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)

//go:embed input.txt
var input []byte

type signal struct {
	value    int
	hasValue bool
	items    []*signal
}

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	sumOfOkPairIndices := 0
	i := 1
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan() // \n

		if compare(parseSignal(line1), parseSignal(line2)) < 0 {
			sumOfOkPairIndices += i
		}
		i++
	}

	return sumOfOkPairIndices
}

func partTwo(scanner *bufio.Scanner) int {
	signals := make([]signal, 2)
	dp1 := signal{items: []*signal{{items: []*signal{{value: 2, hasValue: true}}}}}
	signals[0] = dp1
	dp2 := signal{items: []*signal{{items: []*signal{{value: 6, hasValue: true}}}}}
	signals[1] = dp2

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan() // \n

		signals = append(signals, parseSignal(line1))
		signals = append(signals, parseSignal(line2))
	}

	sort.Slice(signals, func(i, j int) bool {
		return compare(signals[i], signals[j]) < 0
	})

	var idp1, idp2 int
	for i, v := range signals {
		if compare(v, dp1) == 0 {
			idp1 = i + 1
		} else if compare(v, dp2) == 0 {
			idp2 = i + 1
		}
		// fmt.Println(toString(v))
	}

	return idp1 * idp2
}

// <0:mensi, 0:stejny, >0:vetsi
func compare(left, right signal) int {
	if left.hasValue && right.hasValue {
		return left.value - right.value
	} else if !left.hasValue && !right.hasValue {
		for i := 0; i < len(left.items); i++ {
			// If the right list runs out of items first, the inputs are not in the right order
			if i >= len(right.items) {
				return 1
			}
			cmp := compare(*left.items[i], *right.items[i])
			if cmp != 0 {
				return cmp
			}
			// Continue checking the next part of the input
		}

		if len(left.items) < len(right.items) {
			// If the left list runs out of items first, the inputs are in the right order
			return -1
		} else {
			return 0
		}
	} else {
		// convert int to list when necessary
		if left.hasValue {
			left = signal{items: []*signal{{value: left.value, hasValue: true}}}
		}

		if right.hasValue {
			right = signal{items: []*signal{{value: right.value, hasValue: true}}}
		}

		return compare(left, right)
	}
}

func parseSignal(line string) signal {
	var currentSignal *signal
	stack := make([]*signal, 0)

	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			s := signal{items: make([]*signal, 0)}

			if currentSignal != nil {
				currentSignal.items = append(currentSignal.items, &s)
				currentSignal = &s
			} else {
				currentSignal = &s
			}

			stack = append(stack, &s)
		} else if line[i] == ']' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				currentSignal = stack[len(stack)-1]
			}
		} else if line[i] == ',' {
			continue
		} else {
			val := 0
			for ; i <
				len(line); i++ {

				if line[i] == ',' {
					break
				} else if line[i] == '[' || line[i] == ']' {
					i--
					break
				}

				v, _ := strconv.Atoi(string(line[i]))
				val *= 10
				val += v
			}

			s := signal{value: val, hasValue: true}
			currentSignal.items = append(currentSignal.items, &s)
		}
	}

	return *currentSignal
}

func toString(s signal) string {
	if s.hasValue {
		return strconv.FormatInt(int64(s.value), 10)
	}

	var result string
	for i := 0; i < len(s.items); i++ {
		if i > 0 {
			result += ","
		}
		result += toString(*s.items[i])
	}
	return "[" + result + "]"
}
