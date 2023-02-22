package graph

// 1. function arguments are two integers as ther vertext in the graph src,dst(source and vertex) and return all path as multi dimensional array
// 2. initialize a multi dimensioanl allPath to store all path
// 2. first check the src and dst are present in the graph if one of them not present then return empty 2d array
// 3. if src and dst are present then get the edges of array from graphMap to variable edges
// 4. initialize a 1d array path with first value as src
// 5. initialize a map[int]bool to store visited vertex
// using an extra function called findPath to find all path using recursion(DFS)
// 6. findPath function take integer vertex , dst , map[int]bool  visited  , address of 1d array path and 2d array allPaths
// 7. base conditions if vertex is contain in visited map then return
// 8. if not then add int map visited
// 9. check the vertex is the dst then take a copy of the path and append it on allPath
// 10. its all false take vertex's edges and range through it and store each edges on variable edge
// 11. call the same findPath function with vertex as edge ,dst,visited,array path and allPath
// 12. delete the edge from visited
// 13. return the allPath
func (g *graph) AllPathFromOneVertexToAnother(src, dst int) [][]int {

	_, ok1 := g.graphMap[src]
	_, ok2 := g.graphMap[dst]

	if !ok1 || !ok2 {
		return [][]int{}
	}

	allPath := [][]int{}
	path := []int{}

	g.findPaths(src, dst, &map[int]bool{}, path, &allPath)

	return allPath
}

func (g *graph) findPaths(vertex, dst int, visited *map[int]bool, path []int, allPath *[][]int) {

	if (*visited)[vertex] {
		return
	}

	(*visited)[vertex] = true
	path = append(path, vertex)

	if vertex == dst {

		copyPath := make([]int, len(path))
		copy(copyPath, path)

		(*allPath) = append((*allPath), copyPath)
		return
	}

	edges := g.graphMap[vertex]

	for _, edge := range edges {

		g.findPaths(edge, dst, visited, path, allPath)

		delete((*visited), edge)
	}

}
