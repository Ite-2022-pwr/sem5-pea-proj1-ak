package atsp

import (
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/utils"
	"math"
)

// BranchAndBoundSolver to struktura implementująca algorytm rozwiązujący problem ATSP za pomocą metody podziału i ograniczeń
type BranchAndBoundSolver struct {
	graph graph.Graph

	LowBound             int   // dolne ograniczenie
	UpperBound           int   // górne ograniczenie
	BestPath             []int // najlepsza ścieżka
	MinimumOutgoingEdges []int // minimalne wychodzące krawędzie z danego wierzchołka
}

// GetGraph zwraca graf, na którym działa algorytm
func (atsp *BranchAndBoundSolver) GetGraph() graph.Graph {
	return atsp.graph
}

// NewBranchAndBoundSolver tworzy nowy obiekt BranchAndBoundSolver
func NewBranchAndBoundSolver(g graph.Graph) *BranchAndBoundSolver {
	minimumOutgoingEdges := make([]int, g.GetVerticesCount())
	bestPath := make([]int, g.GetVerticesCount())
	minVal := math.MaxInt
	lowBound := 0

	// Obliczanie dolnego ograniczenia
	for i := 0; i < g.GetVerticesCount(); i++ {
		for j := 0; j < g.GetVerticesCount(); j++ {
			if i != j {
				weight, _ := g.GetEdge(i, j)
				if weight < minVal {
					// Znajdź minimalną wagę wychodzącej krawędzi z danego wierzchołka
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

// Solve rozwiązuje problem ATSP na grafie g, zaczynając od wierzchołka startVertex
func (atsp *BranchAndBoundSolver) Solve(startVertex int) (int, []int) {
	return atsp.BranchAndBound(startVertex)
}

// Node to struktura reprezentująca węzeł w drzewie przeszukiwania
type Node struct {
	Vertex     int // wierzchołek
	LowerBound int // dolne ograniczenie dla danego węzła
}

// BranchAndBound to funkcja rozwiązująca problem ATSP za pomocą metody podziału i ograniczeń
func (atsp *BranchAndBoundSolver) BranchAndBound(startVertex int) (int, []int) {
	startNode := Node{Vertex: startVertex, LowerBound: atsp.LowBound}
	visited := make([]bool, atsp.graph.GetVerticesCount())
	helperPath := make([]int, 0)

	return atsp.branchAndBoundRecursive(startNode, visited, helperPath)
}

// calculateLowerBound oblicza dolne ograniczenie dla danego wierzchołka
func (atsp *BranchAndBoundSolver) calculateLowerBound(node Node, nextVertex int) int {
	// Obliczenie dolnego ograniczenia dla następnego węzła polega na odjęciu minimalnej wychodzącej krawędzi z wierchołka
	// poprzedniego węzła oraz dodaniu wagi krawędzi między wierzchołkami poprzedniego i następnego węzła.
	// W momencie odwiedzenia wszystkich wierchołków, w węźle będącym liściem drzewa przeszukiwań, dolne ograniczenie
	// dla tego węzła to rzeczywisty koszt pokonanej ścieżki.
	return node.LowerBound - atsp.MinimumOutgoingEdges[node.Vertex] + atsp.graph.AsMatrix()[node.Vertex][nextVertex]
}

// branchAndBoundRecursive to funkcja rekurencyjna rozwiązująca problem ATSP za pomocą metody podziału i ograniczeń
func (atsp *BranchAndBoundSolver) branchAndBoundRecursive(node Node, visited []bool, helperPath []int) (int, []int) {
	helperPath = append(helperPath, node.Vertex) // Dodaj wierzchołek do ścieżki pomocniczej
	visited[node.Vertex] = true

	var tempNode Node

	// Kolejka priorytetowa węzłów
	bounds := utils.NewPriorityQueue(func(a, b Node) bool {
		return a.LowerBound < b.LowerBound
	})

	// Dla każdego nieodwiedzonego wierzchołka oblicz dolne ograniczenie i dodaj wierzchołek do kolejki
	for i := 0; i < atsp.graph.GetVerticesCount(); i++ {
		if !visited[i] {
			tempNode = Node{Vertex: i, LowerBound: atsp.calculateLowerBound(node, i)}
			bounds.Push(tempNode)
		}
	}

	// Jeśli nie ma wierzchołków do odwiedzenia, to sprawdź czy ścieżka jest lepsza od aktualnej najlepszej ścieżki
	if bounds.IsEmpty() {
		// Oblicz dolne ograniczenie podczas powrotu do wierzchołka startowego
		lowBound := atsp.calculateLowerBound(node, 0)

		// Jeśli znaleziono lepsze rozwiązanie, to zaktualizuj najlepsze rozwiązanie
		if lowBound < atsp.UpperBound {
			atsp.UpperBound = lowBound
			copy(atsp.BestPath, helperPath)
		}
	} else {
		for !bounds.IsEmpty() {
			tempNode = bounds.Pop() // Węzeł o najmniejszym dolnym ograniczeniu

			// Jeśli dolne ograniczenie jest mniejsze od górnego ograniczenia, to kontynuuj przeszukiwanie.
			// W przeciwnym wypadku przerwij przeszukiwanie dla danego węzła.
			if tempNode.LowerBound < atsp.UpperBound {
				atsp.branchAndBoundRecursive(tempNode, visited, helperPath)
			}
		}
	}

	// Powrót do poprzedniego wierzchołka
	visited[node.Vertex] = false
	helperPath = helperPath[:len(helperPath)-1]

	return atsp.UpperBound, atsp.BestPath
}
