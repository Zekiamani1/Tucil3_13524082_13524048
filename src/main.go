package main

import (
	// "bytes"
	// "bufio"
	"fmt"
	"image/color"
	// "os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "log"

	"stima/core"
	"stima/GUI"
)

func main() {
	// var firstgrid, startgrid, endgrid *core.Grid
	// var constraint []*core.Grid
	var peta core.MainGrid
	GUI.GridContainer = container.NewCenter(container.NewStack())
	
	fmt.Println("START")
	
	myApp := app.New()
	mainWindow := myApp.NewWindow("STIMMER101")

	inputPanel := GUI.NewInputPanel(&mainWindow, &peta)

	rightPanel := GUI.MakeRightPanel([]string{"GBFS", "UCS", "A*"}, &mainWindow, &peta)

	bg := canvas.NewImageFromFile("../assets/ryo_bocchi.jpg")
	bg.FillMode = canvas.ImageFillStretch

	mainLayout := container.NewHBox(
		GUI.MakeGap(20, 0),
		container.NewGridWrap(
			fyne.NewSize(250, 650),
			container.NewCenter(
				inputPanel.View(),
			),
		),
		GUI.MakeGap(20, 0),
		container.NewGridWrap(
			fyne.NewSize(600, 700),
			GUI.GridContainer,
		),
		GUI.MakeGap(20, 0),
		container.NewGridWrap(
			fyne.NewSize(250, 700),
			container.NewCenter(
				rightPanel,
			),
		),
	)

	thewholewindow := container.NewStack(
		bg,
		canvas.NewRectangle(color.RGBA{0,0,0,127}),
		mainLayout,
	)

	mainWindow.SetContent(thewholewindow)
	mainWindow.Resize(fyne.NewSize(950, 700))
	mainWindow.ShowAndRun()
}
