package graph

import "fmt"

type AdjacencyMatrix struct {
  Vertices int
  Matrix [][]int
}

func NewAdjacencyMatrix(vertices int) (*AdjacencyMatrix, error) {
  if vertices <= 0 {
    return nil, fmt.Errorf("vertices must be a positive number, given: %d", vertices)
  }

  am := &AdjacencyMatrix{Vertices: vertices}
  am.Matrix = make([][]int, am.Vertices)
  for i := 0; i < am.Vertices; i++ {
    am.Matrix[i] = make([]int, am.Vertices)
  }

  return am, nil
}

func (am *AdjacencyMatrix) PutEdge(startVertex, destinationVertex, weight int) error {
  if startVertex < 0 || startVertex >= am.Vertices {
    return fmt.Errorf("Wrong vertex value: %d", startVertex)
  }

  if destinationVertex < 0 || destinationVertex > am.Vertices {
    return fmt.Errorf("Wrong vertex value: %d", destinationVertex)
  }

  am.Matrix[startVertex][destinationVertex] = weight

  return nil
}

