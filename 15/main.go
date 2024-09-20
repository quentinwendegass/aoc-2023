package main

import (
	"aoc2023/utils"
	"bytes"
	"slices"
	"strconv"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	initSteps := bytes.Split(data, []byte{','})

	sumHashes := 0
	for _, initStep := range initSteps {
		hash := 0
		for _, char := range initStep {
			hash = calculateHash(char, hash)
		}

		sumHashes += hash
	}
	return sumHashes
}

func partTwo(data []byte) int {
	initSteps := bytes.Split(data, []byte{','})

	boxes := make(map[int][][2]string)

	for _, initStep := range initSteps {
		label := 0
		for i, char := range initStep {
			if char == '=' {
				focalLength := string(initStep[i+1])
				key := string(initStep[0:i])

				foundIdx := slices.IndexFunc(boxes[label], func(slice [2]string) bool {
					return slice[0] == key
				})

				if foundIdx != -1 {
					boxes[label][foundIdx][1] = focalLength
				} else {
					boxes[label] = append(boxes[label], [2]string{key, focalLength})
				}
				break
			} else if char == '-' {
				key := string(initStep[0:i])

				foundIdx := slices.IndexFunc(boxes[label], func(slice [2]string) bool {
					return slice[0] == key
				})

				if foundIdx != -1 {
					boxes[label] = utils.RemoveIndex(boxes[label], foundIdx)
				}
				break
			}
			label = calculateHash(char, label)
		}
	}

	sumFocusingPower := 0
	for boxNumber, box := range boxes {
		for i, lens := range box {
			focalLength, err := strconv.Atoi(lens[1])
			if err != nil {
				panic(err)
			}

			sumFocusingPower += (boxNumber + 1) * (i + 1) * focalLength
		}
	}

	return sumFocusingPower
}

func calculateHash(char byte, initHash int) int {
	initHash += int(char)
	initHash *= 17
	initHash %= 256

	return initHash
}
