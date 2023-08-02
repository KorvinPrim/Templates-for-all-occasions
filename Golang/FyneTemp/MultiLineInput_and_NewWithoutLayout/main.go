package main

/*
MultilineEntry is a widget in the Go language that allows you to create and work with multiline full fields in the Fyne library.

With MultilineEntry, you can create text fields that can contain multiple lines of text. This widget provides the ability to enter,
edit and display long texts or indented texts.

It is important to note that MultilineEntry supports multiple input options. It can be used to enter one or more paragraphs of text,
and also supports copy, paste and undo functions.

MultilineEntry also offers custom settings, including the ability to set the maximum number of rows and columns in the field,
as well as set the maximum number of characters that can be contained in the field. You can also customize the style and font s
ize of MultilineEntry using the available methods.

When you create or work with MultilineEntry, you can use it together with other Fyne widgets to create applications with r
ich text input capabilities.
*/

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(500, 500))
	a.Settings().SetTheme(theme.DarkTheme())
	//text := "textttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt"
	//entry := widget.NewMultiLineEntry()
	////entry.SetPlaceHolder("This Place Holder, Enter here ...")
	//entry.SetText(text)
	////entry.Wrapping = fyne.TextWrapOff
	//entry.Wrapping = fyne.TextWrapBreak

	lebel := widget.NewLabel("Блокнот")
	lebel.Resize(fyne.NewSize(100, 30))
	lebel.Move(fyne.NewPos(200, 40))

	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Вводите здесь")
	entry.Resize(fyne.NewSize(400, 300))
	entry.Move(fyne.NewPos(40, 100))

	btn := widget.NewButton("Сохранить", func() {
		file, err := os.Create("info.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		file.WriteString(entry.Text)
	})
	btn.Resize(fyne.NewSize(100, 30))
	btn.Move(fyne.NewPos(200, 400))

	box := container.NewWithoutLayout(lebel, entry, btn)
	w.SetContent(box)

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
