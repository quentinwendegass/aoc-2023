package main

import (
	"os"
	"testing"
)

func BenchmarkPartOneForce(b *testing.B) {
	data := readInput()

	for i := 0; i < b.N; i++ {
		_ = partOne(data, calculateBetterDistanceCountForce)
	}
}

func BenchmarkPartOneMath(b *testing.B) {
	data := readInput()

	for i := 0; i < b.N; i++ {
		_ = partOne(data, calculateBetterDistanceCountMath)
	}
}

func BenchmarkPartTwoForce(b *testing.B) {
	data := readInput()

	for i := 0; i < b.N; i++ {
		_ = partTwo(data, calculateBetterDistanceCountForce)
	}
}

func BenchmarkPartTwoMath(b *testing.B) {
	data := readInput()

	for i := 0; i < b.N; i++ {
		_ = partTwo(data, calculateBetterDistanceCountMath)
	}
}

func readInput() []byte {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return data
}
