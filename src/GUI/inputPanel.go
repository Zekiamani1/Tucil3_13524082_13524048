package GUI

import (
	"bytes"
	"image/color"
	"io"
	"path/filepath"
	"stima/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type InputPanel struct {
	textInput *widget.Entry
	fileLabel *widget.Label
	fileContent []byte
	
	window *fyne.Window
	peta *core.MainGrid
}

func (this *InputPanel) submitFunc(input []byte){
	X, Y, matrix, costMatrix, err := core.ParseInput(bytes.NewReader(input))
	this.peta.X = X
	this.peta.Y = Y
	if err != nil {
		dialog.ShowError(err, *this.window)
		return
	}

	this.peta.Firstgrid, this.peta.Playergrid, this.peta.Endgrid, this.peta.Constraint, err = core.CreateGrid(X, Y, matrix, costMatrix)
	
	if err != nil {
		dialog.ShowError(err, *this.window)
		return
	}
	UpdateMainPanel(X, Y, this.peta.Firstgrid)
}

func NewInputPanel(w *fyne.Window, peta *core.MainGrid) *InputPanel {
	textInput := widget.NewMultiLineEntry()
	textInput.SetPlaceHolder("Masukkan board")
	textInput.SetMinRowsVisible(8)
	return &InputPanel{
		textInput: textInput,
		fileLabel: widget.NewLabel("No file selected"),
		window: w,
		peta: peta,
	}
}

func (this *InputPanel) selectFile() {
	fileOpener := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, *this.window)
			return
		}
		if reader == nil {
			return
		}
		defer reader.Close()

		data, err := io.ReadAll(reader)
		if err != nil {
			dialog.ShowError(err, *this.window)
			return
		}

		this.fileContent = data
		this.fileLabel.SetText(reader.URI().Name())
	}, *this.window)

	path, err1 := filepath.Abs(".")
	if err1 != nil {
		dialog.ShowError(err1, *this.window)
		return
	}
	loc, err2 := storage.ListerForURI(storage.NewFileURI(path))
	if err2 != nil {
		dialog.ShowError(err2, *this.window)
		return
	}
	fileOpener.SetLocation(loc)
	fileOpener.Show()
}

func (this *InputPanel) submit() {
	var input []byte
	if this.textInput.Text != "" {
		input = []byte(this.textInput.Text)
	} else if len(this.fileContent) > 0 {
		input = []byte(this.fileContent)
	} else {
		dialog.ShowInformation("Error", "No input provided", *this.window)
		return
	}
	this.submitFunc(input)
}

func (this *InputPanel) View() fyne.CanvasObject{
	bukaFile := widget.NewButton("Choose File", func() {
		this.selectFile()
	})

	tombolSubmit := widget.NewButton("Submit", func() {
		this.submit()
	})

	title := canvas.NewText("STIMMER101", color.RGBA{255, 240, 89, 255})
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 30

	input := container.NewVBox(
		title,
		widget.NewLabel("Multiline Input:"),
		this.textInput,
		MakeGap(0,1),
		widget.NewLabel("Or load from file:"),
		MakeGap(0,1),
		bukaFile,
		MakeGap(0,1),
		this.fileLabel,
		MakeGap(0,1),
		tombolSubmit,
	)

	return container.NewGridWrap(fyne.NewSize(250, 250), input)
}