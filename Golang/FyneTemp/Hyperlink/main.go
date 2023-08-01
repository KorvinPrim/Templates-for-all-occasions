package main

/*
Demo work hyperlink.
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	url2 "net/url"
)

func main() {
	a := app.New()
	w := a.NewWindow("Демонстрация геперссылки")
	w.Resize(fyne.NewSize(300, 400))

	url, _ := url2.Parse("https://github.com/KorvinPrim")

	link := widget.NewHyperlink("Мой гитхаб", url)

	title := widget.NewLabel("Гиперссылка")

	w.SetContent(container.NewVBox(
		title,
		link,
	))

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
