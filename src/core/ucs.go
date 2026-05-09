package core

import (
	"slices"
	"sort"
)

func (p Player) UCS(end *Grid, constraint []*Grid) *TraversalRecord {
	queue := make([]TraversalRecord, 0)
	closed := make([]TraversalRecord, 0)
	current := TraversalRecord{grid: p.Position}
	for true {
		for _, v := range Allarah {
			temp2 := p
			err := temp2.move(v)
			if err != nil {
				continue
			}

			parent := current
			newNode := TraversalRecord{path: &parent, grid: temp2.Position, arah: v, constraintNow: temp2.CurrentConstraint}
			closedIdx := slices.IndexFunc(closed, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
			})
			if closedIdx != -1 {
				if closed[closedIdx].calculateCost() <= newNode.calculateCost() && !(closed[closedIdx].constraintNow < newNode.constraintNow) {
					continue
				}
			}
			queue = append(queue, newNode)
		}
		closed = append(closed, current)
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].calculateCost() < queue[j].calculateCost()
		})

		if len(queue) == 0 {
			break
		}

		p.Position = queue[0].grid
		p.CurrentConstraint = queue[0].constraintNow
		current = queue[0]
		if p.Position == end && p.CurrentConstraint > constraint[len(constraint)-1].Constraint {
			return &current
		}
		queue = queue[1:]
	}
	return nil
}
