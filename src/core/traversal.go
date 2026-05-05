package core

type traversalRecord struct {
	grid *Grid
	arah Arah
	path *traversalRecord
}

func (h traversalRecord) calculateCost() int {
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


func (h *traversalRecord) calculateFCost(other *Grid) float64 {
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

func (u *traversalRecord) PrintResultPath(player Player, topleft *Grid) {
	parent := u.path
	var chosenPath []traversalRecord
	for parent != nil {
		chosenPath = append([]traversalRecord{*parent}, chosenPath...)
		parent = parent.path
	}
	for i := 0; i < len(chosenPath); i++ {
		player.Position.tipe = TipeEmpty
		player.move(chosenPath[i].arah)
		player.Position.tipe = TipeStart
		println()
		println(arahToString(chosenPath[i].arah))
		topleft.printGrid()
	}
	player.Position.tipe = TipeEmpty
	player.move(u.arah)
	player.Position.tipe = TipeStart
	println()
	println(arahToString(u.arah))
	topleft.printGrid()
}