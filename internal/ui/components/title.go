package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Title(title string) *widget.Label {
	titulo := widget.NewLabelWithStyle(
		title,
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	return titulo
}
