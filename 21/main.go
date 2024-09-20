package main

import (
	"aoc2023/utils"
	"bytes"
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	grid := bytes.Split(data, []byte{'\n'})
	startPoint := findStartingPoint(grid)

	distanceMap := findDistanceAfterSteps(startPoint, grid)

	pointsCount := 0
	for _, distance := range distanceMap {
		if distance <= 64 && distance%2 == 0 {
			pointsCount++
		}
	}

	return pointsCount
}

func partOneForce(data []byte) int {
	grid := bytes.Split(data, []byte{'\n'})

	startPoint := findStartingPoint(grid)

	pointsCount := 0
	findAvailablePositionsAfterXSteps(startPoint, grid, 1, 64, make(map[string]bool), &pointsCount)
	return pointsCount
}

// Couldn't figure this one out myself. Solution from https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21.
// Didn't like this one, since it assumes the input is structured in a certain way, but it's not stated in the problem description.
func partTwo(data []byte) int {
	grid := bytes.Split(data, []byte{'\n'})

	startPoint := findStartingPoint(grid)

	distanceMap := findDistanceAfterSteps(startPoint, grid)

	distanceToEdge := len(grid) / 2

	if distanceToEdge != 65 {
		panic("invalid input")
	}

	n := (26501365 - distanceToEdge) / len(grid)

	if n != 202300 {
		panic("invalid input")
	}

	oddPointCount := (n + 1) * (n + 1)
	evenPointCount := n * n

	oddCornerCount := 0
	evenCornerCount := 0
	oddDistances := 0
	evenDistances := 0

	for _, distance := range distanceMap {
		if distance%2 == 1 {
			if distance > distanceToEdge {
				oddCornerCount++
			}
			oddDistances++
		} else {
			if distance > distanceToEdge {
				evenCornerCount++
			}
			evenDistances++
		}
	}

	return oddPointCount*oddDistances + evenPointCount*evenDistances - ((n + 1) * oddCornerCount) + (n * evenCornerCount)
}

func findAvailablePositionsAfterXSteps(point Point, grid [][]byte, currentStep int, maxSteps int, cache map[string]bool, pointsCount *int) {
	cacheKey := fmt.Sprint(point.x, point.y, currentStep)

	if _, ok := cache[cacheKey]; !ok {
		cache[cacheKey] = true
	} else {
		return
	}

	if currentStep > maxSteps {
		*pointsCount++
		return
	}

	nextPoints := getNextPoints(point)

	for _, nextPoint := range nextPoints {
		if !isPointValid(nextPoint, grid) {
			continue
		}

		findAvailablePositionsAfterXSteps(nextPoint, grid, currentStep+1, maxSteps, cache, pointsCount)
	}
}

func findDistanceAfterSteps(start Point, grid [][]byte) map[Point]int {
	pointQueue := utils.CreateQueue[Point]()
	distanceQueue := utils.CreateQueue[int]()

	distanceMap := make(map[Point]int)

	pointQueue.Push(start)
	distanceQueue.Push(0)

	for pointQueue.Len() > 0 {
		point := pointQueue.Pop()
		distance := distanceQueue.Pop()

		if _, ok := distanceMap[point]; ok {
			continue
		}

		distanceMap[point] = distance

		nextPoints := getNextPoints(point)

		for _, nextPoint := range nextPoints {
			if _, ok := distanceMap[nextPoint]; !isPointValid(nextPoint, grid) || ok {
				continue
			}

			pointQueue.Push(nextPoint)
			distanceQueue.Push(distance + 1)
		}
	}

	return distanceMap
}

func getNextPoints(point Point) []Point {
	return []Point{{x: point.x - 1, y: point.y}, {x: point.x + 1, y: point.y}, {x: point.x, y: point.y - 1}, {x: point.x, y: point.y + 1}}
}

func isPointValid(point Point, grid [][]byte) bool {
	if point.x < 0 || point.x >= len(grid[0]) || point.y < 0 || point.y >= len(grid) {
		return false
	}

	return grid[point.y][point.x] != '#'
}

func findStartingPoint(grid [][]byte) Point {
	for y := range grid {
		x := bytes.Index(grid[y], []byte{'S'})

		if x != -1 {
			return Point{x: x, y: y}
		}
	}

	panic("no start position found")
}
