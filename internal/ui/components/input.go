package components

import "fyne.io/fyne/v2/widget"

func InputNormal(placeholder string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	return input
}
