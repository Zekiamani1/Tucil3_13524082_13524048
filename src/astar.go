package main

import (
	"math"
	"sort"
	"slices"
)

func (h *ucshelper) calculateFCost(other *Grid) float64 {
	if h.grid == nil {
		return math.MaxFloat64
	}
	arr := h.path
	total := 0
	for _, v := range arr {
		if v.grid == nil {
			continue
		}
		total += v.grid.cost
	}
	total += h.grid.cost
	totalF := float64(total) + h.grid.calculateEuclideanDistance(other)
	return totalF
}

func (p Player) astar(end *Grid) {
	save := p
	// step := 1
	open := make([]ucshelper, 0)
	open = append(open, ucshelper{path: nil, grid: p.position})
	closed := make([]ucshelper, 0)
	current := ucshelper{}
	for len(open) > 0 {
		p.position = open[0].grid
		current = open[0]
		if p.position == end {
			current.printastarpath(save)
			return
		}
		closed = append(closed, current)
		open = open[1:]
		for _, v := range Allarah {
			temp2 := p
			err := temp2.move(v)
			if err != nil {
				continue
			} 

			newNode := ucshelper{path: append(current.path, current), grid: temp2.position, arah: v}

			openIdx := slices.IndexFunc(open, func(i ucshelper) bool {
				return i.grid == temp2.position
			})
			if openIdx != -1 {
				if open[openIdx].calculateCost() > newNode.calculateCost() {
					open = append(open[:openIdx], open[openIdx+1:]...)
				} else {
					continue
				}
			}

			closedIdx := slices.IndexFunc(closed, func(i ucshelper) bool {
				return i.grid == temp2.position
			})
			if closedIdx != -1 {
				if closed[closedIdx].calculateCost() <= newNode.calculateCost() {
					continue
				}
			}

			open = append(open, newNode)
		}
		sort.Slice(open, func(i, j int) bool {
			return open[i].calculateFCost(end) < open[j].calculateFCost(end)
		})
	}
}

func (u *ucshelper) printastarpath(player Player) {
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