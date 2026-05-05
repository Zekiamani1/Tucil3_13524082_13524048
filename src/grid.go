package main

import (
	"errors"
	"fmt"
	"math"
	"unicode"
)

var peta *Grid

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
	cost        int
	Kiri        *Grid
	Kanan       *Grid
	Atas        *Grid
	Bawah       *Grid
}
type Player struct {
	position          *Grid
	cost              int
	currentConstraint int
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

func (p *Player) move(arah Arah) error { //kalo false berarti gabisa lewat situ
	if p == nil || p.position == nil {
		return errors.New("player or position is nil")
	}
	temp := p.position
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
			break
		}
		if temp == nil {
			return errors.New("cannot move: reached boundary")
		}
		if temp.Constraint > p.currentConstraint {
			return errors.New("constraint tidak terpenuhi")
		}
		p.position = temp
		p.currentConstraint += p.position.Constraint
		temp.Constraint = 0
		p.cost += p.position.cost
	}
	return nil
}
func createGrid() (firstgrid *Grid, start *Grid, end *Grid) {
	var X int
	var Y int
	fmt.Scan(&X, &Y)
	grid := make([][]*Grid, X)
	for i := 0; i < X; i++ {
		var temp string
		fmt.Scanln(&temp)
		input := []rune(temp)
		if len(input) != Y {
			return nil, nil, nil
		}
		grid[i] = make([]*Grid, Y)
		for j := 0; j < Y; j++ {
			var temp2 *Grid
			switch {
			case input[j] == 'X':
				temp2 = &Grid{tipe: TipeBlock}
			case input[j] == '*':
				temp2 = &Grid{Constraint: 0, tipe: TipeEmpty, coordinateX: j, coordinateY: i}
			case unicode.IsNumber(input[j]):
				temp2 = &Grid{Constraint: int(input[i] - '0'), tipe: TipeEmpty, coordinateX: j, coordinateY: i}
			case input[j] == 'L':
				temp2 = &Grid{Constraint: 0, tipe: TipeLava, coordinateX: j, coordinateY: i}
			case input[j] == 'O':
				temp2 = &Grid{Constraint: 0, tipe: TipeGoal, coordinateX: j, coordinateY: i}
				end = temp2
			case input[j] == 'Z':
				temp2 = &Grid{Constraint: 0, tipe: TipeStart, coordinateX: j, coordinateY: i}
				start = temp2
			}
			grid[i][j] = temp2
		}
	}
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			var temp int
			fmt.Scan(&temp)
			grid[i][j].cost = temp
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
func (g *Grid) printGrid() {
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
				fmt.Print(itu.Constraint)
			}
			itu = itu.Kanan
		}
		fmt.Print("\n")
		now = now.Bawah
	}
}
func main() {
	firstgrid, start, _ := createGrid()
	player := Player{position: start}
	peta = firstgrid
	player.ucs()
	// println()
	// firstgrid.printGrid()

}
