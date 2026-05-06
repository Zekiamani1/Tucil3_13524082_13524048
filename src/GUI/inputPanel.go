package GUI

import (
	"io"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/container"
)

type InputPanel struct {
	textInput *widget.Entry
	fileLabel *widget.Label
	fileContent []byte
	
	window *fyne.Window
	submitFunc func([]byte)
}

func NewInputPanel(w *fyne.Window, submitFunc func([]byte)) *InputPanel {
	return &InputPanel{
		textInput: widget.NewMultiLineEntry(),
		fileLabel: widget.NewLabel("No file selected"),
		submitFunc: submitFunc,
		window: w,
	}
}

func (this *InputPanel) SelectFile() {
	dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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
	}, *this.window).Show()
}

func (this *InputPanel) Submit() {
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
	if this.textInput.PlaceHolder == "" {
		this.textInput.SetPlaceHolder("Masukkan board")
	}
	bukaFile := widget.NewButton("Choose File", func() {
		this.SelectFile()
	})

	tombolSubmit := widget.NewButton("Submit", func() {
		this.Submit()
	})

	input := container.NewVBox(
		widget.NewLabel("Multiline Input:"),
		this.textInput,

		widget.NewSeparator(),

		widget.NewLabel("Or load from file:"),
		bukaFile,
		this.fileLabel,

		widget.NewSeparator(),

		tombolSubmit,
	)

	return container.NewGridWrap(fyne.NewSize(250, 250), input)
}