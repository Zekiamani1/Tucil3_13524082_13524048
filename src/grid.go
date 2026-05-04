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

func (p Player) move(arah Arah) error { //kalo false berarti gabisa lewat situ
	temp := p.position
	for temp.tipe != TipeBlock {
		switch arah {
		case kiri:
			temp = p.position.Kiri
		case kanan:
			temp = p.position.Kanan
		case atas:
			temp = p.position.Atas
		case bawah:
			temp = p.position.Bawah
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
	var now *Grid
	now = nil
	for i := 0; i < X+2; i++ {
		itu := now
		var temp string
		fmt.Scanln(&temp)
		input := []rune(temp)
		if len(input) != Y+2 {
			return nil //salah
		}
		for j := 0; j < Y+2; j++ {
			var temp2 *Grid
			switch {
			case input[j] == 'X':
				temp2 = &Grid{Kiri: itu, tipe: TipeBlock}
			case input[j] == '*':
				temp2 = &Grid{Kiri: itu, Constraint: 0, tipe: TipeEmpty, coordinateX: j, coordinateY: i}
			case unicode.IsNumber(input[j]):
				temp2 = &Grid{Kiri: itu, Constraint: int(input[i] - '0'), tipe: TipeEmpty, coordinateX: j, coordinateY: i}
			case input[j] == 'L':
				temp2 = &Grid{Kiri: itu, Constraint: 0, tipe: TipeLava, coordinateX: j, coordinateY: i}
			case input[j] == 'O':
				temp2 = &Grid{Kiri: itu, Constraint: 0, tipe: TipeGoal, coordinateX: j, coordinateY: i}
			case input[j] == 'Z':
				temp2 = &Grid{Kiri: itu, Constraint: 0, tipe: TipeStart, coordinateX: j, coordinateY: i}
				start = temp2
			}
			if itu != nil {
				itu.Kanan = temp2
				itu = itu.Kanan
			} else {
				itu = temp2
				now = itu
			}

		}
		now.Bawah.Atas = now
		now = now.Bawah
	}
	return start
}
