package GUI

import (
	"image/color"
	"stima/core"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var GridContainer *fyne.Container
var Slider *widget.Slider

func MakeGap(sizeX, sizeY float32) fyne.CanvasObject{
	gap := canvas.NewRectangle(color.Transparent)
	gap.SetMinSize(fyne.NewSize(sizeX, sizeY))
	return gap
}

func MakeGrid(X, Y int, g *core.Grid) fyne.CanvasObject {
	vessel := make([]fyne.CanvasObject, 0, X*Y)

	now := g
	for now != nil {
		itu := now
		for itu != nil {
			var background *canvas.Rectangle
			switch {
			case itu.GetGridType() == core.TipeBlock:
				background = canvas.NewRectangle(color.RGBA{25, 26, 165, 235})
				// fmt.Print("X")
			case itu.GetGridType() == core.TipeGoal:
				background = canvas.NewRectangle(color.RGBA{255, 243, 88, 235})
				// fmt.Print("O")
			case itu.GetGridType() == core.TipeLava:
				background = canvas.NewRectangle(color.RGBA{255, 88, 88, 235})
				// fmt.Print("L")
			case itu.GetGridType() == core.TipeStart:
				background = canvas.NewRectangle(color.RGBA{132, 88, 255, 235})
				// fmt.Print("Z")
			case itu.GetGridType() == core.TipeEmpty:
				background = canvas.NewRectangle(color.RGBA{255, 255, 255, 235})
				// if itu.Constraint != -1 {
				// 	fmt.Print(itu.Constraint)
				// } else {
				// 	fmt.Print(" ")
				// }
			}
			background.StrokeColor = color.Transparent
			background.StrokeWidth = 1
			background.SetMinSize(fyne.NewSize(72, 72))
			costlabel := canvas.NewText(strconv.Itoa(itu.Cost), color.Black)
			costlabel.Alignment = fyne.TextAlignCenter
			costlabel.TextStyle = fyne.TextStyle{Bold: true}
			costlabel.TextSize = 20
			if itu.GetGridType() == core.TipeBlock {
				costlabel.Color = color.White
			}
			cell := container.NewStack(
				background,
				costlabel,
			)
			vessel = append(vessel, cell)
			itu = itu.Kanan
		}
		now = now.Bawah
	}
	return container.NewGridWithColumns(Y, vessel...)
}

func MakeGridFromCell(index int, cells [][][]core.Cell) fyne.CanvasObject {
	var	vessel []fyne.CanvasObject

	for row := range cells[index] {
		for col := range cells[index][row] {
			var background *canvas.Rectangle
			switch {
			case cells[index][row][col].Tipe == core.TipeBlock:
				background = canvas.NewRectangle(color.RGBA{25, 26, 165, 235})
				// fmt.Print("X")
			case cells[index][row][col].Tipe == core.TipeGoal:
				background = canvas.NewRectangle(color.RGBA{255, 243, 88, 235})
				// fmt.Print("O")
			case cells[index][row][col].Tipe == core.TipeLava:
				background = canvas.NewRectangle(color.RGBA{255, 88, 88, 235})
				// fmt.Print("L")
			case cells[index][row][col].Tipe == core.TipeStart:
				background = canvas.NewRectangle(color.RGBA{132, 88, 255, 235})
				// fmt.Print("Z")
			case cells[index][row][col].Tipe == core.TipeEmpty:
				background = canvas.NewRectangle(color.RGBA{255, 255, 255, 235})
				// if itu.Constraint != -1 {
				// 	fmt.Print(itu.Constraint)
				// } else {
				// 	fmt.Print(" ")
				// }
			}
			background.StrokeColor = color.Transparent
			background.StrokeWidth = 1
			background.SetMinSize(fyne.NewSize(72, 72))
			costlabel := canvas.NewText(strconv.Itoa(cells[index][row][col].Cost), color.Black)
			costlabel.Alignment = fyne.TextAlignCenter
			costlabel.TextStyle = fyne.TextStyle{Bold: true}
			costlabel.TextSize = 20
			if cells[index][row][col].Tipe == core.TipeBlock {
				costlabel.Color = color.White
			}
			cell := container.NewStack(
				background,
				costlabel,
			)
			vessel = append(vessel, cell)
		}
	}
	return container.NewGridWithColumns(len(cells[0]), vessel...)
}

func UpdateMainPanel(X, Y int, g *core.Grid) {
	if g == nil {
		return
	}
	// newGrid := MakeGrid(X, Y, g)
	newGrid := container.NewVBox(MakeGrid(X, Y, g))
	GridContainer.Objects = []fyne.CanvasObject{newGrid}
	GridContainer.Refresh()
}

func UpdateBySlider(idx int, solution [][][]core.Cell){
	NewSol := MakeGridFromCell(idx, solution)
	newGrid := container.NewVBox(
		NewSol,
		Slider,
	)
	GridContainer.Objects = []fyne.CanvasObject{newGrid}
	GridContainer.Refresh()
}

func UpdateMainPanelSolution(solution [][][]core.Cell) {
	Slider = widget.NewSlider(0, float64(len(solution)-1))
	Slider.Step = 1

	Slider.OnChanged = func(v float64){
		idx := int(v)
		UpdateBySlider(idx, solution)
	}

	firstSol := MakeGridFromCell(0, solution)
	newGrid := container.NewVBox(
		firstSol,
		Slider,
	)
	GridContainer.Objects = []fyne.CanvasObject{newGrid}
	GridContainer.Refresh()
}