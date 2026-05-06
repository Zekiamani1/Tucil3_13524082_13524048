package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func MakeGap(sizeX, sizeY float32) fyne.CanvasObject{
	gap := canvas.NewRectangle(color.Transparent)
	gap.SetMinSize(fyne.NewSize(sizeX, sizeY))
	return gap
}