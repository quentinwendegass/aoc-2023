package main

import (
	"aoc2023/utils"
	"bytes"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	return getTotalReflected(data, 0)
}

func partTwo(data []byte) int {
	return getTotalReflected(data, 1)
}

func getTotalReflected(data []byte, mandatoryDifferences int) int {
	entries := bytes.Split(data, []byte("\n\n"))

	totalReflectedRows := 0
	totalReflectedColumns := 0

	for _, entry := range entries {
		entryRows := bytes.Split(entry, []byte{'\n'})

		reflectedRow := getReflected(entryRows, mandatoryDifferences)

		if reflectedRow == 0 {
			entryColumns := utils.Transpose2DSlice(entryRows)
			totalReflectedColumns += getReflected(entryColumns, mandatoryDifferences)
		} else {
			totalReflectedRows += reflectedRow
		}
	}

	return totalReflectedRows*100 + totalReflectedColumns
}

func getReflected(entry [][]byte, mandatoryDifferences int) int {
	for i := 0; i < len(entry)-1; i++ {
		differences := utils.FindDifferences(entry[i], entry[i+1])

		if differences < mandatoryDifferences+1 {
			left := i - 1
			right := i + 2

			isReflected := true
			for left >= 0 && right < len(entry) {
				differences += utils.FindDifferences(entry[left], entry[right])
				if differences > mandatoryDifferences {
					isReflected = false
					break
				}

				left--
				right++
			}

			if isReflected && differences == mandatoryDifferences {
				return i + 1
			}
		}
	}

	return 0
}
