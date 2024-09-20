package main

import (
	"aoc2023/utils"
	"strings"
)

type Position int

const (
	PositionFirst Position = iota
	PositionLast  Position = iota
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	sequences := transformInput(data)

	return findTotalValueForPositionInSequence(sequences, PositionLast)
}

func partTwo(data []byte) int {
	sequences := transformInput(data)

	return findTotalValueForPositionInSequence(sequences, PositionFirst)
}

func findTotalValueForPositionInSequence(sequences [][]int, position Position) int {
	total := 0

	for _, sequence := range sequences {
		differencesOverIter := make([]int, 0, 100)
		findDifference(sequence, &differencesOverIter, position)

		lastDifference := 0
		for i := len(differencesOverIter) - 2; i >= 0; i-- {
			if position == PositionFirst {
				lastDifference = differencesOverIter[i] - lastDifference
			} else {
				lastDifference += differencesOverIter[i]
			}
		}

		if position == PositionFirst {
			total += sequence[0] - lastDifference
		} else {
			total += lastDifference + sequence[len(sequence)-1]
		}
	}

	return total
}

func findDifference(sequence []int, differencesOverIter *[]int, position Position) {
	differences := make([]int, len(sequence)-1)

	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
	}

	var difference int
	if position == PositionFirst {
		difference = differences[0]
	} else {
		difference = differences[len(differences)-1]
	}

	*differencesOverIter = append(*differencesOverIter, difference)

	if utils.IsAllZero(differences) {
		return
	}

	findDifference(differences, differencesOverIter, position)
}

func transformInput(data []byte) [][]int {
	lines := strings.Split(string(data), "\n")

	sequences := make([][]int, len(lines))

	for i, line := range lines {
		sequences[i] = utils.AToISlice(utils.ValuesFromString(line, " "))
	}

	return sequences
}
