package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewBookList(books []string) fyne.CanvasObject {
	bookList := widget.NewList(
		func() int { return len(books) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(books[i])
		})

	return container.NewVBox(bookList)
}
