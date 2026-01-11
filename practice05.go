package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	titleLabel  *widget.RichText
	SecondLabel *widget.Label
	thirdLabel  *widget.Label
	fourthLabel *widget.RichText
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("test")
	third := widget.NewLabel("Bold Label")
	third.TextStyle = fyne.TextStyle{Bold: true}

	state := &AppState{
		titleLabel:  widget.NewRichTextFromMarkdown("# Markdown TITLE"),
		SecondLabel: widget.NewLabel("widget sub TITLE"),
		thirdLabel:  third,
		fourthLabel: widget.NewRichTextFromMarkdown("*markdown italic*"),
	}

	window.SetContent(state.makeUI())
	window.Resize(fyne.NewSize(500, 500))
	window.ShowAndRun()
}

func (s *AppState) makeUI() fyne.CanvasObject {
	return container.NewVBox(
		s.titleLabel,
		s.SecondLabel,
		s.thirdLabel,
		s.fourthLabel,
	)
}
