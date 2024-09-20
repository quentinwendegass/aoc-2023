package main

import (
	"aoc2023/utils"
	"bytes"
	"container/heap"
	"strconv"
)

type Direction int

const (
	NONE Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

type State struct {
	x, y      int
	direction Direction
	steps     int
	totalHeat int
}

type StateKey struct {
	x, y      int
	direction Direction
	steps     int
	minSteps  int
	maxSteps  int
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	return findMinimumHeat(convertToIntMatrix(data), 1, 3)
}

func partTwo(data []byte) int {
	return findMinimumHeat(convertToIntMatrix(data), 4, 10)
}

func findMinimumHeat(heatMap [][]int, minSteps, maxSteps int) int {
	rows := len(heatMap)
	cols := len(heatMap[0])
	directions := []Direction{UP, DOWN, LEFT, RIGHT}

	pq := &utils.PriorityQueue[State]{}
	heap.Init(pq)
	heap.Push(pq, &utils.Item[State]{
		Value:    State{x: 0, y: 0, direction: NONE, steps: 0, totalHeat: heatMap[0][0]},
		Priority: heatMap[0][0],
	})

	visited := make(map[StateKey]int)

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*utils.Item[State])
		currentState := currentItem.Value

		if currentState.x == cols-1 && currentState.y == rows-1 {
			return currentState.totalHeat - heatMap[0][0]
		}

		stateKey := StateKey{
			x:         currentState.x,
			y:         currentState.y,
			direction: currentState.direction,
			steps:     currentState.steps,
			minSteps:  minSteps,
			maxSteps:  maxSteps,
		}

		if heat, ok := visited[stateKey]; ok && heat <= currentState.totalHeat {
			continue
		}

		visited[stateKey] = currentState.totalHeat

		for _, dir := range directions {
			if isReverseDirection(currentState.direction, dir) {
				continue
			}

			newSteps := currentState.steps
			if dir == currentState.direction || currentState.direction == NONE {
				newSteps++
			} else {
				if currentState.steps < minSteps {
					continue
				}
				newSteps = 1
			}

			if newSteps > maxSteps {
				continue
			}

			newX, newY := currentState.x, currentState.y
			switch dir {
			case UP:
				newY--
			case DOWN:
				newY++
			case LEFT:
				newX--
			case RIGHT:
				newX++
			}

			if newX >= 0 && newX < cols && newY >= 0 && newY < rows {
				newTotalHeat := currentState.totalHeat + heatMap[newY][newX]
				newState := State{
					x:         newX,
					y:         newY,
					direction: dir,
					steps:     newSteps,
					totalHeat: newTotalHeat,
				}

				heap.Push(pq, &utils.Item[State]{
					Value:    newState,
					Priority: newTotalHeat,
				})
			}
		}
	}

	return -1
}

func isReverseDirection(d1, d2 Direction) bool {
	return (d1 == UP && d2 == DOWN) || (d1 == DOWN && d2 == UP) ||
		(d1 == LEFT && d2 == RIGHT) || (d1 == RIGHT && d2 == LEFT)
}

func convertToIntMatrix(data []byte) [][]int {
	lines := bytes.Split(data, []byte{'\n'})

	heatMap := make([][]int, len(lines))

	for i, line := range lines {
		heatMap[i] = make([]int, len(lines[0]))

		for j, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			heatMap[i][j] = val
		}
	}

	return heatMap
}
