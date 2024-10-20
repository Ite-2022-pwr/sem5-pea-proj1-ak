package atsp

import (
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/utils"
	"math"
)

type BranchAndBoundSolver struct {
	graph graph.Graph

	LowBound             int
	UpperBound           int
	BestPath             []int
	MinimumOutgoingEdges []int
}

func (atsp *BranchAndBoundSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func NewBranchAndBoundSolver(g graph.Graph) *BranchAndBoundSolver {
	minimumOutgoingEdges := make([]int, g.GetVerticesCount())
	bestPath := make([]int, g.GetVerticesCount())
	minVal := math.MaxInt
	lowBound := 0

	for i := 0; i < g.GetVerticesCount(); i++ {
		for j := 0; j < g.GetVerticesCount(); j++ {
			if i != j {
				weight, _ := g.GetEdge(i, j)
				if weight < minVal {
					minVal = weight
				}
			}
		}

		minimumOutgoingEdges[i] = minVal
		lowBound += minVal
		minVal = math.MaxInt
	}

	return &BranchAndBoundSolver{
		graph:                g,
		LowBound:             lowBound,
		UpperBound:           math.MaxInt,
		MinimumOutgoingEdges: minimumOutgoingEdges,
		BestPath:             bestPath,
	}
}

func (atsp *BranchAndBoundSolver) Solve(startVertex int) (int, []int) {
	return atsp.BranchAndBound(startVertex)
}

type Node struct {
	Vertex     int
	LowerBound int
}

func (atsp *BranchAndBoundSolver) BranchAndBound(startVertex int) (int, []int) {
	startNode := Node{Vertex: startVertex, LowerBound: atsp.LowBound}
	visited := make([]bool, atsp.graph.GetVerticesCount())
	helperPath := make([]int, 0)

	return atsp.branchAndBoundRecursive(startNode, visited, helperPath)
}

func (atsp *BranchAndBoundSolver) calculateLowerBound(node Node, nextVertex int) int {
	return node.LowerBound - atsp.MinimumOutgoingEdges[node.Vertex] + atsp.graph.AsMatrix()[node.Vertex][nextVertex]
}

func (atsp *BranchAndBoundSolver) branchAndBoundRecursive(node Node, visited []bool, helperPath []int) (int, []int) {
	helperPath = append(helperPath, node.Vertex)
	visited[node.Vertex] = true

	var tempNode Node
	bounds := utils.NewPriorityQueue(func(a, b Node) bool {
		return a.LowerBound < b.LowerBound
	})

	for i := 0; i < atsp.graph.GetVerticesCount(); i++ {
		if !visited[i] {
			tempNode = Node{Vertex: i, LowerBound: atsp.calculateLowerBound(node, i)}
			bounds.Push(tempNode)
		}
	}

	if bounds.IsEmpty() {
		lowBound := atsp.calculateLowerBound(node, 0)
		if lowBound < atsp.UpperBound {
			atsp.UpperBound = lowBound
			copy(atsp.BestPath, helperPath)
		}
	} else {
		for !bounds.IsEmpty() {
			tempNode = bounds.Pop()
			if tempNode.LowerBound < atsp.UpperBound {
				atsp.branchAndBoundRecursive(tempNode, visited, helperPath)
			}
		}
	}

	visited[node.Vertex] = false
	helperPath = helperPath[:len(helperPath)-1]

	return atsp.UpperBound, atsp.BestPath
}
