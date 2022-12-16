package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var flowRates = make(map[int]int)
var tunnels = make(map[string][]string)
var valve2i = make(map[string]int)
var nonZeroFlowValves []int
var memo = make(map[state]int)

type state struct {
	valve1     int
	valve2     int
	minute1    int
	minute2    int
	openValves int
}

var adjMat [][]int

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	r, _ := regexp.Compile(`^Valve (?P<valve>[A-Z]+) has flow rate=(?P<flow>\d+); tunnel[s]? lead[s]? to valve[s]? (?P<tunnels>[A-Z]+(, ([A-Z]+))*)`)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = m[i]
			}
		}

		v := result["valve"]
		valve2i[v] = i
		flowRate, _ := strconv.Atoi(result["flow"])
		flowRates[i] = flowRate
		connectedValves := strings.Split(result["tunnels"], ", ")
		tunnels[v] = connectedValves
		if flowRate > 0 {
			nonZeroFlowValves = append(nonZeroFlowValves, i)
		}
		i++
	}
	// compute distances between valves
	floydWarshall()

	fmt.Printf("part one: %v\n", solve(valve2i["AA"], valve2i["AA"], 30, 1, 0))
	fmt.Printf("part two: %v\n", solve(valve2i["AA"], valve2i["AA"], 5, 5, 0))

}

func solve(v1, v2 int, m1 int, m2 int, openInt int) int {
	// time is up
	if m1 > 30 || m2 > 30 {
		return 0
	}

	key1 := state{valve1: v1, valve2: v2, minute1: m1, minute2: m2, openValves: openInt}
	key2 := state{valve1: v2, valve2: v1, minute1: m2, minute2: m1, openValves: openInt}

	result, alreadySeen := memo[key1] // have we already seen this state?
	if alreadySeen {
		return result
	}

	maxFlow := 0
	for _, vn1 := range nonZeroFlowValves {
		if !isOn(openInt, vn1) {
			d1 := adjMat[v1][vn1]
			newFlowVn1 := gain(vn1, d1, m1)

			// Stay at valve v1 in case we cannot move to the next one
			if newFlowVn1 <= 0 {
				vn1 = v1
				d1 = -1
				newFlowVn1 = 0
			}
			openInt = turnOnNthBit(openInt, vn1)

			for _, vn2 := range nonZeroFlowValves {
				if !isOn(openInt, vn2) {
					d2 := adjMat[v2][vn2]
					newFlowVn2 := gain(vn2, d2, m2)

					// Stay at valve v2 in case we cannot move to the next one
					if newFlowVn2 <= 0 {
						vn2 = v2
						d2 = -1
						newFlowVn2 = 0
					}
					openInt = turnOnNthBit(openInt, vn2)

					// Did we move at least once?
					if newFlowVn1+newFlowVn2 > 0 {
						maxFlow = max(maxFlow, newFlowVn1+newFlowVn2+solve(vn1, vn2, m1+d1+1, m2+d2+1, openInt))
					}

					// Turn off bit only if we moved
					if newFlowVn2 > 0 {
						openInt = turnOffNthBit(openInt, vn2)
					}
				}
			}

			// Turn off bit only if we moved
			if newFlowVn1 > 0 {
				openInt = turnOffNthBit(openInt, vn1)
			}
		}
	}

	// problem is symetrical
	memo[key1] = maxFlow
	memo[key2] = maxFlow

	return maxFlow
}

func floydWarshall() {
	// FloydWarshall
	valveCount := len(tunnels)
	adjMat = make([][]int, valveCount)
	for i := 0; i < valveCount; i++ {
		adjMat[i] = make([]int, valveCount)
	}

	for i := 0; i < valveCount; i++ {
		for j := 0; j < valveCount; j++ {
			if i == j {
				adjMat[i][j] = 0
			} else {
				adjMat[i][j] = math.MaxInt32
			}
		}
	}

	for u, uTunnels := range tunnels {
		for _, v := range uTunnels {
			ui := valve2i[u]
			vi := valve2i[v]
			adjMat[ui][vi] = 1
			adjMat[vi][ui] = 1
		}
	}

	for k := 0; k < valveCount; k++ {
		for i := 0; i < valveCount; i++ {
			for j := 0; j < valveCount; j++ {
				adjMat[i][j] = min(adjMat[i][j], adjMat[i][k]+adjMat[k][j])

			}
		}
	}
}

func gain(v int, d, m int) int {
	return flowRates[v] * (30 - m - d)
}

func isOn(a, n int) bool {
	return a&(1<<n) != 0
}

func turnOnNthBit(a, n int) int {
	return a | (1 << n)
}

func turnOffNthBit(a, n int) int {
	return a & (^(1 << n))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
