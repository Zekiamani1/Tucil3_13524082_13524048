package main

import (
	"bytes"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/widget"

	"log"
	"stima/core"
	"stima/GUI"
)

// func main() {
// 	var X int
// 	var Y int
// 	var matrix []string
// 	var costMatrix [][]int

// 	fmt.Println("Choose input option: ")
// 	fmt.Println("    1. Console")
// 	fmt.Println("    2. File input")
// 	fmt.Print(">> ")
// 	var choice int
// 	fmt.Scanln(&choice)
// 	if choice == 1 {
// 		fmt.Scan(&X, &Y)
// 		for i := 0; i < X; i++ {
// 			var temp string
// 			fmt.Scan(&temp)
// 			matrix = append(matrix, temp)
// 		}
// 		for i := 0; i < X; i++ {
// 			var temparray []int
// 			for j := 0; j < Y; j++ {
// 				var temp int
// 				fmt.Scan(&temp)
// 				temparray = append(temparray, temp)
// 			}
// 			costMatrix = append(costMatrix, temparray)
// 		}
// 	} else if choice == 2 {
// 		var filepath string
// 		fmt.Print("File path: ")
// 		fmt.Scanln(&filepath)
// 		file, err := os.Open(filepath)
// 		if err != nil {
// 			fmt.Println("File is invalid")
// 			return
// 		}
// 		defer file.Close()
// 		scanner := bufio.NewScanner(file)
// 		for scanner.Scan() {

// 		}
// 		return
// 	} else {
// 		return
// 	}
// 	// firstgrid, _, _, _ := core.CreateGrid(X, Y, matrix, costMatrix)
// 	firstgrid, start, end, _ := core.CreateGrid(X, Y, matrix, costMatrix)
// 	if firstgrid != nil {
// 		firstgrid.PrintGrid()
// 	}
// 	player := core.Player{Position: start}

// 	fmt.Println("Choose pathfinding option: ")
// 	fmt.Println("    1. UCS")
// 	fmt.Println("    2. GBFS")
// 	fmt.Println("    3. A*")
// 	fmt.Print(">> ")

// 	fmt.Scanln(&choice)

// 	var result *core.TraversalRecord
// 	if choice == 1 {
// 		result = player.UCS(end)
// 	} else if choice == 2 {
// 		result = player.ASTAR(end)
// 	} else if choice == 3 {
// 		// not implemented
// 		return
// 	} else {
// 		return
// 	}

// 	if result != nil {
// 		result.PrintResultPath(player, firstgrid)
// 	} else {
// 		fmt.Println("result null!")
// 	}
// }

type MainGrid struct {
	X int
	Y int
	firstgrid *core.Grid
	playergrid *core.Grid
	endgrid *core.Grid
	constraint []*core.Grid
}

func main() {
	// var firstgrid, startgrid, endgrid *core.Grid
	// var constraint []*core.Grid
	var peta MainGrid
	mainPanel := container.NewCenter(container.NewStack())

	fmt.Println("START")

	myApp := app.New()
	mainRunner := myApp.NewWindow("STIMMER101")

	inputPanel := GUI.NewInputPanel(&mainRunner, func (input []byte){
		X, Y, matrix, costMatrix, err := core.ParseInput(bytes.NewReader(input))
		peta.X = X
		peta.Y = Y
		if err != nil {
			dialog.ShowError(err, mainRunner)
			return
		}

		peta.firstgrid, peta.playergrid, peta.endgrid, peta.constraint, err = core.CreateGrid(X, Y, matrix, costMatrix)
		
		if err != nil {
			dialog.ShowError(err, mainRunner)
			return
		}
		GUI.UpdateMainPanel(X, Y, peta.firstgrid, mainPanel)
		log.Println(X, Y)
		log.Println(matrix)
		log.Println(costMatrix)
	})

	thewholewindow := container.NewHBox(
		GUI.MakeGap(20, 0),
		container.NewCenter(
			container.NewGridWrap(
				fyne.NewSize(250, 250),
				inputPanel.View(),
		)),
		GUI.MakeGap(20, 0),
		mainPanel,
	)

	mainRunner.SetContent(thewholewindow)
	mainRunner.Resize(fyne.NewSize(1000, 900))
	mainRunner.ShowAndRun()
}