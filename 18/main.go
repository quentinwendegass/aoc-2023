package main

import (
	"aoc2023/utils"
	"strconv"
	"strings"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

type DigStep struct {
	direction byte
	amount    int
}

func partOne(data []byte) int {
	digPlan := convertToDigPlan(data)
	return findHoleArea(digPlan)
}

func partTwo(data []byte) int {
	digPlan := convertHexToDigPlan(data)
	return findHoleArea(digPlan)
}

func findHoleArea(digPlan []DigStep) int {
	x := 0
	y := 0
	vertices := make([][]int, len(digPlan))
	for i, digStep := range digPlan {
		switch digStep.direction {
		case 'U':
			y += digStep.amount
		case 'D':
			y -= digStep.amount
		case 'L':
			x -= digStep.amount
		case 'R':
			x += digStep.amount
		}

		vertices[i] = []int{x, y}
	}

	boundaryPoints := utils.CountBoundaryPoints(vertices)
	return utils.CountPointsInsideArea(utils.ShoelaceArea(vertices), boundaryPoints) + boundaryPoints
}

func convertToDigPlan(data []byte) []DigStep {
	lines := strings.Split(string(data), "\n")

	digPlan := make([]DigStep, len(lines))

	for i, line := range lines {
		stepEntries := strings.Split(line, " ")
		amount, err := strconv.Atoi(stepEntries[1])
		if err != nil {
			panic(err)
		}

		digPlan[i] = DigStep{direction: stepEntries[0][0], amount: amount}
	}

	return digPlan
}

func convertHexToDigPlan(data []byte) []DigStep {
	lines := strings.Split(string(data), "\n")

	digPlan := make([]DigStep, len(lines))

	for i, line := range lines {
		stepEntries := strings.Split(line, " ")
		color := stepEntries[2][2 : len(stepEntries[2])-1]

		var direction byte
		switch color[5] {
		case '0':
			direction = 'R'
		case '1':
			direction = 'D'
		case '2':
			direction = 'L'
		case '3':
			direction = 'U'
		}

		amount, err := strconv.ParseInt(color[0:5], 16, 64)
		if err != nil {
			panic(err)
		}

		digPlan[i] = DigStep{direction: direction, amount: int(amount)}
	}

	return digPlan
}
