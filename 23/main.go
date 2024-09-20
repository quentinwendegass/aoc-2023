package main

import (
	"aoc2023/utils"
	"bytes"
)

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	splitData := bytes.Split(data, []byte("\n"))
	graph := buildGraph(splitData)

	startNode := 1
	endNode := (len(splitData)-1)*len(splitData[0]) + len(splitData[0]) - 2
	longestPath, _ := graph.LongestPath(startNode, endNode)

	return len(longestPath) - 1
}

// Takes a some time to run, didn't feel like optimizing it. It's really inefficient currently since the graph is built in a very naive way.
func partTwo(data []byte) int {
	splitData := bytes.Split(data, []byte("\n"))
	splitData = replaceDirectionMarkers(splitData)
	graph := buildGraph(splitData)

	startNode := 1
	endNode := (len(splitData)-1)*len(splitData[0]) + len(splitData[0]) - 2
	longestPath, _ := graph.LongestPath(startNode, endNode)

	return len(longestPath) - 1
}

func buildGraph(data [][]byte) *utils.Graph {
	graph := utils.NewGraph()

	for y, row := range data {
		for x, cell := range row {
			if cell == '#' {
				continue
			}

			node := y*len(row) + x

			if y > 0 && (data[y-1][x] == '.' || data[y-1][x] == '^') {
				graph.AddEdge(node, node-len(row))
			}

			if y < len(data)-1 && (data[y+1][x] == '.' || data[y+1][x] == 'v') {
				graph.AddEdge(node, node+len(row))
			}

			if x > 0 && (data[y][x-1] == '.' || data[y][x-1] == '<') {
				graph.AddEdge(node, node-1)
			}

			if x < len(row)-1 && (data[y][x+1] == '.' || data[y][x+1] == '>') {
				graph.AddEdge(node, node+1)
			}
		}
	}

	return graph
}

func replaceDirectionMarkers(data [][]byte) [][]byte {
	data = utils.Copy2DSlice(data)
	for y, row := range data {
		for x, cell := range row {
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				data[y][x] = '.'
			}
		}
	}

	return data
}
