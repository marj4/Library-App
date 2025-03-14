package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewBookForm(onSubmit func(title string)) fyne.CanvasObject {
	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Введите название книги")

	submitButton := widget.NewButton("Добавить", func() {
		onSubmit(titleEntry.Text)
	})

	return container.NewVBox(titleEntry, submitButton)
}
