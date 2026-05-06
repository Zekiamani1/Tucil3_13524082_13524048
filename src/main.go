package main

import (
	"bytes"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	// 	"os"
	"bufio"
	"io"
	"log"
	// "stima/core"
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

func parseInput(r io.Reader) (int, int, []string, [][]int, error) {
	reader := bufio.NewReader(r)

	var X, Y int
	if _, err := fmt.Fscan(reader, &X, &Y); err != nil {
		return 0, 0, nil, nil, err
	}

	matrix := make([]string, X)
	for i := 0; i < X; i++ {
		if _, err := fmt.Fscan(reader, &matrix[i]); err != nil {
			return 0, 0, nil, nil, err
		}
	}

	costMatrix := make([][]int, X)
	for i := 0; i < X; i++ {
		costMatrix[i] = make([]int, Y)
		for j := 0; j < Y; j++ {
			if _, err := fmt.Fscan(reader, &costMatrix[i][j]); err != nil {
				return 0, 0, nil, nil, err
			}
		}
	}

	return X, Y, matrix, costMatrix, nil
}

func main() {
	// var firstgrid, startgrid, endgrid *core.Grid
	// var constraint []*core.Grid

	fmt.Println("START")

	myApp := app.New()
	mainRunner := myApp.NewWindow("STIMMER101")

	inputPanel := GUI.NewInputPanel(&mainRunner, func (input []byte){
		X, Y, matrix, costMatrix, err := parseInput(bytes.NewReader(input))
		if err != nil {
			dialog.ShowError(err, mainRunner)
			return
		}

		// 	firstgrid, startgrid, endgrid, _ := core.CreateGrid(X, Y, matrix, costMatrix)
		
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
		container.NewCenter(widget.NewLabel("Input here")),
	)

	mainRunner.SetContent(thewholewindow)
	mainRunner.Resize(fyne.NewSize(1000, 900))
	mainRunner.ShowAndRun()
}