package main

import (
	"aoc2023/utils"
	"bytes"
	"math"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

type Galaxy struct {
	x float64
	y float64
}

func partOne(data []byte) int {
	return calculateTotalShortestPath(data, 1)
}

func partTwo(data []byte) int {
	return calculateTotalShortestPath(data, 1000000-1)
}

func calculateTotalShortestPath(data []byte, expansionRate int) int {
	lines := bytes.Split(data, []byte{'\n'})

	galaxies := make([]Galaxy, 0, 100)

	emptyRowsIdx, emptyColumnsIdx := getEmptyCells(lines)

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{x: float64(x), y: float64(y)})
			}
		}
	}

	totalShortestPath := 0

	for _, galaxyA := range galaxies {
		for _, galaxyB := range galaxies {
			if galaxyA == galaxyB {
				continue
			}

			expandedCellsX := countExpandedCellsBetweenGalaxyPos(galaxyA.x, galaxyB.x, emptyColumnsIdx) * expansionRate
			expandedCellsY := countExpandedCellsBetweenGalaxyPos(galaxyA.y, galaxyB.y, emptyRowsIdx) * expansionRate

			totalShortestPath += int(math.Max(galaxyA.x, galaxyB.x) - math.Min(galaxyA.x, galaxyB.x) + float64(expandedCellsX) + math.Max(galaxyA.y, galaxyB.y) - math.Min(galaxyA.y, galaxyB.y) + float64(expandedCellsY))
		}
	}
	return totalShortestPath / 2
}

func getEmptyCells(lines [][]byte) (rows []int, columns []int) {
	emptyRowsIdx := make([]int, 0, 100)
	emptyColumnsIdx := make([]int, 0, 100)
	notEmptyColumnMap := make([]bool, len(lines[0]))

	for y, line := range lines {
		isEmptyRow := true
		for x, char := range line {
			if char == '#' {
				isEmptyRow = false
				notEmptyColumnMap[x] = true
			}
		}

		if isEmptyRow {
			emptyRowsIdx = append(emptyRowsIdx, y)
		}
	}

	for i, notEmptyColumn := range notEmptyColumnMap {
		if !notEmptyColumn {
			emptyColumnsIdx = append(emptyColumnsIdx, i)
		}
	}

	return emptyRowsIdx, emptyColumnsIdx
}

func countExpandedCellsBetweenGalaxyPos(pos1, pos2 float64, expandedCells []int) int {
	stopPos := math.Max(pos1, pos2)
	startPos := math.Min(pos1, pos2)

	expandedCellCount := 0
	for _, cell := range expandedCells {
		if cell > int(startPos) && cell < int(stopPos) {
			expandedCellCount++
		}
	}

	return expandedCellCount
}
