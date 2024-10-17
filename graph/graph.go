package graph

type Graph interface {
	PutEdge(sourceVertex, destinationVertex, weight int)
	GetEdge(sourceVertex, destinationVertex int) (int, error)
	GetVerticesCount() int
	GetEdgesCount() int
}
