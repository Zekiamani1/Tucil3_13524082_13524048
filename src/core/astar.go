package core

import (
	"sort"
	"slices"
)

func (p Player) ASTAR(end *Grid) *TraversalRecord {
	// step := 1
	open := make([]TraversalRecord, 0)
	open = append(open, TraversalRecord{path: nil, grid: p.Position})
	closed := make([]TraversalRecord, 0)
	current := TraversalRecord{}
	for len(open) > 0 {
		p.Position = open[0].grid
		current = open[0]
		if p.Position == end {
			return &current
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
			newNode := TraversalRecord{path: &parent, grid: temp2.Position, arah: v}

			openIdx := slices.IndexFunc(open, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
			})
			if openIdx != -1 {
				if open[openIdx].calculateCost() > newNode.calculateCost() {
					open = append(open[:openIdx], open[openIdx+1:]...)
				} else {
					continue
				}
			}

			closedIdx := slices.IndexFunc(closed, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
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
	return nil
}