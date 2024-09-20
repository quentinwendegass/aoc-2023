package main

import (
	"aoc2023/utils"
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	dataLoader := func() (io.Reader, func(), func()) {
		file, err := os.Open("./input.txt")
		if err != nil {
			log.Fatal("Cannot open input file", err)
		}

		return file, func() {
				file.Close()
			}, func() {
				file.Seek(0, 0)
			}
	}

	utils.WithAOC(partOne, partTwo, dataLoader)
}

func partOne(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	totalPoints := 0

	for scanner.Scan() {
		scratchcard := scanner.Text()

		winningNumbersCount := countMatchingNumbers(scratchcard)

		if winningNumbersCount > 0 {
			totalPoints += int(math.Pow(2, float64(winningNumbersCount-1)))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalPoints
}

func partTwo(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	scratchcards := [300]int{}
	currentScratchcard := 0

	for scanner.Scan() {
		scratchcard := scanner.Text()
		scratchcards[currentScratchcard]++

		winningNumbersCount := countMatchingNumbers(scratchcard)

		for i := 0; i < winningNumbersCount; i++ {
			scratchcards[currentScratchcard+1+i] += scratchcards[currentScratchcard]
		}

		currentScratchcard++
	}

	totalScratchcards := 0

	for i := 0; i < currentScratchcard; i++ {
		totalScratchcards += scratchcards[i]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalScratchcards
}

func countMatchingNumbers(scratchcard string) int {
	scratchcardNumbers := strings.Split(scratchcard, ":")[1]
	splitScratchcardNumbers := strings.Split(scratchcardNumbers, "|")
	winningNumbers := strings.Trim(splitScratchcardNumbers[0], " ")
	myNumbers := strings.Trim(splitScratchcardNumbers[1], " ")

	myNumbersSplit := strings.Split(myNumbers, " ")
	winningNumbersSplit := strings.Split(winningNumbers, " ")

	winningNumbersMap := make(map[string]bool, 100)
	for _, winningNumber := range winningNumbersSplit {
		if winningNumber != "" {
			winningNumbersMap[winningNumber] = true
		}
	}

	winningNumbersCount := 0

	for _, myNumber := range myNumbersSplit {
		if winningNumbersMap[myNumber] {
			winningNumbersCount++
		}
	}

	return winningNumbersCount
}
