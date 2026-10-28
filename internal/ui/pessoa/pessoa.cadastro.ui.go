package pessoa

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/BrunoPessoa097/gestorvideo/internal/ui/components"
)

type CamposInput struct {
	nome     *widget.Entry
	senha    *widget.Entry
	telefone *widget.Entry
}

func inputsVall(c CamposInput) {
	nomeValor := c.nome.Text
	fmt.Print(nomeValor)
}

func Cadastro(w fyne.Window) {
	titulo := components.Title("Cadastro de Pessoas")

	nome := components.InputNormal("Digite seu nome")
	senha := components.InputPass("Digite sua senha")
	telefone := components.InputNormal("Digite seu telefone")

	labels := map[string]*widget.Entry{
		"Nome":     nome,
		"Senha":    senha,
		"Telefone": telefone,
	}

	formulario := components.Formulario(labels)

	botaoCadastar := components.Botao("Cadastro")
	botaoCadastar2 := components.Botao("Reset")

	botao := container.NewHBox(botaoCadastar, botaoCadastar2)

	conteudo := container.NewVBox(
		container.NewCenter(titulo),
		formulario,
		container.NewCenter(botao),
	)

	w.SetContent(container.NewPadded(conteudo))
}
