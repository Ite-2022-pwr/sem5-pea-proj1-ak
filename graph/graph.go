package graph

type Graph interface {
  PutEdge(sourceVertex, destinationVertex, weight int)
  GetEdge(sourceVertex, destinationVertex int) int
}
