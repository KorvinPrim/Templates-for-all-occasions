package main

/*
Component card
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(500, 500))
	a.Settings().SetTheme(theme.DarkTheme())

	card1 := widget.NewCard(
		"Angel",
		"photo angel",
		canvas.NewImageFromFile("angel.png"),
	)

	card2 := widget.NewCard(
		"Channel",
		"Info",
		widget.NewButton("Sub", func() {

		}),
	)

	box := container.NewGridWithColumns(2, card1, card2)
	w.SetContent(box)

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
