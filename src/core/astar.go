package core

import (
	"slices"
	"sort"
)

func (p Player) ASTAR(end *Grid, hereustic int) (int, *TraversalRecord) {
	// step := 1
	open := make([]TraversalRecord, 0)
	open = append(open, TraversalRecord{path: nil, grid: p.Position})
	closed := make([]TraversalRecord, 0)
	current := TraversalRecord{}
	iteration := 0
	for len(open) > 0 {
		iteration += 1
		p.Position = open[0].grid
		p.CurrentConstraint = open[0].constraintNow
		current = open[0]
		if p.Position == end {
			return iteration, &current
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
			newNode := TraversalRecord{path: &parent, grid: temp2.Position, arah: v, constraintNow: temp2.CurrentConstraint}

			openIdx := slices.IndexFunc(open, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
			})
			if openIdx != -1 {
				if open[openIdx].calculateCost(hereustic, end) > newNode.calculateCost(hereustic, end) {
					open = append(open[:openIdx], open[openIdx+1:]...)
				} else {
					continue
				}
			}

			closedIdx := slices.IndexFunc(closed, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
			})
			if closedIdx != -1 {
				if closed[closedIdx].calculateCost(hereustic, end) <= newNode.calculateCost(hereustic, end) {
					continue
				}
			}

			open = append(open, newNode)
		}
		sort.Slice(open, func(i, j int) bool {
			return open[i].calculateCost(hereustic, end) < open[j].calculateCost(hereustic, end)
		})
	}
	return 0, nil
}
