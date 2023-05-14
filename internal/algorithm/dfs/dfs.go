package dfs

import (
	"bot_forge_back/internal/states"
	"fmt"
)

type Vertex struct {
	neighbors []*Vertex
	Visited   bool
	Name      string
}

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) Dfs(v *Vertex) {
	v.Visited = true
	fmt.Println(v.Name)
	for _, neighbor := range v.neighbors {
		if !neighbor.Visited {
			g.Dfs(neighbor)
		}
	}
}

func MakeGraphFromConfig(s *states.BotStates) Graph {
	var graph = new(Graph)
	// adding all vertexes to graph
	for key := range s.States {
		vertexName := s.States[key].Name
		vertex := &Vertex{
			Name: vertexName,
		}
		graph.Vertices = append(graph.Vertices, vertex)
	}
	for _, vertex := range graph.Vertices {
		for _, state := range s.States {
			if vertex.Name == state.Name {
				for _, button := range state.Actions {
					for _, v := range graph.Vertices {
						if v.Name == button.NextBlock {
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
