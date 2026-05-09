package main

import (
	// "bytes"
	"bufio"
	"fmt"
	"os"

	// "fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/widget"
	// "log"

	"stima/core"
)

func main() {
	var X int
	var Y int
	var matrix []string
	var costMatrix [][]int

	fmt.Println("Choose input option: ")
	fmt.Println("    1. Console")
	fmt.Println("    2. File input")
	fmt.Print(">> ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Scan(&X, &Y)
		for i := 0; i < X; i++ {
			var temp string
			fmt.Scan(&temp)
			matrix = append(matrix, temp)
		}
		for i := 0; i < X; i++ {
			var temparray []int
			for j := 0; j < Y; j++ {
				var temp int
				fmt.Scan(&temp)
				temparray = append(temparray, temp)
			}
			costMatrix = append(costMatrix, temparray)
		}
	} else if choice == 2 {
		var filepath string
		fmt.Print("File path: ")
		fmt.Scanln(&filepath)
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("File is invalid")
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

		}
		return
	} else {
		return
	}
	// firstgrid, _, _, _ := core.CreateGrid(X, Y, matrix, costMatrix)
	firstgrid, start, end, _, _ := core.CreateGrid(X, Y, matrix, costMatrix)
	if firstgrid != nil {
		firstgrid.PrintGrid()
	}
	player := core.Player{Position: start}

	fmt.Println("Choose pathfinding option: ")
	fmt.Println("    1. UCS")
	fmt.Println("    2. GBFS")
	fmt.Println("    3. A*")
	fmt.Print(">> ")

	fmt.Scanln(&choice)

	var result *core.TraversalRecord
	if choice == 1 {
		result = player.UCS(end)
	} else if choice == 2 {
		result = player.GBFS(firstgrid, end)
	} else if choice == 3 {
		result = player.ASTAR(end)
		return
	} else {
		return
	}

	if result != nil {
		result.PrintResultPath(player, firstgrid)
	} else {
		fmt.Println("result null!")
	}
}

// func main() {
// 	// var firstgrid, startgrid, endgrid *core.Grid
// 	// var constraint []*core.Grid
// 	var peta core.MainGrid
// 	mainPanel := container.NewCenter(container.NewStack())

// 	fmt.Println("START")

// 	myApp := app.New()
// 	mainWindow := myApp.NewWindow("STIMMER101")

// 	inputPanel := GUI.NewInputPanel(&mainWindow, &peta, mainPanel)

// 	rightPanel := GUI.MakeRightPanel([]string{"GBFS", "UCS", "A*"}, &mainWindow, &peta)

// 	bg := canvas.NewImageFromFile("../assets/ryo_bocchi.jpg")
// 	bg.FillMode = canvas.ImageFillStretch

// 	mainLayout := container.NewHBox(
// 		GUI.MakeGap(20, 0),
// 		container.NewGridWrap(
// 			fyne.NewSize(250, 650),
// 			container.NewCenter(
// 				inputPanel.View(),
// 			),
// 		),
// 		GUI.MakeGap(20, 0),
// 		container.NewGridWrap(
// 			fyne.NewSize(600, 700),
// 			mainPanel,
// 		),
// 		GUI.MakeGap(20, 0),
// 		container.NewGridWrap(
// 			fyne.NewSize(250, 700),
// 			container.NewCenter(
// 				rightPanel,
// 			),
// 		),
// 	)

// 	thewholewindow := container.NewStack(
// 		bg,
// 		canvas.NewRectangle(color.RGBA{0,0,0,127}),
// 		mainLayout,
// 	)

// 	mainWindow.SetContent(thewholewindow)
// 	mainWindow.Resize(fyne.NewSize(950, 700))
// 	mainWindow.ShowAndRun()
// }
