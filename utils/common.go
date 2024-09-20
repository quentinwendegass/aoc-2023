package utils

import (
	"fmt"
	"os"
	"strconv"
)

type Executor[DataType any, ReturnType any] func(DataType) ReturnType
type DataLoader[DataType any] func() (data DataType, cleanup func(), intermediate func())

func WithAOC[DataType any, ReturnType any](partOne Executor[DataType, ReturnType], partTwo Executor[DataType, ReturnType], dataLoader DataLoader[DataType]) {
	data, cleanup, intermediate := dataLoader()
	defer cleanup()

	PrintSolution(1, partOne(data))
	intermediate()
	PrintSolution(2, partTwo(data))
}

var DefaultDataLoader = func() ([]byte, func(), func()) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return data, func() {}, func() {}
}

func PrintSolution(solutionNumber int, solution any) {
	var color string
	if solutionNumber%2 == 0 {
		color = Red
	} else {
		color = Cyan
	}

	fmt.Println(color+fmt.Sprintf("Solution %d:", solutionNumber)+Reset, Bold, solution, Reset)
}

func ToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
