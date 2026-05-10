package core

import (
	"bytes"
	"fmt"
)

type TraversalRecord struct {
	constraintNow int
	grid          *Grid
	arah          Arah
	path          *TraversalRecord
}

func (h TraversalRecord) calculateCost(pilihan int, end *Grid) float64 {
	parent := &h
	var chosenPath []TraversalRecord
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	total := 0
	for i, x := range chosenPath {
		if x.grid != nil {
			tranvers := x.grid
			if i != len(chosenPath)-1 {
				for tranvers != nil && tranvers != chosenPath[i+1].grid {
					total += tranvers.Cost
					switch chosenPath[i+1].arah {
					case kiri:
						tranvers = tranvers.Kiri
					case kanan:
						tranvers = tranvers.Kanan
					case atas:
						tranvers = tranvers.Atas
					case bawah:
						tranvers = tranvers.Bawah
					}
				}
			}
		}
	}
	total += h.grid.Cost
	var totalF float64
	switch pilihan {
	case 0:
		totalF = float64(total) //0 hereustic
	case 1:
		totalF = float64(total) + h.grid.calculateEuclideanDistance(end) //1 euclidean hereustic
	case 2:
		totalF = float64(total) + h.grid.calculateManhattanDistance(end) //2 manhattan hereustic
	}
	return totalF
}

func (u *TraversalRecord) PrintResultPath(player Player, topleft *Grid) {
	parent := u.path
	var chosenPath []TraversalRecord
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		player.Position.tipe = TipeEmpty
		player.move(chosenPath[i].arah)
		player.Position.tipe = TipeStart
		println()
		println("arah: ", arahToString(false, chosenPath[i].arah))
		println("Cost saat ini: ", chosenPath[i].calculateCost(0, nil))
		topleft.PrintGrid()
	}
	player.Position.tipe = TipeEmpty
	player.move(u.arah)
	player.Position.tipe = TipeStart
	println()
	println("arah: ", arahToString(false, u.arah))
	println("Cost saat ini: ", u.calculateCost(0, nil))
	topleft.PrintGrid()
}

func (this *TraversalRecord) GetResultPath(player *Player, topleft *Grid) ([]byte, [][][]Cell) {
	parent := this.path
	var chosenPath []TraversalRecord
	var result [][][]Cell
	var output bytes.Buffer
	ConstStep := this.GetConstraint()
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		player.Position.tipe = TipeEmpty
		player.move(chosenPath[i].arah)
		player.Position.tipe = TipeStart
		result = append(result, topleft.ToCells())
		fmt.Fprintln(&output, "")
		fmt.Fprintln(&output, "Arah: ", arahToString(false, chosenPath[i].arah))
		fmt.Fprintln(&output, "Cost saat ini: ", chosenPath[i].calculateCost(0, nil))
		topleft.ToBytes(&output, ConstStep[i])
	}
	player.Position.tipe = TipeEmpty
	player.move(this.arah)
	player.Position.tipe = TipeStart
	result = append(result, topleft.ToCells())
	fmt.Fprintln(&output, "")
	fmt.Fprintln(&output, "Arah: ", arahToString(false, this.arah))
	fmt.Fprintln(&output, "Cost saat ini: ", this.calculateCost(0, nil))
	topleft.ToBytes(&output, this.constraintNow)
	return output.Bytes(), result
}

func (this *TraversalRecord) GetDirectionsAsString(simplified bool) string {
	parent := this.path
	var chosenPath []TraversalRecord
	var result string
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 1; i < len(chosenPath); i++ {
		result = result + arahToString(true, chosenPath[i].arah)
	}
	result = result + arahToString(true, this.arah)
	return result
}

func (this *TraversalRecord) GetAccumulatedCost() []int {
	parent := this.path
	var chosenPath []TraversalRecord
	var cost []int
	total := 0
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		total = int(chosenPath[i].calculateCost(0, nil))
		cost = append(cost, total)
	}
	total = int(this.calculateCost(0, nil))
	cost = append(cost, total)
	return cost
}

func (this *TraversalRecord) GetConstraint() []int {
	parent := this.path
	var chosenPath []TraversalRecord
	var cost []int
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		cost = append(cost, chosenPath[i].constraintNow)
	}
	cost = append(cost, this.constraintNow)
	return cost
}