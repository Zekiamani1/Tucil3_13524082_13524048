package core

import (
	"slices"
	"sort"
)

func (p Player) ASTAR(end *Grid, constraints []*Grid, hereustic int, NeedToDoAllConstraint bool) (int, *TraversalRecord) {
	var constraint []*Grid
	if !NeedToDoAllConstraint {
		constraint = []*Grid{end}
	} else {
		constraint = append(constraints, end)
	}
	target := 0
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
		target = slices.IndexFunc(constraint, func(i *Grid) bool {
			return i.Constraint == p.CurrentConstraint
		})
		if target == -1 {
			target = len(constraint) - 1
		}
		if p.Position == end && (!NeedToDoAllConstraint || (len(constraint) < 2 || p.CurrentConstraint > constraint[len(constraint)-2].Constraint)) {
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
				if closed[closedIdx].calculateCost(hereustic, end) <= newNode.calculateCost(hereustic, end) && !(closed[closedIdx].constraintNow < newNode.constraintNow) {
					continue
				}
			}

			open = append(open, newNode)
		}
		sort.Slice(open, func(i, j int) bool {
			return open[i].calculateCost(hereustic, constraint[target]) < open[j].calculateCost(hereustic, constraint[target])
		})
	}
	return 0, nil
}
