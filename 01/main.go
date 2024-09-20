package main

import (
	"aoc2023/utils"
	"bytes"
	"strconv"
	"strings"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	lines := strings.Split(string(data), "\n")

	total := 0
	for _, line := range lines {
		var first byte
		var last byte

		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= 48 && c <= 57 {
				first = c
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= 48 && c <= 57 {
				last = c
				break
			}
		}

		var concat bytes.Buffer

		concat.WriteByte(first)
		concat.WriteByte(last)

		num, _ := strconv.Atoi(concat.String())
		total += num
	}

	return total
}

func partTwo(data []byte) int {
	var writtenOutNumbers = [][]byte{[]byte("zero"), []byte("one"), []byte("two"), []byte("three"), []byte("four"), []byte("five"), []byte("six"), []byte("seven"), []byte("eight"), []byte("nine")}

	lines := strings.Split(string(data), "\n")

	total := 0
	for _, line := range lines {
		var first byte
		var last byte
		var firstIdx int
		var lastIdx int

		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= 48 && c <= 57 {
				first = c
				firstIdx = i
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= 48 && c <= 57 {
				lastIdx = i
				last = c
				break
			}
		}

		for i := 0; i < len(writtenOutNumbers); i++ {
			firstWIdx := bytes.Index([]byte(line), writtenOutNumbers[i])
			lastWIdx := bytes.LastIndex([]byte(line), writtenOutNumbers[i])

			if firstWIdx != -1 && firstIdx > firstWIdx {
				firstIdx = firstWIdx
				first = byte(48 + i)
			}

			if lastWIdx != -1 && lastIdx < lastWIdx {
				lastIdx = lastWIdx
				last = byte(48 + i)
			}

		}

		combined := []byte{first, last}

		num, _ := strconv.Atoi(string(combined))
		total += num
	}

	return total
}
