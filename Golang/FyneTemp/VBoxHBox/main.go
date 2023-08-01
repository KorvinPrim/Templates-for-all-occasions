package main

/*
H and V Box
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("H and V Box")
	w.Resize(fyne.NewSize(500, 500))

	lbl1 := widget.NewLabel("Lbl1")
	lbl2 := widget.NewLabel("Lbl2")
	lbl3 := widget.NewLabel("Lbl3")
	lbl4 := widget.NewLabel("Lbl4")

	btn1 := widget.NewButton("Button1", func() {})
	btn2 := widget.NewButton("Button2", func() {})
	btn3 := widget.NewButton("Button3", func() {})
	btn4 := widget.NewButton("Button4", func() {})
	btn_box := container.NewHBox(btn1, btn2, btn3, btn4)
	label_box := container.NewHBox(lbl1, lbl2, lbl3, lbl4)

	content := container.NewVBox(label_box, btn_box)

	w.SetContent(content)

	w.ShowAndRun()
}

//func main() {
//	a := app.New()
//	w := a.NewWindow("H and V Box")
//	w.Resize(fyne.NewSize(500, 200))
//
//	lbl1 := widget.NewLabel("Lbl1")
//	lbl2 := widget.NewLabel("Lbl2")
//	lbl3 := widget.NewLabel("Lbl3")
//	lbl4 := widget.NewLabel("Lbl4")
//
//	btn1 := widget.NewButton("Button1", func() {})
//	btn2 := widget.NewButton("Button2", func() {})
//	btn3 := widget.NewButton("Button3", func() {})
//	btn4 := widget.NewButton("Button4", func() {})
//	btn_box := container.NewGridWithColumns(4, btn1, btn2, btn3, btn4)
//	label_box := container.NewGridWithColumns(4, lbl1, lbl2, lbl3, lbl4)
//
//	content := container.NewGridWithRows(2, label_box, btn_box)
//
//	w.SetContent(content)
//
//	w.ShowAndRun()
//}
//fyne package -os windows -icon myapp.png
