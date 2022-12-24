package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	myContainer := widget.NewLabel("")
	content := ""
	myContainer.SetText(content)

	// creating calculator buttons
	one := widget.NewButton("1", func() {
		content += "1"
	})
	two := widget.NewButton("2", func() {
		content += "2"
	})
	three := widget.NewButton("3", func() {
		content += "3"
	})
	four := widget.NewButton("4", func() {
		content += "4"
	})
	five := widget.NewButton("5", func() {
		content += "5"
	})
	six := widget.NewButton("6", func() {
		content += "6"
	})
	seven := widget.NewButton("7", func() {
		content += "7"
	})
	eight := widget.NewButton("8", func() {
		content += "8"
	})
	nine := widget.NewButton("9", func() {
		content += "9"
	})
	zero := widget.NewButton("0", func() {
		content += "0"
	})
	clear := widget.NewButton("C", func() {
		content += ""
	})
	plus := widget.NewButton("+", func() {
		content += "+"
	})
	minus := widget.NewButton("-", func() {
		content += "-"
	})
	multiply := widget.NewButton("×", func() {
		content += "×"
	})
	divide := widget.NewButton("÷", func() {
		content += "÷"
	})
	equal := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(content)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				content = strconv.FormatFloat(result.(float64), 'f', -1, 64)
			} else {
				content = err.Error()
			}
		} else {
			content = err.Error()
		}
	})

	// updating it each time in a detached mode
	go func() {
		for {
			myContainer.SetText(content)
		}
	}()

	// creating layout from the given buttons
	myWindow.SetContent(container.NewVBox(
		myContainer,
		container.NewGridWithColumns(3,
			divide,
			multiply,
			minus,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithRows(3,
				container.NewGridWithColumns(3,
					seven,
					eight,
					nine,
				),
				container.NewGridWithColumns(3,
					four,
					five,
					six,
				),
				container.NewGridWithColumns(3,
					one,
					two,
					three,
				),
			),
			plus,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				clear,
				zero,
			),
			equal,
		),
	),
	)

	// show window and run the go app
	myWindow.ShowAndRun()
}
