package algorithm

import (
	dfs "advocate-back/internal/algorithm/dfs"
	"advocate-back/internal/states"
	_ "fmt"
	_ "unicode/utf8"
)

func CheckAlgorithm(s *states.BotStates) {
	g := dfs.MakeGraphFromConfig(s)
	for _, vertex := range g.Vertices {
		if !vertex.Visited {
			g.Dfs(vertex)
		}
	}
}
