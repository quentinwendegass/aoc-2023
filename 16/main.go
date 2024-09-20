package main

import (
	"aoc2023/utils"
	"bytes"
)

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	RIGHT Direction = 2
	LEFT  Direction = 3
)

type BeamPosition struct {
	x, y      int
	direction Direction
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	mirrorMap := bytes.Split(data, []byte{'\n'})

	return countEnergizedTiles(mirrorMap, BeamPosition{x: 0, y: 0, direction: RIGHT})
}

func partTwo(data []byte) int {
	mirrorMap := bytes.Split(data, []byte{'\n'})

	maxEnergizedTiles := 0

	for i := range mirrorMap {
		energizedTiles := countEnergizedTiles(mirrorMap, BeamPosition{x: 0, y: i, direction: RIGHT})

		if maxEnergizedTiles < energizedTiles {
			maxEnergizedTiles = energizedTiles
		}

		energizedTiles = countEnergizedTiles(mirrorMap, BeamPosition{x: len(mirrorMap[i]) - 1, y: i, direction: LEFT})

		if maxEnergizedTiles < energizedTiles {
			maxEnergizedTiles = energizedTiles
		}
	}

	for i := range mirrorMap[0] {
		energizedTiles := countEnergizedTiles(mirrorMap, BeamPosition{x: i, y: 0, direction: DOWN})

		if maxEnergizedTiles < energizedTiles {
			maxEnergizedTiles = energizedTiles
		}

		energizedTiles = countEnergizedTiles(mirrorMap, BeamPosition{x: i, y: len(mirrorMap) - 1, direction: UP})

		if maxEnergizedTiles < energizedTiles {
			maxEnergizedTiles = energizedTiles
		}
	}

	return maxEnergizedTiles
}

func countEnergizedTiles(mirrorMap [][]byte, start BeamPosition) int {
	visitedMap := make([][]bool, len(mirrorMap))
	for i := range visitedMap {
		visitedMap[i] = make([]bool, len(mirrorMap[i]))
	}

	cache := make(map[BeamPosition]bool)

	moveBeam(start, mirrorMap, visitedMap, cache)

	energizedTiles := 0
	for i := range visitedMap {
		for j := range visitedMap[i] {
			if visitedMap[i][j] {
				energizedTiles++
			}
		}
	}

	return energizedTiles
}

func moveBeam(pos BeamPosition, mirrorMap [][]byte, visitedMap [][]bool, cache map[BeamPosition]bool) {
	if pos.y < 0 || pos.y >= len(mirrorMap) {
		return
	}

	if pos.x < 0 || pos.x >= len(mirrorMap[pos.y]) {
		return
	}

	_, ok := cache[pos]

	if ok {
		return
	} else {
		cache[pos] = true
	}

	visitedMap[pos.y][pos.x] = true
	tile := mirrorMap[pos.y][pos.x]

	switch pos.direction {
	case UP:
		if tile == '.' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y - 1, direction: UP}, mirrorMap, visitedMap, cache)
			return
		}

		if tile == '\\' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x - 1, y: pos.y, direction: LEFT}, mirrorMap, visitedMap, cache)
		}

		if tile == '/' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x + 1, y: pos.y, direction: RIGHT}, mirrorMap, visitedMap, cache)
		}
	case DOWN:
		if tile == '.' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y + 1, direction: DOWN}, mirrorMap, visitedMap, cache)
			return
		}

		if tile == '\\' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x + 1, y: pos.y, direction: RIGHT}, mirrorMap, visitedMap, cache)
		}

		if tile == '/' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x - 1, y: pos.y, direction: LEFT}, mirrorMap, visitedMap, cache)
		}
	case RIGHT:
		if tile == '.' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x + 1, y: pos.y, direction: RIGHT}, mirrorMap, visitedMap, cache)
			return
		}

		if tile == '\\' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y + 1, direction: DOWN}, mirrorMap, visitedMap, cache)
		}

		if tile == '/' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y - 1, direction: UP}, mirrorMap, visitedMap, cache)
		}
	case LEFT:
		if tile == '.' || tile == '-' {
			moveBeam(BeamPosition{x: pos.x - 1, y: pos.y, direction: LEFT}, mirrorMap, visitedMap, cache)
			return
		}

		if tile == '\\' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y - 1, direction: UP}, mirrorMap, visitedMap, cache)
		}

		if tile == '/' || tile == '|' {
			moveBeam(BeamPosition{x: pos.x, y: pos.y + 1, direction: DOWN}, mirrorMap, visitedMap, cache)
		}
	}
}
