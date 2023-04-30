package algorithm

import (
	"advocate-back/internal/algorithm/dfs"
	"advocate-back/internal/states"
	_ "fmt"
	_ "unicode/utf8"
)

func CheckAlgorithm(s *states.Bot) {
	g := dfs.MakeGraphFromConfig(s.BotStates)
	for _, vertex := range g.Vertices {
		if !vertex.Visited {
			g.Dfs(vertex)
		}
	}
}
