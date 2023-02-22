package interfaces

type GraphInterface interface {
	Insert(src, dst int, biDirectional bool) // srch is verted and dst is edge
	BFS()                                    // Bridth First Search
	DFS()                                    // Depth First Search
	DisplayVertexWithEdges()                 // to display graph as vertex with its edges

	AllPathFromOneVertexToAnother(src, dst int) [][]int // give two vertex return all possible paths in unidirectional graph
}
