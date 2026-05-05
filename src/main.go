package main

import(
	"stima/core"
)

func main() {
	firstgrid, start, end, _ := core.CreateGrid()
	player := core.Player{Position: start}

	result := player.UCS(end)
	result.PrintResultPath(player, firstgrid)
}
