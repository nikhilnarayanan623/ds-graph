package graph

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-graph/graph/interfaces"
	"github.com/nikhilnarayanan623/ds-stackAndQueue/queue"
	"github.com/nikhilnarayanan623/ds-stackAndQueue/stack"
)

type graph struct {
	// vertex and its edges
	graphMap map[int][]int
}

func GetNewGraph() interfaces.GraphInterface {

	return &graph{graphMap: map[int][]int{}}
}

// func to insert a new vertex and edge connection
// 1. function parameters are src,ds(vertex , edge) and IsbiDirectional to determine the connection is bidirectional or not
// 2. first check the srch(vertex) is present if not present then create a new vertex
// 3. then check the edge is present as a vertex in graph is not then create new vertex as edge
// 4. add dst (edge) to the list of src(vertex)
// 5. check the connection is bi directional if bidirectional then add src to the edge list of edge
func (g *graph) Insert(src, dst int, IsbiDriectional bool) {

	if _, ok := g.graphMap[src]; !ok {
		g.graphMap[src] = []int{}
	}

	if _, ok := g.graphMap[dst]; !ok {
		g.graphMap[dst] = []int{}
	}

	g.graphMap[src] = append(g.graphMap[src], dst)

	if IsbiDriectional {
		g.graphMap[dst] = append(g.graphMap[dst], src)
	}
}

// 1. range through the grapMap
// 2. in range return key as vertex and its value as edges
// 3. print the vertex then loop through the edges list and print it
func (g *graph) DisplayVertexWithEdges() {

	for vertex, edges := range g.graphMap {
		fmt.Print(vertex, ": ")
		for _, edge := range edges {
			fmt.Print(edge, " ")
		}
		fmt.Println()
	}
}

// 1. map[int]bool to store the visited vertex
// 2. imported package queue and using queue to add the edges of a vertex

// 3. range through the graph map and take all vertex in a variable vertex
// 4. check the vertex is visited or not if visited then continue the range
// 5. add the vertex to the queue

// 6. set a loop until the queue is empty
// 7. dequeu the queue and store the deuque value in currentVertex
// 9. if not visited then print the value and add the currentVertex to visited
// 10. get the list of  edges using currentVertex from graphMap
// 11. then range through the list of edges and store
// 12. when ranging check each edge is visited or not if not visited then add to queue

func (g *graph) BFS() {

	visited := map[int]bool{}

	queueOfVertex := queue.NewQueue()

	for vertex := range g.graphMap {

		if visited[vertex] {
			continue
		}

		queueOfVertex.Enqueue(vertex)

		for !queueOfVertex.IsQueueEmpty() {

			currentVertex := queueOfVertex.Dequeue()
			fmt.Print(currentVertex, " ")

			visited[currentVertex] = true

			for _, edge := range g.graphMap[currentVertex] {
				if !visited[edge] {
					queueOfVertex.Enqueue(edge)
				}
			}
		}
		fmt.Println()
	}
}

// 1. map[int]bool to store the visited vertex
// 2. imported package stack and using stack to add the edges of a vertex

// 3. range through the graph map and take all vertex in a variable vertex
// 4. check the vertex is visited or not if visited then continue the range
// 5. add the vertex to the queue

// 6. set a loop until the queue is empty
// 7. pop the stack and store the poped value in currentVertex
// 9. if not visited then print the value and add the currentVertex to visited
// 10. get the list of  edges using currentVertex from graphMap
// 11. then range through the list of edges and store
// 12. when ranging check each edge is visited or not if not visited then add to stack
func (g *graph) DFS() {

	visited := map[int]bool{}

	stackOfVertices := stack.NewStack()

	for vertex := range g.graphMap {

		if visited[vertex] {
			continue
		}

		stackOfVertices.Push(vertex)

		for !stackOfVertices.IsStackEmpty() {

			currentVertex := stackOfVertices.Pop()
			fmt.Print(currentVertex, " ")

			visited[currentVertex] = true

			edges := g.graphMap[currentVertex]

			for _, edge := range edges {
				if !visited[edge] {
					stackOfVertices.Push(edge)
				}
			}
		}
	}
	fmt.Println()
}
