package main

import (
	"encoding/json"
	"fmt"
)

type (
	Graph struct {
		VertexesNum  int                  `json:"num"`
		AdjacentList map[int]map[int]bool `json:"adjacent_list"`
	}
)

func NewGrah() *Graph {
	return &Graph{
		AdjacentList: map[int]map[int]bool{},
	}
}

func (g *Graph) addVertex(value int) {
	if _, ok := g.AdjacentList[value]; !ok {
		g.AdjacentList[value] = map[int]bool{}
		g.VertexesNum++
	}
}

func (g *Graph) addEdge(node1, node2 int) {
	// make sure the vertexes are already existing
	g.addVertex(node1)
	g.addVertex(node2)

	if _, exists := g.AdjacentList[node1][node2]; !exists {
		g.AdjacentList[node1][node2] = true
	}

	if _, exists := g.AdjacentList[node2][node1]; !exists {
		g.AdjacentList[node2][node1] = true
	}
}

func (g *Graph) Print() {
	jstr, err := json.MarshalIndent(g.AdjacentList, "", " ")
	if err != nil {
		fmt.Println("\n\n Error in marshaling the adjacent list: ", err.Error())
	}

	fmt.Println("\n\njstr is: ", string(jstr))
}

func main() {
	g := NewGrah()

	g.addVertex(0)
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)
	g.addVertex(6)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(3, 4)
	g.addEdge(2, 4)
	g.addEdge(4, 5)
	g.addEdge(5, 6)

	g.Print()
}
