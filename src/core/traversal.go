package core

type TraversalRecord struct {
	constraintNow int
	grid          *Grid
	arah          Arah
	path          *TraversalRecord
}

func (h TraversalRecord) calculateCost() int {
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
	return total
}

func (h TraversalRecord) calculateFCost(other *Grid) float64 {
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
	totalF := float64(total) + h.grid.calculateEuclideanDistance(other)
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
		println("arah: ", arahToString(chosenPath[i].arah))
		println("Cost saat ini: ", chosenPath[i].calculateCost())
		topleft.PrintGrid()
	}
	player.Position.tipe = TipeEmpty
	player.move(u.arah)
	player.Position.tipe = TipeStart
	println()
	println("arah: ", arahToString(u.arah))
	println("Cost saat ini: ", u.calculateCost())
	topleft.PrintGrid()
}

func (this *TraversalRecord) ToCells(player *Player, topleft *Grid) [][][]Cell {
	parent := this.path
	var chosenPath []TraversalRecord
	var result [][][]Cell
	for parent != nil {
		chosenPath = append([]TraversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		player.Position.tipe = TipeEmpty
		player.move(chosenPath[i].arah)
		player.Position.tipe = TipeStart
		result = append(result, topleft.ToCells())
	}
	player.Position.tipe = TipeEmpty
	player.move(this.arah)
	player.Position.tipe = TipeStart
	result = append(result, topleft.ToCells())
	return result
}