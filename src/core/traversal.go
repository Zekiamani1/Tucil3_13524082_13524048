package core

type TraversalRecord struct {
	constraintNow int
	grid          *Grid
	arah          Arah
	path          *TraversalRecord
}

func (h TraversalRecord) calculateCost() int {
	parent := h.path
	total := 0
	for parent != nil {
		if parent.grid != nil {
			total += parent.grid.Cost
		}
		parent = parent.path
	}
	total += h.grid.Cost
	return total
}

func (h *TraversalRecord) calculateFCost(other *Grid) float64 {
	parent := h.path
	total := 0
	for parent != nil {
		if parent.grid != nil {
			total += parent.grid.Cost
		}
		parent = parent.path
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
		println(arahToString(chosenPath[i].arah))
		topleft.PrintGrid()
	}
	player.Position.tipe = TipeEmpty
	player.move(u.arah)
	player.Position.tipe = TipeStart
	println()
	println(arahToString(u.arah))
	topleft.PrintGrid()
}
