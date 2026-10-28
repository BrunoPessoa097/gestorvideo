package components

import "fyne.io/fyne/v2/widget"

func Botao(name string) *widget.Button {
	botao := widget.NewButton(name, func() {
		// funcao+()
	})
	return botao
}
