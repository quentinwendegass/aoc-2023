package utils

import (
	"slices"
	"strconv"
)

func AToISlice(strValues []string) []int {
	values := make([]int, len(strValues))

	for i, strValue := range strValues {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			panic(err)
		}
		values[i] = value
	}

	return values
}

func ASliceToISlice(asciArray [][]byte) []int {
	asciConverted := make([]int, len(asciArray))

	for i, code := range asciArray {
		val, err := strconv.Atoi(string(code))
		if err != nil {
			panic(err)
		}
		asciConverted[i] = val
	}

	return asciConverted
}

func Transpose2DSlice[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func FindDifferences[T comparable](a, b []T) int {
	differences := 0
	for i := range a {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences
}

func RemoveIndex[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func Copy2DSlice[T any](slice [][]T) [][]T {
	newSlice := make([][]T, len(slice))
	for i, row := range slice {
		newSlice[i] = make([]T, len(row))
		copy(newSlice[i], row)
	}
	return newSlice
}

func FindCommonElements[T comparable](s1, s2 []T) []T {
	result := make([]T, 0)

	for _, val := range s1 {
		if slices.Contains(s2, val) {
			result = append(result, val)
		}
	}

	return result
}
