package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

var snafuDigits2Digits = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var digits2snafuDigits = map[int]rune{
	2: '2',
	1: '1',
	0: '0',
	3: '=',
	4: '-',
}

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += snafu2decimal(line)
	}

	return decimal2snafu(sum)
}

func snafu2decimal(snafuNumber string) int {
	decimalValue := 0
	n := 0
	for i := len(snafuNumber) - 1; i >= 0; i-- {
		decimalValue += snafuDigits2Digits[rune(snafuNumber[i])] * PowInts(5, n)
		n++
	}

	return decimalValue
}

func decimal2snafu(decimalNumber int) string {
	snafuValue := ""

	for decimalNumber > 0 {
		remainder := decimalNumber % 5
		snafuValue = string(digits2snafuDigits[remainder]) + snafuValue
		if remainder >= 3 {
			decimalNumber += 5
		}
		decimalNumber /= 5
	}

	return snafuValue
}

// Assumption: n >= 0
func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func test() {
	numbers := []int{
		1,         //              1
		2,         //              2
		3,         //             1=
		4,         //             1-
		5,         //             10
		6,         //             11
		7,         //             12
		8,         //             2=
		9,         //             2-
		10,        //             20
		15,        //            1=0
		20,        //            1-0
		2022,      //         1=11-2
		12345,     //        1-0---0
		314159265, //  1121-1110-1=0
	}
	for _, n := range numbers {
		fmt.Println(n, decimal2snafu(n))
	}

	snafuNumbers := []string{
		"1=-0-2", //     1747
		"12111",  //      906
		"2=0=",   //      198
		"21",     //       11
		"2=01",   //      201
		"111",    //       31
		"20012",  //     1257
		"112",    //       32
		"1=-1=",  //      353
		"1-12",   //      107
		"12",     //        7
		"1=",     //        3
		"122",    //       37
	}
	for _, n := range snafuNumbers {
		fmt.Println(n, snafu2decimal(n), decimal2snafu(snafu2decimal(n)))
	}
}
