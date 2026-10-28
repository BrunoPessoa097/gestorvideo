package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/BrunoPessoa097/gestorvideo/internal/ui/pessoa"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gestor de Vídeo")
	w.Resize(fyne.NewSize(600, 100))

	pessoa.Cadastro(w)

	w.ShowAndRun()
}
