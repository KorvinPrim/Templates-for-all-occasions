package main

/*
The code implementation focuses on creating a versatile menu system
in Go using the Fine library.

It enables the creation of a main menu for your application,
along with child menus and buttons. By leveraging the Go language
and Fine library, developers com seamlessly structure and organize
menus, define menu item actions, and handle user interactions.

This code implementation empowers developers to build a dynamic
and user-friendly menu system for their Go applications.
*/

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main")
	w.Resize(fyne.NewSize(500, 370))
	a.Settings().SetTheme(theme.DarkTheme())

	file_item1 := fyne.NewMenuItem("New File", func() {
		file, _ := os.Create("create.txt")
		defer file.Close()

	})
	file_item2 := fyne.NewMenuItem("Save", func() {
		fmt.Println("File Saved!")
	})

	menu1 := fyne.NewMenu("File", file_item1, file_item2)

	actions_item1 := fyne.NewMenuItem("Hello", func() { fmt.Println("Hello from menu!") })
	actions_item2 := fyne.NewMenuItem("Bye", func() { fmt.Println("Bye from menu!") })
	actions_item3 := fyne.NewMenuItem("Button", func() { fmt.Println("Button from menu!") })
	menu2 := fyne.NewMenu("Actions", actions_item1, actions_item2, actions_item3)

	info_item1 := fyne.NewMenuItem("Hello", func() { fmt.Println("Hello from menu!") })
	info_item2 := fyne.NewMenuItem("Bye", func() { fmt.Println("Bye from menu!") })
	info_item3 := fyne.NewMenuItem("Button", func() { fmt.Println("Button from menu!") })
	info_item4 := fyne.NewMenuItem("Button", func() { fmt.Println("Button from menu!") })
	info_item5 := fyne.NewMenuItem("Button", func() { fmt.Println("Button from menu!") })
	info_item6 := fyne.NewMenuItem("Button", func() { fmt.Println("Button from menu!") })
	menu3 := fyne.NewMenu("Info",
		info_item1,
		info_item2,
		info_item3,
		info_item4,
		info_item5,
		info_item6)

	main_menu := fyne.NewMainMenu(menu1, menu2, menu3)

	w.SetMainMenu(main_menu)

	box := container.NewWithoutLayout()
	w.SetContent(box)
	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
