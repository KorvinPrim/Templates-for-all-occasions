package main

/*
Adding Images
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Демонстрация изображения")
	w.Resize(fyne.NewSize(900, 400))
	
	color_for := color.NRGBA{R: 127, G: 12, B: 42, A: 255}
	label := canvas.NewText("Цветное", color_for)

	w.SetContent(container.NewGridWithColumns(1, label))

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
