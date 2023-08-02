package main

/*
This training program focuses on implementing resizing and positioning
elements using the Go language of the Fyne library. You will learn how
to create a user-friendly registration form, where elements such as
text fields, buttons, labels, and checkboxes can be adjusted in size
and positioned precisely within the form's layout. This program will
provide hands-on exercises and guidance on leveraging the Fyne library's
capabilities to create a dynamic and visually appealing user interface
for your registration form.
*/

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(500, 370))
	a.Settings().SetTheme(theme.DarkTheme())

	label := canvas.NewText("Регистрация", color.White)
	label.TextSize = 20
	label.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	label.Resize(fyne.NewSize(100, 30))
	label.Move(fyne.NewPos(190, 5))

	username := widget.NewEntry()
	username.Resize(fyne.NewSize(450, 40))
	username.Move(fyne.NewPos(25, 40))
	username.SetPlaceHolder("Ваше имя")

	pass := widget.NewPasswordEntry()
	pass.Resize(fyne.NewSize(450, 40))
	pass.Move(fyne.NewPos(25, 85))
	pass.SetPlaceHolder("Придумайте пароль")

	email := widget.NewEntry()
	email.Resize(fyne.NewSize(450, 40))
	email.Move(fyne.NewPos(25, 130))
	email.SetPlaceHolder("Ваша почта")

	info := widget.NewMultiLineEntry()
	info.Resize(fyne.NewSize(450, 100))
	info.Move(fyne.NewPos(25, 190))
	info.SetPlaceHolder("Дополнительная информация")

	submit := widget.NewButton("Зарегестрироваться", func() {
		file, err := os.Create("user.txt")
		if err != nil {
			fmt.Sprintf("Error create file: %v \n", err)
			os.Exit(1)
		}
		defer file.Close()

		data := fmt.Sprintf(
			"Имя пользователя: %s\nПароль: %s\nEmail: %s\nДополнительная информация: %s\n",
			username.Text, pass.Text, email.Text, info.Text,
		)
		file.WriteString(data)
	})

	submit.Resize(fyne.NewSize(300, 50))
	submit.Move(fyne.NewPos(100, 300))

	box := container.NewWithoutLayout(
		label,
		username,
		pass,
		email,
		info,
		submit,
	)
	w.SetContent(box)
	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
