package main

/*
Chenge Theme
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("H and V Box")
	w.Resize(fyne.NewSize(500, 500))

	a.Settings().SetTheme(theme.DarkTheme())

	lbl1 := widget.NewLabel("Lbl1")
	lbl2 := widget.NewLabel("Lbl2")
	lbl3 := widget.NewLabel("Lbl3")
	lbl4 := widget.NewLabel("Lbl4")

	btn1 := widget.NewButton("Button1", func() {})
	btn2 := widget.NewButton("Button2", func() {})
	btn3 := widget.NewButton("менить тему на тёмную", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	btn4 := widget.NewButton("Сменить тему на светлую", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	btn_box := container.NewHBox(btn1, btn2, btn3, btn4)
	label_box := container.NewHBox(lbl1, lbl2, lbl3, lbl4)

	content := container.NewVBox(label_box, btn_box)

	w.SetContent(content)

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
