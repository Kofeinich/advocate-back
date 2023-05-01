package algorithm

import (
	"advocate-back/internal/algorithm/dfs"
	"advocate-back/internal/states"
	"fmt"
	_ "fmt"
	_ "unicode/utf8"
)

func CheckAlgorithm(s states.BotStates) error {
	err := ValidateReferences(&s)
	if err != nil {
		return err
	}
	g := dfs.MakeGraphFromConfig(&s)
	for _, vertex := range g.Vertices {
		if vertex.Name == s.InitialState {
			g.Dfs(vertex)
			break
		}
	}
	var newVertices []*dfs.Vertex
	for _, vertex := range g.Vertices {
		if vertex.Visited {
			newVertices = append(newVertices, vertex)
			fmt.Println(vertex)
		}
	}
	g.Vertices = newVertices

	return nil
}
