package GUI

import (
	"fmt"
	"image/color"
	"stima/core"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var RightPanel fyne.CanvasObject
var chosenAlgo string
var chosenMode string
var Output []byte

func AlgoChooser(options []string) *widget.RadioGroup {
	return widget.NewRadioGroup(options, func(chosen string){
		chosenAlgo = chosen
	})
}

func modeChooser() *widget.RadioGroup {
	return widget.NewRadioGroup([]string{"Visit all constraint", "Go to goal following constraint"}, func(chosen string){
		chosenMode = chosen
	})
}

func SaveFileDialog(widnwo *fyne.Window){
	dialog.NewFileSave(func(IO fyne.URIWriteCloser, err error){
		if err != nil {
			dialog.ShowError(err, *widnwo)
			return
		}
		if IO == nil {
			return
		}
		defer IO.Close()

		IO.Write(Output)
	}, *widnwo).Show()
}

func MakeRightPanel(options []string, window *fyne.Window, peta *core.MainGrid) fyne.CanvasObject {
	Solution := canvas.NewText("", color.RGBA{255, 240, 89, 255})
	SolutionDetail := canvas.NewText("", color.RGBA{255, 240, 89, 255})
	Solution.Hide()
	SolutionDetail.Hide()
	SaveFile := widget.NewButton("Save File", func(){
		SaveFileDialog(window)
	})
	SaveFile.Hide()

	tombolSubmit := widget.NewButton("Submit", func() {
		if peta.Firstgrid == nil {
			dialog.ShowInformation("Error", "No map yet", *window)
			return
		}
		if chosenAlgo == "" {
			dialog.ShowInformation("Error", "No algorithm choosen", *window)
			return
		}
		if peta.Endgrid.GetGridType() == core.TipeStart {
			dialog.ShowInformation("Error", "Map is already solved", *window)
			return
		}
		if chosenMode == "" {
			dialog.ShowInformation("Error", "No mode choosen", *window)
			return
		}
		
		// BACKENDBACKENDBACKENDBACKEND
		player := core.Player{Position: peta.Playergrid}
		start := time.Now()
		mode := chosenMode == "Visit all constraint"
		iteration, pathResults := peta.RunAlgo(&player, chosenAlgo, mode)
		duration := time.Since(start)
		if pathResults == nil {
			dialog.ShowInformation("Error", "Map doesn't have solution", *window)
			return
		}
		
		Solution.Text = fmt.Sprintf("Solution: %s", pathResults.GetDirectionsAsString(true))
		Solution.TextStyle = fyne.TextStyle{Bold: true}
		Solution.TextSize = 16
		Solution.Alignment = fyne.TextAlignCenter
		Solution.Show()
		
		SolutionDetail.Text = fmt.Sprintf("Time: %s Iteration: %d", duration, iteration)
		SolutionDetail.TextStyle = fyne.TextStyle{Bold: true}
		SolutionDetail.TextSize = 16
		SolutionDetail.Alignment = fyne.TextAlignCenter
		SolutionDetail.Show()
		SaveFile.Show()

		RightPanel.Refresh()
		var pathFrames [][][]core.Cell
		Output, pathFrames = pathResults.GetResultPath(&player, peta.Firstgrid)
		AccCost = pathResults.GetAccumulatedCost()
		UpdateMainPanelSolution(pathFrames)
	})

	title := canvas.NewText("STIMMER101", color.RGBA{255, 240, 89, 255})
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 30
	selection := AlgoChooser(options)
	ModeSelection := modeChooser()

	rightPanel := container.NewVBox(
		widget.NewLabel("Choose Algorithm"),
		selection,
		MakeGap(0,1),
		widget.NewLabel("Choose Mode"),
		ModeSelection,
		MakeGap(0,1),
		tombolSubmit,
		MakeGap(0,1),
		Solution,
		SolutionDetail,
		SaveFile,
	)

	return container.NewGridWrap(fyne.NewSize(250, 250), rightPanel)
}