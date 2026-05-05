package main

import (
	"sort"
)

type ucshelper struct {
	grid *Grid
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
			if len(current.path) > 0 {
				if temp2.position == current.path[len(current.path)-1].grid {
					continue
				}
			}
			if err != nil {
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
