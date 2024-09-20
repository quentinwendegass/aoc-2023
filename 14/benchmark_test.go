package main

import (
	"os"
	"testing"
)

func BenchmarkPartTwo(b *testing.B) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		dataCp := make([]byte, len(data))
		copy(dataCp, data)
		partTwo(dataCp)
	}
}
