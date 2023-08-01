package main

/*
Adding Images
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"path"
)

func main() {
	a := app.New()
	w := a.NewWindow("Демонстрация изображения")
	w.Resize(fyne.NewSize(900, 400))

	img1 := canvas.NewImageFromFile(path.Join("img1.jpg"))

	label := widget.NewLabel("Картинка")

	w.SetContent(container.NewGridWithColumns(2, img1, label))

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
