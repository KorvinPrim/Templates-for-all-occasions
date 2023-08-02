package main

/*
The code is an example of a simple application using the Fyne library in Go.
It creates a window with a card widget that contains an accordion.
Each accordion item represents a paragraph with sub-items, and each
sub-item contains a label with text.

The application initializes the Fyne app, creates a window, sets the
window's size and theme, and creates the card with the accordion.
The card is then set as the content of the window, and the window
is displayed and the application is run.

Fyne is a cross-platform toolkit for creating graphical interfaces in Go.
It provides UI components that are easy to customize and scale for
different platforms.
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main")
	w.Resize(fyne.NewSize(500, 370))
	a.Settings().SetTheme(theme.DarkTheme())

	//text := "iriure adipisci persius facilis splendide eius civibus sagittis sententiae prompta ornare cubilia ornatus fabulas animal gubergren platea bibendum leo hac elementum dolor consectetuer vivamus leo quidam torquent donec postea inciderint amet no gravida fabulas molestiae ligula nostrum bibendum aperiri consectetur dissentiunt efficiantur dictas posidonium phasellus porro hendrerit maecenas delectus alienum"
	//text_label := widget.NewLabel(text)
	//text_label.Wrapping = fyne.TextWrapBreak
	//item1 := widget.NewAccordionItem(
	//	"Read more info",
	//	text_label,
	//)
	//
	//item2 := widget.NewAccordionItem(
	//	"Button",
	//	widget.NewButton("Say hello!", func() {
	//		println("say hello")
	//	}),
	//)
	//
	//item3 := widget.NewAccordionItem(
	//	"Chapters",
	//	widget.NewAccordion(
	//		widget.NewAccordionItem(
	//			"Chapter1",
	//			widget.NewLabel("Some text here..."),
	//		),
	//		widget.NewAccordionItem(
	//			"Chapter2",
	//			widget.NewLabel("Some text here..."),
	//		),
	//		widget.NewAccordionItem(
	//			"Chapter3",
	//			widget.NewLabel("Some text here..."),
	//		),
	//	))

	card := widget.NewCard(
		"Korvin GROUP",
		"Ознакомтесь с пользовательским соглашением",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Параграф 1",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Пункт 1",
						widget.NewLabel("Текст пункта 1")),
					widget.NewAccordionItem(
						"Пункт 2",
						widget.NewLabel("Текст пункта 2")),
					widget.NewAccordionItem(
						"Пункт 3",
						widget.NewLabel("Текст пункта 3")),
				),
			),
			widget.NewAccordionItem(
				"Параграф 2",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Пункт 1",
						widget.NewLabel("Текст пункта 1")),
					widget.NewAccordionItem(
						"Пункт 2",
						widget.NewLabel("Текст пункта 2")),
					widget.NewAccordionItem(
						"Пункт 3",
						widget.NewLabel("Текст пункта 3")),
				),
			),
			widget.NewAccordionItem(
				"Параграф 3",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Пункт 1",
						widget.NewLabel("Текст пункта 1")),
					widget.NewAccordionItem(
						"Пункт 2",
						widget.NewLabel("Текст пункта 2")),
					widget.NewAccordionItem(
						"Пункт 3",
						widget.NewLabel("Текст пункта 3")),
				),
			),
		),
	)

	//box := container.NewWithoutLayout(accordion)
	w.SetContent(card)
	w.ShowAndRun()
}

//fyne package -os windows -icon myapp.png
