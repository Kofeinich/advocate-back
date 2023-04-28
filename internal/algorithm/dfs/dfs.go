package dfs

import (
	"advocate-back/internal/states"
	"fmt"
)

type Vertex struct {
	neighbors []*Vertex
	Visited   bool
	name      string
}

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) Dfs(v *Vertex) {
	v.Visited = true
	fmt.Println(v.name)
	for _, neighbor := range v.neighbors {
		if !neighbor.Visited {
			g.Dfs(neighbor)
		}
	}
}

func MakeGraphFromConfig(s *states.BotStates) Graph {
	var graph = new(Graph)

	for key := range s.States {
		vertexName := s.States[key].Name
		vertex := &Vertex{
			name: vertexName,
		}
		graph.Vertices = append(graph.Vertices, vertex)
	}

	for _, vertex := range graph.Vertices {
		for _, state := range s.States {
			if vertex.name == state.Name {
				for _, button := range state.Buttons {
					for _, v := range graph.Vertices {
						if v.name == button.NextBlock {
							vertex.neighbors = append(vertex.neighbors, v)
							break
						}
					}
				}
				break
			}
		}
	}
	return *graph
}
