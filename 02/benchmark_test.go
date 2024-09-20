package main

import (
	"os"
	"testing"
)

func BenchmarkPartOneOptimized(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = partOneOptimized(data)
	}
}

func BenchmarkPartTwoOptimized(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = partTwoOptimized(data)
	}
}

func BenchmarkPartTwoUnoptimized(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = partTwo(data)
	}
}
