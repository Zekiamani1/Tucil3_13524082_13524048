package core

import (
	"errors"
	"fmt"
	"math"
	"unicode"
	"sort"
)

type Tipe int

const (
	TipeEmpty Tipe = iota
	TipeBlock
	TipeStart
	TipeGoal
	TipeLava
)

type Grid struct {
	coordinateX int
	coordinateY int
	tipe        Tipe
	Constraint  int
	Cost        int
	Kiri        *Grid
	Kanan       *Grid
	Atas        *Grid
	Bawah       *Grid
}
type Player struct {
	Position          *Grid
	Cost              int
	CurrentConstraint int
}

type Arah int

const (
	kiri Arah = iota
	kanan
	atas
	bawah
)

var Allarah = []Arah{
	kiri,
	kanan,
	atas,
	bawah,
}

func arahToString(arah Arah) string {
	switch arah {
		case kiri:
			return "kiri"
		case kanan:
			return "kanan"
		case bawah:
			return "bawah"
		case atas:
			return "atas"
	}
	return "arah invalid"
}

func (p *Player) move(arah Arah) error { //kalo false berarti gabisa lewat situ
	if p == nil || p.Position == nil {
		return errors.New("player or Position is nil")
	}
	temp := p.Position
	canstop := false
	for temp != nil {
		stop := false
		switch arah {
		case kiri:
			if temp.Kiri.tipe == TipeBlock {
				stop = true
				break
			}
			temp = temp.Kiri
		case kanan:
			if temp.Kanan.tipe == TipeBlock {
				stop = true
				break
			}
			temp = temp.Kanan
		case atas:
			if temp.Atas.tipe == TipeBlock {
				stop = true
				break
			}
			temp = temp.Atas
		case bawah:
			if temp.Bawah.tipe == TipeBlock {
				stop = true
				break
			}
			temp = temp.Bawah
		}
		if stop {
			if !canstop {
				return errors.New("tembok le")
			}
			break
		}
		if temp == nil {
			return errors.New("cannot move: reached boundary")
		}
		// if temp.Constraint > p.currentConstraint {
		// 	return errors.New("constraint tidak terpenuhi")
		// }
		p.Position = temp
		// p.currentConstraint += p.Position.Constraint
		// temp.Constraint = -1
		p.Cost += p.Position.Cost
		canstop = true
	}
	return nil
}
func CreateGrid(X, Y int, matrix []string, costMatrix [][]int) (firstgrid *Grid, start *Grid, end *Grid, constraint []*Grid, err error) {
	grid := make([][]*Grid, X)
	err = nil
	for i := 0; i < X; i++ {
		temp := matrix[i]
		input := []rune(temp)
		if len(input) != Y {
			return nil, nil, nil, nil, errors.New("Input length invalid")
		}
		grid[i] = make([]*Grid, Y)
		for j := 0; j < Y; j++ {
			var temp2 *Grid
			switch {
			case input[j] == 'X':
				temp2 = &Grid{tipe: TipeBlock, coordinateX: j, coordinateY: i}
			case input[j] == '*':
				temp2 = &Grid{Constraint: -1, tipe: TipeEmpty, coordinateX: j, coordinateY: i}
			case unicode.IsNumber(input[j]):
				temp2 = &Grid{Constraint: int(input[j] - '0'), tipe: TipeEmpty, coordinateX: j, coordinateY: i}
				constraint = append(constraint, temp2)
			case input[j] == 'L':
				temp2 = &Grid{Constraint: -1, tipe: TipeLava, coordinateX: j, coordinateY: i}
			case input[j] == 'O':
				temp2 = &Grid{Constraint: -1, tipe: TipeGoal, coordinateX: j, coordinateY: i}
				end = temp2
			case input[j] == 'Z':
				temp2 = &Grid{Constraint: -1, tipe: TipeStart, coordinateX: j, coordinateY: i}
				start = temp2
			default:
				temp2 = &Grid{Constraint: -1, tipe: TipeStart, coordinateX: j, coordinateY: i}
			}
			grid[i][j] = temp2
		}
	}
	if start == nil {
		return nil, nil, nil, nil, errors.New("Input is invalid: missing player tile")
	}
	if end == nil {
		return nil, nil, nil, nil, errors.New("Input is invalid: missing goal tile")
	}
	sort.Slice(constraint, func(i, j int) bool {
		return constraint[i].Constraint < constraint[j].Constraint
	})
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			grid[i][j].Cost = costMatrix[i][j]
		}
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			temp2 := grid[i][j]
			if j > 0 {
				temp2.Kiri = grid[i][j-1]
			}
			if j < Y-1 {
				temp2.Kanan = grid[i][j+1]
			}
			if i > 0 {
				temp2.Atas = grid[i-1][j]
			}
			if i < X-1 {
				temp2.Bawah = grid[i+1][j]
			}
		}
	}
	firstgrid = grid[0][0]
	return
}
func (g *Grid) calculateEuclideanDistance(other *Grid) float64 {
	return math.Sqrt(math.Pow(float64(other.coordinateX)-float64(g.coordinateX), 2) + math.Pow(float64(other.coordinateY)-float64(g.coordinateY), 2))

}
func (g *Grid) PrintGrid() {
	now := g
	for now != nil {
		itu := now
		for itu != nil {
			switch {
			case itu.tipe == TipeBlock:
				fmt.Print("X")
			case itu.tipe == TipeGoal:
				fmt.Print("O")
			case itu.tipe == TipeLava:
				fmt.Print("L")
			case itu.tipe == TipeStart:
				fmt.Print("Z")
			case itu.tipe == TipeEmpty:
				if itu.Constraint != -1 {
					fmt.Print(itu.Constraint)
				} else {
					fmt.Print(" ")
				}
			}
			itu = itu.Kanan
		}
		fmt.Print("\n")
		now = now.Bawah
	}
}

func (this *Grid) GetGridType() Tipe {
	return this.tipe
}
// func main() {
// 	firstgrid, start, end, _ := CreateGrid()
// 	player := Player{Position: start}
// 	Peta = firstgrid
// 	// player.astar(end)
// 	player.ucs(end)
// 	// println()
// 	// firstgrid.printGrid()
// }
