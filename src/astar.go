package main

import (
	"sort"
	"slices"
)

func (h *traversalRecord) calculateFCost(other *Grid) float64 {
	parent := h.path
	total := 0
	for parent != nil {
		if parent.grid != nil {
			total += parent.grid.cost
		}
		parent = parent.path
	}
	total += h.grid.cost
	totalF := float64(total) + h.grid.calculateEuclideanDistance(other)
	return totalF
}

func (p Player) astar(end *Grid) {
	save := p
	// step := 1
	open := make([]traversalRecord, 0)
	open = append(open, traversalRecord{path: nil, grid: p.position})
	closed := make([]traversalRecord, 0)
	current := traversalRecord{}
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

			parent := current
			newNode := traversalRecord{path: &parent, grid: temp2.position, arah: v}

			openIdx := slices.IndexFunc(open, func(i traversalRecord) bool {
				return i.grid == temp2.position
			})
			if openIdx != -1 {
				if open[openIdx].calculateCost() > newNode.calculateCost() {
					open = append(open[:openIdx], open[openIdx+1:]...)
				} else {
					continue
				}
			}

			closedIdx := slices.IndexFunc(closed, func(i traversalRecord) bool {
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

func (u *traversalRecord) printastarpath(player Player) {
	parent := u.path
	var chosenPath []traversalRecord
	for parent != nil {
		chosenPath = append([]traversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		player.position.tipe = TipeEmpty
		player.move(chosenPath[i].arah)
		player.position.tipe = TipeStart
		println()
		println(arahToString(chosenPath[i].arah))
		peta.printGrid()
	}
	player.position.tipe = TipeEmpty
	player.move(u.arah)
	player.position.tipe = TipeStart
	println()
	println(arahToString(u.arah))
	peta.printGrid()
}