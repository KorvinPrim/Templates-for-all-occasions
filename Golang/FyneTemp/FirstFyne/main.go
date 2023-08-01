package main

/*
The first application program implemented on Fine. A calculator that allows you to perform basic arithmetic calculations.
*/

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Почти калькулятор")

	label1 := widget.NewLabel("Введите первое число")
	first_num := widget.NewEntry()

	label2 := widget.NewLabel("Введите второе число")
	second_num := widget.NewEntry()

	label3 := widget.NewLabel("Результат:")
	label_result := widget.NewLabel("")

	btn_plus := widget.NewButton("+", func() {
		n1, err1 := strconv.ParseFloat(first_num.Text, 64)
		n2, err2 := strconv.ParseFloat(second_num.Text, 64)

		if err1 != nil || err2 != nil {
			label_result.SetText("Ошибка ввода")
		} else {
			sum := n1 + n2
			label_result.SetText(fmt.Sprintf("%v", sum))
		}

	})

	btn_minus := widget.NewButton("-", func() {
		n1, err1 := strconv.ParseFloat(first_num.Text, 64)
		n2, err2 := strconv.ParseFloat(second_num.Text, 64)

		if err1 != nil || err2 != nil {
			label_result.SetText("Ошибка ввода")
		} else {
			sum := n1 - n2
			label_result.SetText(fmt.Sprintf("%v", sum))
		}

	})

	btn_multiplication := widget.NewButton("*", func() {
		n1, err1 := strconv.ParseFloat(first_num.Text, 64)
		n2, err2 := strconv.ParseFloat(second_num.Text, 64)

		if err1 != nil || err2 != nil {
			label_result.SetText("Ошибка ввода")
		} else {
			sum := n1 * n2
			label_result.SetText(fmt.Sprintf("%v", sum))
		}

	})

	btn_division := widget.NewButton("/", func() {
		n1, err1 := strconv.ParseFloat(first_num.Text, 64)
		n2, err2 := strconv.ParseFloat(second_num.Text, 64)

		if err1 != nil || err2 != nil {
			label_result.SetText("Ошибка ввода")
		} else {
			sum := n1 / n2
			label_result.SetText(fmt.Sprintf("%v", sum))
		}

	})

	w.SetContent(container.NewVBox(
		label1,
		first_num,
		label2,
		second_num,

		btn_plus,
		btn_minus,
		btn_multiplication,
		btn_division,

		label3,
		label_result,
	))

	w.ShowAndRun()
}
