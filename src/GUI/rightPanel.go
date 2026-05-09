package GUI

import (
	"image/color"
	"stima/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var chosenAlgo string

func AlgoChooser(options []string) *widget.RadioGroup {
	return widget.NewRadioGroup(options, func(chosen string){
		chosenAlgo = chosen
	})
}

func MakeRightPanel(options []string, window *fyne.Window, peta *core.MainGrid) fyne.CanvasObject {
	tombolSubmit := widget.NewButton("Submit", func() {
		if peta == nil {
			dialog.ShowInformation("Error", "No map yet", *window)
			return
		}
		if chosenAlgo == "" {
			dialog.ShowInformation("Error", "No input provided", *window)
			return
		}
		if peta.Endgrid.GetGridType() == core.TipeStart {
			dialog.ShowInformation("Error", "Map is already solved", *window)
			return
		}
		
		// BACKENDBACKENDBACKENDBACKEND
		player := core.Player{Position: peta.Playergrid}
		pathResults := peta.RunAlgo(&player, chosenAlgo)
		if pathResults == nil {
			dialog.ShowInformation("Error", "Map doesn't have solution", *window)
			return
		}
		pathFrames := pathResults.ToCells(&player, peta.Firstgrid)
		UpdateMainPanelSolution(pathFrames)
	})

	title := canvas.NewText("STIMMER101", color.RGBA{255, 240, 89, 255})
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 30
	selection := AlgoChooser(options)

	rightPanel := container.NewVBox(
		widget.NewLabel("Choose Algorithm"),
		selection,

		widget.NewSeparator(),

		tombolSubmit,
	)

	return container.NewGridWrap(fyne.NewSize(250, 250), rightPanel)
}