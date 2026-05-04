package main

import (
	"errors"
	"fmt"
	"unicode"
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

func (p *Player) move(arah Arah) error { //kalo false berarti gabisa lewat situ
	if p == nil || p.position == nil {
		return errors.New("player or position is nil")
	}
	temp := p.position
	for temp != nil && temp.tipe != TipeBlock {
		switch arah {
		case kiri:
			temp = temp.Kiri
		case kanan:
			temp = temp.Kanan
		case atas:
			temp = temp.Atas
		case bawah:
			temp = temp.Bawah
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
func createGrid() *Grid {
	var X int
	var Y int
	var start *Grid
	fmt.Scan(&X, &Y)
	grid := make([][]*Grid, X)
	for i := 0; i < X; i++ {
		var temp string
		fmt.Scanln(&temp)
		input := []rune(temp)
		if len(input) != Y {
			return nil
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
			case input[j] == 'Z':
				temp2 = &Grid{Constraint: 0, tipe: TipeStart, coordinateX: j, coordinateY: i}
				// start = temp2
			}
			grid[i][j] = temp2
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
	start = grid[0][0]
	return start
}
func (g *Grid) printGrid() {
	now := g
	for now != nil {
		itu := now
		for itu != nil {
			switch {
			case itu.tipe == TipeBlock:
				fmt.Print("X")
			case itu.tipe == TipeLava:
				fmt.Print("L")
			case itu.tipe == TipeStart:
				fmt.Print("S")
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
	fmt.Println()
	start := createGrid()
	start.printGrid()

}
