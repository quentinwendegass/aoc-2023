package main

import (
	"aoc2023/utils"
	"strings"
)

type BreakCondition func(string) bool

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	directions, pathMap, _ := transformInput(data)

	return findSteps("AAA", directions, pathMap, func(currentPath string) bool { return currentPath == "ZZZ" })
}

func partTwo(data []byte) int {
	directions, pathMap, startingPaths := transformInput(data)
	allSteps := make([]int, len(startingPaths))

	for i, path := range startingPaths {
		allSteps[i] = findSteps(path, directions, pathMap, func(currentPath string) bool { return currentPath[2] == 'Z' })
	}

	return utils.LCM(allSteps)
}

func findSteps(path string, directions string, pathMap map[string][]string, breakCondition BreakCondition) int {
	directionIdx := 0
	steps := 0
	currentPath := path

	for {
		direction := directions[directionIdx]
		nextPaths := pathMap[currentPath]

		if direction == 'L' {
			currentPath = nextPaths[0]
		} else {
			currentPath = nextPaths[1]
		}

		steps++

		if breakCondition(currentPath) {
			break
		}

		directionIdx++
		if directionIdx >= len(directions) {
			directionIdx = 0
		}
	}

	return steps
}

func transformInput(data []byte) (string, map[string][]string, []string) {
	lines := strings.Split(string(data), "\n")

	directions := lines[0]
	pathLines := lines[2:]

	startingPaths := make([]string, 0, 100)

	pathMap := make(map[string][]string, len(pathLines))

	for _, pathLine := range pathLines {
		currentPath := pathLine[:3]
		nextLeftPath := pathLine[7:10]
		nextRightPath := pathLine[12:15]

		if currentPath[2] == 'A' {
			startingPaths = append(startingPaths, currentPath)
		}

		pathMap[currentPath] = []string{nextLeftPath, nextRightPath}
	}

	return directions, pathMap, startingPaths
}
