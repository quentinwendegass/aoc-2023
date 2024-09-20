package utils

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

func (graph *Graph) AddEdges(node int, edgeNodes []int) {
	graph.adjacencyList[node] = append(graph.adjacencyList[node], edgeNodes...)
}

func (graph *Graph) AddEdge(node, edgeNode int) {
	graph.adjacencyList[node] = append(graph.adjacencyList[node], edgeNode)
}

func (graph *Graph) DetectCycleFrom(startNode int) ([]int, bool) {
	visited := make(map[int]bool)
	path := []int{}
	pathSet := make(map[int]bool)

	cycleNodes, foundCycle := graph.isCyclic(startNode, -1, visited, path, pathSet)
	return cycleNodes, foundCycle
}

func (graph *Graph) isCyclic(node, parentNode int, visited map[int]bool, path []int, pathSet map[int]bool) ([]int, bool) {
	visited[node] = true
	path = append(path, node)
	pathSet[node] = true

	for _, adjacentNode := range graph.adjacencyList[node] {
		if !visited[adjacentNode] {
			if cycleNodes, foundCycle := graph.isCyclic(adjacentNode, node, visited, path, pathSet); foundCycle {
				return cycleNodes, true
			}
		} else if pathSet[adjacentNode] && adjacentNode != parentNode {
			cycleStartIndex := -1
			for i, n := range path {
				if n == adjacentNode {
					cycleStartIndex = i
					break
				}
			}
			if cycleStartIndex != -1 {
				cycleNodes := path[cycleStartIndex:]
				return cycleNodes, true
			}
		}
	}

	pathSet[node] = false
	return nil, false
}

func (graph *Graph) LongestPath(startNode, endNode int) ([]int, bool) {
	visited := make(map[int]bool)
	currentPath := []int{startNode}
	var longestPath []int
	visited[startNode] = true

	graph.longestPathHelper(startNode, endNode, visited, currentPath, &longestPath)

	if len(longestPath) > 0 {
		return longestPath, true
	} else {
		return nil, false
	}
}

func (graph *Graph) longestPathHelper(currentNode, endNode int, visited map[int]bool, currentPath []int, longestPath *[]int) {
	if currentNode == endNode {
		if len(currentPath) > len(*longestPath) {
			*longestPath = make([]int, len(currentPath))
			copy(*longestPath, currentPath)
		}
		return
	}

	for _, neighbor := range graph.adjacencyList[currentNode] {
		if !visited[neighbor] {
			visited[neighbor] = true
			currentPath = append(currentPath, neighbor)

			graph.longestPathHelper(neighbor, endNode, visited, currentPath, longestPath)

			currentPath = currentPath[:len(currentPath)-1]
			visited[neighbor] = false
		}
	}
}
