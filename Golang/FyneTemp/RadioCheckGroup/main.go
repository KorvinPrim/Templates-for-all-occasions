package main

/*
An application program implemented on Fine. Demonstrates the logic of working with check and radio groups.
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Заказ еды")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Оформление заказа")

	name_label := widget.NewLabel("Ваше имя:")

	name := widget.NewEntry()

	food_label := widget.NewLabel("Выберете еду для заказа")
	food := widget.NewCheckGroup([]string{"Лапша", "Борщ", "Пельмени", "Суши", "Роллы"}, func(strings []string) {

	})

	male_label := widget.NewLabel("Ваш пол:")
	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(s string) {

	})

	result := widget.NewLabel("")

	btn := widget.NewButton("Оформить заказ", func() {
		username := name.Text
		food_arr := food.Selected
		male_val := male.Selected

		result.SetText(username + "\n" + male_val + "\n")

		for _, i := range food_arr {
			result.SetText(result.Text + "\n" + i)
		}

	})

	w.SetContent(container.NewVBox(
		title,

		name_label,
		name,

		food_label,
		food,

		male_label,
		male,

		btn,
		result,
	))

	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
