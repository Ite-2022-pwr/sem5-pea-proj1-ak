package generator

import (
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"math/rand"
)

func GenerateAdjacencyMatrix(vertices int) (*graph.AdjacencyMatrix, error) {
	matrix, err := graph.NewAdjacencyMatrix(vertices)

	if err != nil {
		return nil, err
	}

	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			if i == j {
				continue
			}

			randomCost := rand.Intn(99) + 1
			if err = matrix.PutEdge(i, j, randomCost); err != nil {
				return nil, err
			}
		}
	}

	return matrix, nil
}
