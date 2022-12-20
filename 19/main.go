package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

//go:embed input.txt
var input []byte

type state struct {
	ore                int
	clay               int
	obsidian           int
	geode              int
	oreRobotCount      int
	clayRobotCount     int
	obsidianRobotCount int
	geodeRobotCount    int
	oreInc             int
	clayInc            int
	obsidianInc        int
	geodeInc           int
}

const maxStates = 3000

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	r, _ := regexp.Compile(`^Blueprint \d+: Each ore robot costs (?P<ore_ore>[0-9]+) ore. Each clay robot costs (?P<clay_ore>\d+) ore. Each obsidian robot costs (?P<obsidian_ore>\d+) ore and (?P<obsidian_clay>\d+) clay. Each geode robot costs (?P<geode_ore>\d+) ore and (?P<geode_obsidian>\d+) obsidian.`)

	blueprints := make([]map[string]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)

		blueprint := make(map[string]int)
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				vi, _ := strconv.Atoi(m[i])
				blueprint[name] = vi
			}
		}
		blueprints = append(blueprints, blueprint)
	}

	fmt.Printf("part one: %v\n", partOne(blueprints, 24))
	fmt.Printf("part two: %v\n", partTwo(blueprints, 32))
}

func partOne(blueprints []map[string]int, minutes int) int {
	qualityLevelSum := 0
	for i := 0; i < len(blueprints); i++ {
		qualityLevelSum += (i + 1) * runSimulation(blueprints[i], minutes)
	}
	return qualityLevelSum
}

func partTwo(blueprints []map[string]int, minutes int) int {
	result := 1
	for i := 0; i < len(blueprints); i++ {
		if i >= 3 {
			break
		}
		result *= runSimulation(blueprints[i], minutes)
	}
	return result
}

func runSimulation(params map[string]int, minutes int) int {
	states := make([]state, 0)
	states = append(states, state{oreRobotCount: 1})

	for i := 1; i <= minutes; i++ {
		// calculate resource increments
		for i := 0; i < len(states); i++ {
			states[i].oreInc = states[i].oreRobotCount
			states[i].clayInc = states[i].clayRobotCount
			states[i].obsidianInc = states[i].obsidianRobotCount
			states[i].geodeInc = states[i].geodeRobotCount
		}

		// try robot constructing
		unprocessedStates := make([]state, len(states))
		copy(unprocessedStates, states)

		for len(unprocessedStates) > 0 {
			s := unprocessedStates[0]
			unprocessedStates = unprocessedStates[1:]

			if canConstructGeodeRobot(params, s) {
				newS := s
				newS.ore -= params["geode_ore"]
				newS.obsidian -= params["geode_obsidian"]
				newS.geodeRobotCount++
				states = append(states, newS)
			}
			if canConstructObsidianRobot(params, s) {
				newS := s
				newS.ore -= params["obsidian_ore"]
				newS.clay -= params["obsidian_clay"]
				newS.obsidianRobotCount++
				states = append(states, newS)
			}
			if canConstructClayRobot(params, s) {
				newS := s
				newS.ore -= params["clay_ore"]
				newS.clayRobotCount++
				states = append(states, newS)
			}
			if canConstructOreRobot(params, s) {
				newS := s
				newS.ore -= params["ore_ore"]
				newS.oreRobotCount++
				states = append(states, newS)
			}
		}

		// add resource increments
		for i := 0; i < len(states); i++ {
			states[i].ore += states[i].oreInc
			states[i].clay += states[i].clayInc
			states[i].obsidian += states[i].obsidianInc
			states[i].geode += states[i].geodeInc
		}

		// sort and prune states
		sortStates(states)

		newStates := make([]state, 0)
		newStates = append(newStates, states[0])
		for _, s := range states {
			if len(newStates) >= maxStates {
				break
			}
			// ignore equal states
			if s == newStates[len(newStates)-1] {
				continue
			}
			newStates = append(newStates, s)
		}
		states = newStates
	}

	return states[0].geode
}

func sortStates(states []state) {
	sort.Slice(states, func(i, j int) bool {
		s1 := states[i]
		s2 := states[j]

		if s1.geode > s2.geode {
			return true
		} else if s1.geode < s2.geode {
			return false
		}

		if s1.geodeRobotCount > s2.geodeRobotCount {
			return true
		} else if s1.geodeRobotCount < s2.geodeRobotCount {
			return false
		}

		if s1.obsidian > s2.obsidian {
			return true
		} else if s1.obsidian < s2.obsidian {
			return false
		}

		if s1.obsidianRobotCount > s2.obsidianRobotCount {
			return true
		} else if s1.obsidianRobotCount < s2.obsidianRobotCount {
			return false
		}

		if s1.clay > s2.clay {
			return true
		} else if s1.clay < s2.clay {
			return false
		}

		if s1.clayRobotCount > s2.clayRobotCount {
			return true
		} else if s1.clayRobotCount < s2.clayRobotCount {
			return false
		}

		if s1.ore > s2.ore {
			return true
		} else if s1.ore < s2.ore {
			return false
		}

		if s1.oreRobotCount > s2.oreRobotCount {
			return true
		} else if s1.oreRobotCount < s2.oreRobotCount {
			return false
		}

		return false
	})
}

func canConstructGeodeRobot(params map[string]int, s state) bool {
	return s.ore >= params["geode_ore"] && s.obsidian >= params["geode_obsidian"]
}

func canConstructObsidianRobot(params map[string]int, s state) bool {
	return s.ore >= params["obsidian_ore"] && s.clay >= params["obsidian_clay"]
}

func canConstructClayRobot(params map[string]int, s state) bool {
	return s.ore >= params["clay_ore"]
}

func canConstructOreRobot(params map[string]int, s state) bool {
	return s.ore >= params["ore_ore"]
}
