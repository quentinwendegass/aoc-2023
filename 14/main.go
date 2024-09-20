package main

import (
	"aoc2023/utils"
	"bytes"
	"crypto/md5"
	"fmt"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(rawData []byte) int {
	data := make([]byte, len(rawData))
	copy(data, rawData)
	lines := bytes.Split(data, []byte{'\n'})

	rollNorth(lines)

	return countRocks(lines)
}

// Happy with this one. Runs really fast (~16ms on my i7 Macbook).
func partTwo(data []byte) int {
	lines := bytes.Split(data, []byte{'\n'})

	seenBefore := make(map[string]int)
	allPossibleStates := make([][]byte, 0, 100)
	const NumberOfCycles = 1_000_000_000

	for i := 0; i < NumberOfCycles; i++ {
		rollNorth(lines)
		rollWest(lines)
		rollSouth(lines)
		rollEastBetter(lines)

		h := md5.New()
		h.Write(data)
		key := fmt.Sprintf("%x", h.Sum(nil))

		repeatStart, ok := seenBefore[key]

		if !ok {
			dataCp := make([]byte, len(data))
			copy(dataCp, data)
			allPossibleStates = append(allPossibleStates, dataCp)
			seenBefore[key] = i
		} else {
			repeatingRange := len(allPossibleStates) - repeatStart
			idx := ((NumberOfCycles - 1 - repeatStart) % repeatingRange) + repeatStart
			return countRocks(bytes.Split(allPossibleStates[idx], []byte{'\n'}))
		}
	}

	return countRocks(lines)
}

func countRocks(lines [][]byte) int {
	total := 0
	for i, line := range lines {
		numberOfRocks := 0
		for _, pos := range line {
			if pos == 'O' {
				numberOfRocks++
			}
		}

		total += (len(lines) - i) * numberOfRocks
	}
	return total
}

func rollNorth(lines [][]byte) {
	for i := 1; i < len(lines); i++ {
		for j, pos := range lines[i] {
			if pos == 'O' {
				moveUp := 0
				for k := i - 1; k >= 0; k-- {
					if lines[k][j] == '.' {
						moveUp++
					} else {
						break
					}
				}

				if moveUp != 0 {
					lines[i][j] = '.'
					lines[i-moveUp][j] = 'O'
				}
			}
		}
	}
}

func rollSouth(lines [][]byte) {
	for i := len(lines) - 2; i >= 0; i-- {
		for j, pos := range lines[i] {
			if pos == 'O' {
				moveDown := 0
				for k := i + 1; k < len(lines); k++ {
					if lines[k][j] == '.' {
						moveDown++
					} else {
						break
					}
				}

				if moveDown != 0 {
					lines[i][j] = '.'
					lines[i+moveDown][j] = 'O'
				}
			}
		}
	}
}

func rollWest(lines [][]byte) {
	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			if line[i] == 'O' {
				moveLeft := 0
				for k := i - 1; k >= 0; k-- {
					if line[k] == '.' {
						moveLeft++
					} else {
						break
					}
				}

				if moveLeft != 0 {
					line[i] = '.'
					line[i-moveLeft] = 'O'
				}
			}
		}
	}
}

func rollEastBetter(lines [][]byte) {
	for _, line := range lines {
		for i := len(line) - 2; i >= 0; i-- {
			if line[i] == 'O' {
				moveRight := 0
				for k := i + 1; k < len(line); k++ {
					if line[k] == '.' {
						moveRight++
					} else {
						break
					}
				}

				if moveRight != 0 {
					line[i] = '.'
					line[i+moveRight] = 'O'
				}
			}
		}
	}
}
