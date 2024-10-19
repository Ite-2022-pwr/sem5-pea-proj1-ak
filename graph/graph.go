package graph

type Graph interface {
	PutEdge(sourceVertex, destinationVertex, weight int) error
	GetEdge(sourceVertex, destinationVertex int) (int, error)
	GetVerticesCount() int
	GetEdgesCount() int
	ToString() string
}
