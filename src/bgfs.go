package main

import "math"

func (p *Player) bgfs(end *Grid) {
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
			neighbor = append(neighbor, temp2.position)
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
		p.position.tipe = TipeEmpty
		p.move(Allarah[choose])
		if p.position.tipe == TipeGoal {
			println("step ", step, Allarah[choose])
			p.position.tipe = TipeStart
			println()
			peta.printGrid()
			return
		}
		p.position.tipe = TipeStart
		println()
		peta.printGrid()
		p.position.tipe = TipeEmpty
	}
}
