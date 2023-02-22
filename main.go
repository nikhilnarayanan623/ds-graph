package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-graph/graph"
)

func main() {

	gr := graph.GetNewGraph()

	gr.Insert(1, 2, false)
	gr.Insert(1, 3, false)

	gr.Insert(2, 4, false)
	gr.Insert(3, 5, false)

	gr.Insert(5, 4, false)
	gr.Insert(5, 7, false)
	gr.Insert(5, 6, false)

	gr.Insert(7, 4, false)

	fmt.Println("All vertex with its edges ")
	gr.DisplayVertexWithEdges()
	// fmt.Println("\nAll pathes from ", 1, "to ", 4)
	// fmt.Println(gr.AllPathFromOneVertexToAnother(1, 4))

	// fmt.Println("DFS")
	// gr.DFS()
}
