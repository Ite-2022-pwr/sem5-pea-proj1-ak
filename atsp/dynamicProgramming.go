package atsp

import (
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"math"
)

type DynamicProgrammingSolver struct {
	graph graph.Graph

	visitedAll int
	memoCosts  [][]int
	memoPaths  [][]int
	start      int
}

func (atsp *DynamicProgrammingSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func NewDynamicProgrammingSolver(g graph.Graph) *DynamicProgrammingSolver {
	visitedAll := (1 << uint(g.GetVerticesCount())) - 1

	dp := make([][]int, 1<<uint(g.GetVerticesCount()))
	memo := make([][]int, 1<<uint(g.GetVerticesCount()))
	for i := 0; i < 1<<uint(g.GetVerticesCount()); i++ {
		dp[i] = make([]int, g.GetVerticesCount())
		memo[i] = make([]int, g.GetVerticesCount())
	}

	return &DynamicProgrammingSolver{
		graph:      g,
		visitedAll: visitedAll,
		memoCosts:  dp,
		memoPaths:  memo}
}

func (atsp *DynamicProgrammingSolver) Solve(startVertex int) (int, []int) {
	return atsp.DynamicProgramming(startVertex)
}

func (atsp *DynamicProgrammingSolver) DynamicProgramming(startVertex int) (int, []int) {
	for i := 0; i < 1<<uint(atsp.graph.GetVerticesCount()); i++ {
		for j := 0; j < atsp.graph.GetVerticesCount(); j++ {
			atsp.memoCosts[i][j] = -1
		}
	}
	atsp.start = startVertex

	cost := atsp.dynamicProgrammingRecursive(1<<uint(startVertex), startVertex)
	path := atsp.lookupPath()

	return cost, path
}

func (atsp *DynamicProgrammingSolver) dynamicProgrammingRecursive(mask, position int) int {
	if mask == atsp.visitedAll {
		cost, _ := atsp.graph.GetEdge(position, atsp.start)
		return cost
	}

	if atsp.memoCosts[mask][position] != -1 {
		return atsp.memoCosts[mask][position]
	}

	answer := math.MaxInt
	answerCity := -1

	for city := 0; city < atsp.graph.GetVerticesCount(); city++ {
		if mask&(1<<uint(city)) == 0 {
			cost, _ := atsp.graph.GetEdge(position, city)
			newAnswer := cost + atsp.dynamicProgrammingRecursive(mask|(1<<uint(city)), city)
			if newAnswer < answer {
				answer = newAnswer
				answerCity = city
			}
		}
	}

	atsp.memoCosts[mask][position] = answer
	atsp.memoPaths[mask][position] = answerCity

	return answer
}

func (atsp *DynamicProgrammingSolver) lookupPath() []int {
	mask := 1 << uint(atsp.start)

	path := make([]int, 0, atsp.graph.GetVerticesCount())
	path = append(path, atsp.start)
	lastCity := atsp.start

	for len(path) < atsp.graph.GetVerticesCount() {
		nextCity := atsp.memoPaths[mask][lastCity]
		mask |= 1 << uint(nextCity)
		path = append(path, nextCity)
		lastCity = nextCity
	}

	return path
}
