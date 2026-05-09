package core

import "math"

func (p Player) GBFS(left, end *Grid) *TraversalRecord {
	current := TraversalRecord{grid: p.Position}
	for true {
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
			dist := n.calculateEuclideanDistance(end)
			if dist < min {
				min = dist
				choose = i
			}
		}
		p.move(Allarah[choose])
		temp := current
		current = TraversalRecord{grid: p.Position, path: &temp, arah: Allarah[choose]}
		if p.Position.tipe == TipeGoal {
			return &current
		}
	}
	return nil
}
