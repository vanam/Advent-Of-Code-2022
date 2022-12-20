package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	list := make([][2]int, 0)
	list2 := make([][2]int, 0)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)
		list = append(list, [2]int{v, i})
		list2 = append(list2, [2]int{811589153 * v, i})
		i++
	}

	fmt.Printf("part one: %v\n", partOne(list))
	fmt.Printf("part two: %v\n", partTwo(list2))
}

func partOne(input [][2]int) int {
	queue := make([][2]int, len(input))
	list := make([][2]int, len(input))

	copy(queue, input)
	copy(list, input)

	for _, value := range queue {
		iv := find(value, list)
		list = concatMultipleSlices([][][2]int{list[0:iv], list[iv+1:]})
		newI := ((iv+value[0])%len(list) + len(list)) % len(list)
		list = concatMultipleSlices([][][2]int{list[0:newI], {value}, list[newI:]})
	}

	return grooveCoordinateSum(list)
}

func partTwo(input [][2]int) int {
	queue := make([][2]int, len(input))
	list := make([][2]int, len(input))

	copy(queue, input)
	copy(list, input)

	for i := 0; i < 10; i++ {
		for _, value := range queue {
			if value[0] == 0 {
				continue
			}
			iv := find(value, list)
			list = concatMultipleSlices([][][2]int{list[0:iv], list[iv+1:]})
			newI := ((iv+value[0])%len(list) + len(list)) % len(list)
			list = concatMultipleSlices([][][2]int{list[0:newI], {value}, list[newI:]})
		}
	}

	return grooveCoordinateSum(list)
}

func find(value [2]int, list [][2]int) int {
	for i, v := range list {
		if v[0] == value[0] && v[1] == value[1] {
			return i
		}
	}

	return -1
}

func grooveCoordinateSum(list [][2]int) int {
	var i0 int
	for i, v := range list {
		if v[0] == 0 {
			i0 = i
			break
		}
	}
	len := len(list)
	return list[(i0+1000)%len][0] + list[(i0+2000)%len][0] + list[(i0+3000)%len][0]
}

// https://freshman.tech/snippets/go/concatenate-slices/
func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
