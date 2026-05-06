package core

import "math"

func (p *Player) GBFS(topleft, end *Grid) {
	step := 1
	for true {
		neighbor := make([]*Grid, 0, 4)
		for _, v := range Allarah {
			temp2 := *p
			err := (&temp2).move(v)

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
		p.Position.tipe = TipeEmpty
		p.move(Allarah[choose])
		if p.Position.tipe == TipeGoal {
			println("step ", step, Allarah[choose])
			p.Position.tipe = TipeStart
			println()
			topleft.PrintGrid()
			return
		}
		p.Position.tipe = TipeStart
		println()
		topleft.PrintGrid()
		p.Position.tipe = TipeEmpty
		step++
	}
}
