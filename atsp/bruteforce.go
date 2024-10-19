package atsp

import (
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"math"
)

type BruteForceSolver struct {
	graph graph.Graph
}

func (atsp *BruteForceSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func NewBruteForceSolver(g graph.Graph) *BruteForceSolver {
	return &BruteForceSolver{graph: g}
}

func (atsp *BruteForceSolver) Solve(startVertex int) (int, []int) {
	return atsp.BruteForce(startVertex)
}

func (atsp *BruteForceSolver) BruteForce(startVertex int) (int, []int) {
	visited := make([]bool, atsp.graph.GetVerticesCount())
	path := make([]int, 0, atsp.graph.GetVerticesCount())
	bestPath := make([]int, atsp.graph.GetVerticesCount())
	bestCost := math.MaxInt

	visited[startVertex] = true
	path = append(path, startVertex)

	bestCost, bestPath = atsp.bruteForceRecursive(startVertex, visited, startVertex, 0, bestCost, bestPath, path)

	return bestCost, bestPath
}

func (atsp *BruteForceSolver) bruteForceRecursive(startVertex int, visited []bool, currentVertex int, currentCost int, bestCost int, bestPath []int, path []int) (int, []int) {
	if len(path) == atsp.graph.GetVerticesCount() {
		cost, _ := atsp.graph.GetEdge(currentVertex, startVertex)
		currentCost += cost

		if currentCost < bestCost {
			bestCost = currentCost
			copy(bestPath, path)
		}

		return bestCost, bestPath
	}

	for i := 0; i < atsp.graph.GetVerticesCount(); i++ {
		if !visited[i] {
			visited[i] = true
			cost, _ := atsp.graph.GetEdge(currentVertex, i)
			currentCost += cost
			path = append(path, i)

			bestCost, bestPath = atsp.bruteForceRecursive(startVertex, visited, i, currentCost, bestCost, bestPath, path)

			visited[i] = false
			currentCost -= cost
			path = path[:len(path)-1]
		}
	}

	return bestCost, bestPath
}
