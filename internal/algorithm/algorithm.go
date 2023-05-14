package algorithm

import (
	"bot_forge_back/internal/algorithm/dfs"
	"bot_forge_back/internal/states"
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
