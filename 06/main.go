package main

import (
	"aoc2023/utils"
	"math"
	"strings"
)

type CalculateBetterDistanceCount func(int, int) int

func main() {
	var partOne = func(data []byte) int {
		return partOne(data, calculateBetterDistanceCountMath)
	}
	var partTwo = func(data []byte) int {
		return partTwo(data, calculateBetterDistanceCountMath)
	}
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte, calculate CalculateBetterDistanceCount) int {
	times, distances := transformInput(data)

	total := 1

	for i, time := range times {
		total *= calculate(time, distances[i])
	}

	return total
}

func partTwo(data []byte, calculate CalculateBetterDistanceCount) int {
	times, distances := transformInput(data)

	time := utils.CombineIntsToString(times)
	distance := utils.CombineIntsToString(distances)

	return calculate(time, distance)
}

func calculateBetterDistanceCountForce(time int, recordDistance int) int {
	betterDistancesCount := 0

	for holdTime := 0; holdTime < time+1; holdTime++ {
		distance := holdTime * (time - holdTime)

		if distance > recordDistance {
			betterDistancesCount++
		}
	}

	return betterDistancesCount
}

func calculateBetterDistanceCountMath(time int, recordDistance int) int {
	holdTime := math.Ceil((math.Sqrt(float64(time*time-4*recordDistance)) + float64(time)) / 2)

	betterDistancesCount := time - 2*(time-int(holdTime)) - 1

	return betterDistancesCount
}

func transformInput(data []byte) ([]int, []int) {
	lines := strings.Split(string(data), "\n")

	_, timeValues := utils.ValuesFromLabelString(lines[0], ":", " ")
	_, distanceValues := utils.ValuesFromLabelString(lines[1], ":", " ")

	return utils.AToISlice(timeValues), utils.AToISlice(distanceValues)
}
