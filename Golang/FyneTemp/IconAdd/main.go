package main

/*
Url Img
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Ico for app")
	w.Resize(fyne.NewSize(500, 500))

	ic, _ := fyne.LoadResourceFromPath("ico.png")
	w.SetIcon(ic)

	a.Settings().SetTheme(theme.DarkTheme())

	res, _ := fyne.LoadResourceFromURLString("https://phonoteka.org/uploads/posts/2021-07/1625711842_31-phonoteka-org-p-krasivie-arti-kosmosa-krasivo-31.jpg")

	img := canvas.NewImageFromResource(res)

	w.SetContent(img)

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
