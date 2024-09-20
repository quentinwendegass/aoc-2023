package main

import (
	"aoc2023/utils"
	"strconv"
	"strings"
)

func main() {
	utils.WithAOC(partOneOptimized, partTwoOptimized, utils.DefaultDataLoader)
}

func partOneOptimized(data []byte) int {
	var MAX_BLUE byte = 14
	var MAX_RED byte = 12
	var MAX_GREEN byte = 13

	total := 0
	isGame := false
	currentNum := [2]byte{0, 0}
	failed := false
	lineDone := false
	idx := 1

	for _, b := range data {
		if b == '\n' {
			lineDone = false
			isGame = false
			if !failed {
				total += idx
			}
			failed = false
			idx++
			continue
		}

		if lineDone {
			continue
		}

		if b == ':' {
			isGame = true
			continue
		}

		if !isGame {
			continue
		}

		if b == 'r' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > MAX_RED {
				failed = true
				lineDone = true
				continue
			}
		}

		if b == 'g' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > MAX_GREEN {
				failed = true
				lineDone = true
				continue
			}
		}

		if b == 'b' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > MAX_BLUE {
				failed = true
				lineDone = true
				continue
			}
		}

		if b >= 48 && b <= 57 {
			if currentNum[0] == 0 {
				currentNum[0] = b
			} else {
				currentNum[1] = b
			}
		}
	}

	return total
}

func partTwoOptimized(data []byte) int {
	var total int = 0
	isGame := false
	currentNum := [2]byte{0, 0}
	var maxRed byte = 0
	var maxBlue byte = 0
	var maxGreen byte = 0

	for _, b := range data {
		if b == '\n' {
			isGame = false

			total += int(maxRed) * int(maxGreen) * int(maxBlue)
			maxRed = 0
			maxGreen = 0
			maxBlue = 0
			continue
		}

		if b == ':' {
			isGame = true
			continue
		}

		if !isGame {
			continue
		}

		if b == 'r' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > maxRed {
				maxRed = num
				continue
			}
		}

		if b == 'g' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > maxGreen {
				maxGreen = num
				continue
			}
		}

		if b == 'b' && currentNum[0] != 0 {
			num := convertToInt(currentNum)
			currentNum[0] = 0
			currentNum[1] = 0
			if num > maxBlue {
				maxBlue = num
				continue
			}
		}

		if b >= 48 && b <= 57 {
			if currentNum[0] == 0 {
				currentNum[0] = b
			} else {
				currentNum[1] = b
			}
		}
	}

	total += int(maxRed) * int(maxGreen) * int(maxBlue)

	return total
}

func partTwo(data []byte) int {
	lines := strings.Split(string(data), "\n")

	total := 0
	for _, line := range lines {
		labelIdx := strings.Index(line, ":")

		gamesString := line[labelIdx+1:]

		games := strings.Split(gamesString, ";")

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, game := range games {
			turns := strings.Split(game, ",")

			for _, turn := range turns {
				turn = strings.Trim(turn, " \n")
				idxAfterNum := strings.Index(turn, " ")
				num, _ := strconv.Atoi(turn[:idxAfterNum])

				if turn[idxAfterNum+1] == 'r' && num > maxRed {
					maxRed = num
				} else if turn[idxAfterNum+1] == 'g' && num > maxGreen {
					maxGreen = num
				} else if turn[idxAfterNum+1] == 'b' && num > maxBlue {
					maxBlue = num
				}
			}
		}

		total += maxRed * maxGreen * maxBlue
	}

	return total
}

func convertToInt(r [2]byte) byte {
	var num byte = 0
	if r[1] != 0 {
		num += (r[0] - 48) * 10
		num += (r[1] - 48)
	} else {
		num += (r[0] - 48)
	}
	return num
}
