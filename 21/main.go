package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) string {
	exp := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		exp[lineParts[0]] = strings.Split(lineParts[1], " ")
	}

	for len(exp["root"]) != 1 {
		for k, v := range exp {
			if len(v) == 1 {
				continue
			}

			a, errA := strconv.Atoi(v[0])
			b, errB := strconv.Atoi(v[2])

			if errA != nil && len(exp[v[0]]) == 1 {
				v[0] = exp[v[0]][0]
			}

			if errB != nil && len(exp[v[2]]) == 1 {
				v[2] = exp[v[2]][0]
			}

			if errA != nil || errB != nil {
				exp[k] = v
				continue
			}

			var c int
			switch v[1] {
			case "+":
				c = a + b
			case "-":
				c = a - b
			case "*":
				c = a * b
			case "/":
				c = a / b
			}

			exp[k] = []string{strconv.Itoa(c)}
		}
	}

	return exp["root"][0]
}

func partTwo(scanner *bufio.Scanner) string {
	exp := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		exp[lineParts[0]] = strings.Split(lineParts[1], " ")
	}
	exp["humn"] = []string{"-", "-"}
	vRoot := exp["root"]
	vRoot[1] = "="
	exp["root"] = vRoot

	change := true
	for change {
		change = false
		for k, v := range exp {
			if len(v) == 1 || k == "humn" {
				continue
			}

			a, errA := strconv.Atoi(v[0])
			b, errB := strconv.Atoi(v[2])

			if errA != nil && len(exp[v[0]]) == 1 {
				v[0] = exp[v[0]][0]
				change = true
			}

			if errB != nil && len(exp[v[2]]) == 1 {
				v[2] = exp[v[2]][0]
				change = true
			}

			if errA != nil || errB != nil {
				exp[k] = v
				continue
			}

			var c int
			switch v[1] {
			case "+":
				c = a + b
			case "-":
				c = a - b
			case "*":
				c = a * b
			case "/":
				c = a / b
			}

			exp[k] = []string{strconv.Itoa(c)}
			change = true
		}
	}

	// Just print equation and use external solver
	// fmt.Println(toString(exp, exp["root"][0]) + "=" + exp["root"][2])
	// fmt.Println(strconv.Itoa(801541236563053696 / 232925)) // https://www.mathpapa.com/equation-solver/

	rightSide, _ := decimal.NewFromString(exp["root"][2])
	return solve(exp, exp["root"][0], rightSide).String()
}

func solve(exp map[string][]string, key string, value decimal.Decimal) decimal.Decimal {
	if key == "humn" {
		return value
	}

	v := exp[key]

	a, errA := decimal.NewFromString(v[0])
	b, errB := decimal.NewFromString(v[2])

	if errA == nil {
		switch v[1] {
		case "+":
			value = value.Sub(a)
		case "-":
			value = value.Sub(a).Mul(decimal.NewFromInt(-1))
		case "*":
			value = value.Div(a)
		case "/":
			value = value.Mul(a)
		}

	} else if errB == nil {
		switch v[1] {
		case "+":
			value = value.Sub(b)
		case "-":
			value = value.Add(b)
		case "*":
			value = value.Div(b)
		case "/":
			value = value.Mul(b)
		}
	}

	if errA != nil {
		return solve(exp, v[0], value)
	}

	return solve(exp, v[2], value)
}

func toString(exp map[string][]string, key string) string {
	if key == "humn" {
		return "x"
	}

	_, err := strconv.Atoi(key)

	if err == nil {
		return key
	}

	v := exp[key]

	return "(" + toString(exp, v[0]) + v[1] + toString(exp, v[2]) + ")"
}
