package main

import (
	"aoc2023/utils"
	"bytes"
)

const (
	North byte = 0
	South byte = 1
	East  byte = 3
	West  byte = 4
)

type NodeMetadata struct {
	directions []byte
	x          int
	y          int
	value      byte
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	nodes, _ := getNodesInCycle(data)
	return len(nodes) / 2
}

func partTwo(data []byte) int {
	nodes, nodeDirectionsMap := getNodesInCycle(data)

	getCurrentDirection := func(from, to int) byte {
		for _, direction := range nodeDirectionsMap[nodes[from]].directions {
			if direction == North && hasSymbolDirection(South, nodeDirectionsMap[nodes[to]].directions) {
				return North
			} else if direction == South && hasSymbolDirection(North, nodeDirectionsMap[nodes[to]].directions) {
				return South
			} else if direction == East && hasSymbolDirection(West, nodeDirectionsMap[nodes[to]].directions) {
				return East
			} else if direction == West && hasSymbolDirection(East, nodeDirectionsMap[nodes[to]].directions) {
				return West
			}
		}

		panic("nodes in a cycle should always go into a direction")
	}

	var currentDirection byte = 5
	vertices := make([][]int, 0, 100)

	for i := 0; i < len(nodes)-1; i++ {
		newDirection := getCurrentDirection(i, i+1)
		if currentDirection != newDirection {
			currentDirection = newDirection
			vertices = append(vertices, []int{nodeDirectionsMap[nodes[i]].x, nodeDirectionsMap[nodes[i]].y})
		}
	}

	return utils.CountPointsInsideArea(utils.ShoelaceArea(vertices), len(nodes)-1)
}

func getNodesInCycle(data []byte) ([]int, map[int]NodeMetadata) {
	lines := bytes.Split(data, []byte{'\n'})
	graph := utils.NewGraph()

	nodeMetadataMap := make(map[int]NodeMetadata)

	startingNode := -1

	for i, line := range lines {
		for j, symbol := range line {
			node := i*len(lines[0]) + j

			symbolDirections := getDirectionsOfSymbol(symbol)
			connectedDirections := make([]byte, 0, 2)
			neighbours := make([]int, 0, 4)

			// north
			if hasSymbolDirection(North, symbolDirections) && i > 0 && hasSymbolDirection(South, getDirectionsOfSymbol(lines[i-1][j])) {
				neighbours = append(neighbours, node-len(lines[0]))
				connectedDirections = append(connectedDirections, North)
			}

			// south
			if hasSymbolDirection(South, symbolDirections) && i < len(lines)-1 && hasSymbolDirection(North, getDirectionsOfSymbol(lines[i+1][j])) {
				neighbours = append(neighbours, node+len(lines[0]))
				connectedDirections = append(connectedDirections, South)
			}

			// east
			if hasSymbolDirection(East, symbolDirections) && j < len(lines[0])-1 && hasSymbolDirection(West, getDirectionsOfSymbol(line[j+1])) {
				neighbours = append(neighbours, node+1)
				connectedDirections = append(connectedDirections, East)
			}

			// west
			if hasSymbolDirection(West, symbolDirections) && j > 0 && hasSymbolDirection(East, getDirectionsOfSymbol(line[j-1])) {
				neighbours = append(neighbours, node-1)
				connectedDirections = append(connectedDirections, West)
			}

			if symbol == 'S' {
				startingNode = node
			}

			nodeMetadataMap[node] = NodeMetadata{directions: connectedDirections, x: j, y: i, value: symbol}

			graph.AddEdges(node, neighbours)
		}
	}

	nodes, _ := graph.DetectCycleFrom(startingNode)

	return nodes, nodeMetadataMap
}

func hasSymbolDirection(char byte, directions []byte) bool {
	for _, direction := range directions {
		if char == direction {
			return true
		}
	}

	return false
}

func getDirectionsOfSymbol(char byte) []byte {
	directions := make([]byte, 0, 2)

	// north
	if char == '|' || char == 'L' || char == 'J' {
		directions = append(directions, North)
	}

	// south
	if char == '|' || char == '7' || char == 'F' {
		directions = append(directions, South)
	}

	// east
	if char == '-' || char == 'L' || char == 'F' {
		directions = append(directions, East)
	}

	// west
	if char == '-' || char == 'J' || char == '7' {
		directions = append(directions, West)
	}

	if char == 'S' {
		directions = append(directions, North, East, West, South)
	}

	return directions
}
