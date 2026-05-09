package core

import (
	"math"
	"slices"
)

func (p Player) GBFS(end *Grid, constraints []*Grid) *TraversalRecord {
	target := 0
	constraint := append(constraints, end)
	current := TraversalRecord{grid: p.Position}
	for true {
		if current.path != nil {
			if current.path.path != nil {
				if current.path.path.path != nil {
					if current.path.path.path.grid == current.path.grid && current.grid == current.path.path.grid {
						return nil //terjadi stuck jir
					}
				}
			}
		}
		neighbor := make([]*Grid, 0, 4)
		for _, v := range Allarah {
			temp2 := p
			err := temp2.move(v)
			if err != nil {
				neighbor = append(neighbor, nil)
				continue
			}
			neighbor = append(neighbor, temp2.Position)
		}
		var choose int
		min := math.MaxFloat64
		for i, n := range neighbor {
			if n == nil {
				continue
			}
			dist := n.calculateEuclideanDistance(constraint[target])
			if dist <= min {
				min = dist
				choose = i
			}
		}
		p.move(Allarah[choose])
		temp := current
		current = TraversalRecord{grid: p.Position, path: &temp, arah: Allarah[choose]}
		target = slices.IndexFunc(constraint, func(i *Grid) bool {
			return i.Constraint == p.CurrentConstraint
		})
		if target == -1 {
			target = len(constraint) - 1
		}
		if p.Position.tipe == TipeGoal && p.CurrentConstraint > constraint[len(constraint)-2].Constraint {
			return &current
		}
	}
	return nil
}
