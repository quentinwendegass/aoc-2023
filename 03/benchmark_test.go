package main

import (
	"os"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = partOne(data)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = partTwo(data)
	}
}
