package main

import (
	"math"
	"sort"
)

type ucshelper struct {
	grid *Grid
	cost int
	arah Arah
	path []ucshelper
}

func (h ucshelper) calculateCost() int {
	arr := h.path
	total := 0
	for _, v := range arr {
		if v.grid == nil {
			continue
		}
		total += v.grid.cost
	}
	total += h.grid.cost
	return total
}
func (p Player) ucs() {
	save := p
	// step := 1
	queue := make([]ucshelper, 0)
	current := ucshelper{}
	for true {
		for _, v := range Allarah {
			temp2 := p
			err := temp2.move(v)
			if err != nil {
				queue = append(queue, ucshelper{path: nil, grid: nil, cost: math.MaxInt64, arah: v})
				continue
			}
			queue = append(queue, ucshelper{path: append(current.path, current), grid: temp2.position, arah: v})
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].calculateCost() < queue[j].calculateCost()
		})
		p.position = queue[0].grid
		current = queue[0]
		if p.position.tipe == TipeGoal {
			current.printucspath(save)
			return
		}
		queue = queue[1:]
	}
}

func (u ucshelper) printucspath(player Player) {
	for i := 1; i < len(u.path); i++ {
		player.position.tipe = TipeEmpty
		player.move(u.path[i].arah)
		player.position.tipe = TipeStart
		println()
		println(u.path[i].arah)
		peta.printGrid()
	}
	player.position.tipe = TipeEmpty
	player.move(u.arah)
	player.position.tipe = TipeStart
	println()
	println(u.arah)
	peta.printGrid()
}
