package main

import (
	"aoc2023/utils"
	"bytes"
	"fmt"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	lines := bytes.Split(data, []byte{'\n'})

	total := 0
	for _, line := range lines {
		arrangements, brokenCode := transformInputLine(line)
		total += getValidCodesForce(brokenCode, arrangements)
	}

	return total
}

func partTwo(data []byte) int {
	lines := bytes.Split(data, []byte{'\n'})

	total := 0
	for _, line := range lines {
		arrangements, brokenCode := transformInputLine(line)
		brokenCode = bytes.Join([][]byte{brokenCode, brokenCode, brokenCode, brokenCode, brokenCode}, []byte{'?'})

		duplicatedArrangements := make([]int, 0, len(arrangements)*5)
		for i := 0; i < 5; i++ {
			duplicatedArrangements = append(duplicatedArrangements, arrangements...)
		}
		arrangements = duplicatedArrangements

		total += getValidCodesRecursive(brokenCode, arrangements, make(map[string]int))
	}

	return total
}

func getValidCodesForce(brokenCode []byte, arrangements []int) int {
	unknownPositions := make([]int, 0, 10)

	for i, pos := range brokenCode {
		if pos == '?' {
			unknownPositions = append(unknownPositions, i)
		}
	}

	validCodes := 0
	numStates := 1 << len(unknownPositions)

	for i := 0; i < numStates; i++ {
		stateArray := make([]byte, len(brokenCode))
		copy(stateArray, brokenCode)

		for j, pos := range unknownPositions {
			if (i>>j)&1 == 1 {
				stateArray[pos] = '.'
			} else {
				stateArray[pos] = '#'
			}
		}

		if isValid(string(stateArray), arrangements) {
			validCodes++
		}
	}

	return validCodes
}

func isValid(code string, arrangements []int) bool {
	currentHashes := 0
	currentArrangementIdx := 0

	for _, pos := range code {
		if pos == '#' {
			currentHashes++
			continue
		}

		if currentHashes != 0 {
			if len(arrangements) <= currentArrangementIdx || arrangements[currentArrangementIdx] != currentHashes {
				return false
			}
			currentHashes = 0
			currentArrangementIdx++
		}
	}

	if currentHashes != 0 {
		if len(arrangements) <= currentArrangementIdx || arrangements[currentArrangementIdx] != currentHashes {
			return false
		}
		currentHashes = 0
		currentArrangementIdx++
	}

	if currentArrangementIdx != len(arrangements) {
		return false
	}

	return true
}

// More or less copied from https://medium.com/@jatinkrmalik/day-12-hot-springs-advent-of-code-2023-python-77506773abfb
// Couldn't get my brain to implementing a recursive solution with memoization.
func getValidCodesRecursive(code []byte, arr []int, cache map[string]int) int {
	key := fmt.Sprint(code, arr)

	val, ok := cache[key]
	if ok {
		return val
	}
	if len(arr) == 0 {
		if bytes.Contains(code, []byte{'#'}) {
			return 0
		} else {
			return 1
		}
	}

	if len(code) == 0 {
		return 0
	}

	total := 0

	if code[0] == '.' || code[0] == '?' {
		total += getValidCodesRecursive(code[1:], arr, cache)
	}

	if code[0] == '#' || code[0] == '?' {
		if valid(code, arr) {
			newCode := []byte{}

			if arr[0]+1 < len(code) {
				newCode = code[arr[0]+1:]
			}
			total += getValidCodesRecursive(newCode, arr[1:], cache)
		}
	}

	cache[key] = total
	return total
}

func valid(code []byte, arr []int) bool {
	return arr[0] <= len(code) && !bytes.Contains(code[0:arr[0]], []byte{'.'}) && (arr[0] == len(code) || code[arr[0]] != '#')
}

func transformInputLine(line []byte) ([]int, []byte) {
	a := bytes.Split(line, []byte{' '})

	arrangements := bytes.Split(a[1], []byte{','})
	brokenCode := a[0]

	arrangementsConverted := utils.ASliceToISlice(arrangements)

	return arrangementsConverted, brokenCode
}
