package main

import (
	"aoc2023/utils"
	"math/rand"
	"slices"
	"strings"
)

type Graph map[int][]int

var nodeIDMap = make(map[string]int)
var nextID = 0

func main() {
	utils.WithAOC(partOne, func([]byte) int { return -1 }, utils.DefaultDataLoader)
}

// Can either be really fast or really slow, due to the random nature of the algorithm.
func partOne(data []byte) int {
	graph := constructGraph(data)

	for i := 0; i < 1000; i++ {
		graphCopy := deepCopyGraph(graph)
		cut, a, b := kargerMinCut(graphCopy)
		if cut == 3 {
			return a * b
		}
	}

	panic("cut of 3 not found. try again")
}

func kargerMinCut(graph Graph) (int, int, int) {
	vertices := len(graph)

	componentSize := make(map[int]int)
	for u := range graph {
		componentSize[u] = 1
	}

	for vertices > 2 {
		edges := getEdges(graph)
		if len(edges) == 0 {
			break
		}

		randomEdge := edges[rand.Intn(len(edges))]

		u := randomEdge[0]
		v := randomEdge[1]

		graph[u] = append(graph[u], graph[v]...)
		delete(graph, v)

		componentSize[u] += componentSize[v]
		delete(componentSize, v)

		for k := range graph {
			for i := range graph[k] {
				if graph[k][i] == v {
					graph[k][i] = u
				}
			}
		}

		removeSelfLoops(graph, u)

		vertices--
	}

	edges := getEdges(graph)
	remainingNodes := make([]int, 0, len(componentSize))
	for _, size := range componentSize {
		remainingNodes = append(remainingNodes, size)
	}

	return len(edges), remainingNodes[0], remainingNodes[1]
}

func getNodeID(node string) int {
	if id, exists := nodeIDMap[node]; exists {
		return id
	}
	nodeIDMap[node] = nextID
	nextID++
	return nodeIDMap[node]
}

func constructGraph(data []byte) Graph {
	graph := make(Graph)

	lines := strings.Split(string(data), "\n")

	input := map[string][]string{}

	for _, line := range lines {
		l := strings.Split(line, ": ")
		input[l[0]] = strings.Split(l[1], " ")
	}

	for key, values := range input {
		u := getNodeID(key)
		for _, value := range values {
			v := getNodeID(value)

			if !slices.Contains(graph[v], u) {
				graph[v] = append(graph[v], u)
			}

			if !slices.Contains(graph[u], v) {
				graph[u] = append(graph[u], v)
			}
		}
	}

	return graph
}

func removeSelfLoops(graph Graph, u int) {
	edges := graph[u]
	newEdges := []int{}
	for _, v := range edges {
		if v != u {
			newEdges = append(newEdges, v)
		}
	}
	graph[u] = newEdges
}

func getEdges(graph Graph) [][2]int {
	var edges [][2]int
	for u, neighbors := range graph {
		for _, v := range neighbors {
			if u < v {
				edges = append(edges, [2]int{u, v})
			}
		}
	}
	return edges
}

func deepCopyGraph(graph Graph) Graph {
	newGraph := make(Graph)
	for u, neighbors := range graph {
		newNeighbors := make([]int, len(neighbors))
		copy(newNeighbors, neighbors)
		newGraph[u] = newNeighbors
	}
	return newGraph
}
