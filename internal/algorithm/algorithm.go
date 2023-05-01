package algorithm

import (
	"advocate-back/internal/algorithm/dfs"
	"advocate-back/internal/states"
	"fmt"
	_ "fmt"
	_ "unicode/utf8"
)

func CheckAlgorithm(s *states.Bot) error {
	err := ValidateReferences(s.BotStates)
	if err != nil {
		return err
	}
	g := dfs.MakeGraphFromConfig(s.BotStates)
	for _, vertex := range g.Vertices {
		if vertex.Name == s.BotStates.InitialState {
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
