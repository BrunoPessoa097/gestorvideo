package components

import "fyne.io/fyne/v2/widget"

func Formulario(campos map[string]*widget.Entry) *widget.Form {
	entradas := []*widget.FormItem{}

	for texto, input := range campos {
		entradas = append(
			entradas,
			widget.NewFormItem(texto, input),
		)
	}

	formulario := widget.NewForm(
		entradas...,
	)

	return formulario
}
