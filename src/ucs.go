package main

import (
	"sort"
	"slices"
)

type traversalRecord struct {
	grid *Grid
	arah Arah
	path *traversalRecord
}

func (h traversalRecord) calculateCost() int {
	parent := h.path
	total := 0
	for parent != nil {
		if parent.grid != nil {
			total += parent.grid.cost
		}
		parent = parent.path
	}
	total += h.grid.cost
	return total
}
func (p Player) ucs(end *Grid) {
	save := p
	// step := 1
	queue := make([]traversalRecord, 0)
	closed := make([]traversalRecord, 0)
	current := traversalRecord{grid: p.position}
	for true {
		for _, v := range Allarah {
			temp2 := p
			err := temp2.move(v)
			if err != nil {
				continue
			}
			
			parent := current
			newNode := traversalRecord{path: &parent, grid: temp2.position, arah: v}
			closedIdx := slices.IndexFunc(closed, func(i traversalRecord) bool {
				return i.grid == temp2.position
			})
			if closedIdx != -1 {
				if closed[closedIdx].calculateCost() <= newNode.calculateCost() {
					continue
				}
			}
			queue = append(queue, newNode)
		}
		closed = append(closed, current)
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].calculateCost() < queue[j].calculateCost()
		})
		p.position = queue[0].grid
		current = queue[0]
		if p.position == end {
			current.printucspath(save)
			return
		}
		queue = queue[1:]
	}
}

func (u traversalRecord) printucspath(player Player) {
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
