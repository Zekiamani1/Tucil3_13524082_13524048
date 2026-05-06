package core
import (
	"sort"
	"slices"
)

func (p Player) UCS(end *Grid) *TraversalRecord {
	// step := 1
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
			newNode := TraversalRecord{path: &parent, grid: temp2.Position, arah: v}
			closedIdx := slices.IndexFunc(closed, func(i TraversalRecord) bool {
				return i.grid == temp2.Position
			})
			if closedIdx != -1 {
				if closed[closedIdx].calculateCost() <= newNode.calculateCost() {
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
		current = queue[0]
		if p.Position == end {
			return &current
		}
		queue = queue[1:]
	}
	return nil
}
