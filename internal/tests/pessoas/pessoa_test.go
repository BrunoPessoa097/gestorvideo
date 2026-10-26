package pessoas_test

import (
	"testing"

	"github.com/BrunoPessoa097/gestorvideo/internal/back/db"
	"github.com/BrunoPessoa097/gestorvideo/internal/back/pessoas"
)

func TestCriarUsuario(t *testing.T) {
	_, err := pessoas.NewPessoa("Bruno", "senha", "88888888")

	if err != nil {
		t.Fatal("Pessoa não criada")
	}
}

func TestCriarUsuarioVazio(t *testing.T) {
	_, err := pessoas.NewPessoa("", "", "")
	mensagem := "Existem campos vazios"

	if err == nil {
		t.Fatal("Erro esperava vazio")
	}

	if err.Error() != mensagem {
		t.Errorf("mensagem de erro incorreta.\nEsperado: %s\nRecebido: %s", mensagem, err.Error())
	}
}

func TestSenhaHash(t *testing.T) {
	pessoa, _ := pessoas.NewPessoa("Bruno", "senha", "8888888")

	if pessoa.Senha == "senha" {
		t.Fatal("Problemas do hash")
	}
}

func TestSalvar(t *testing.T) {
	pessoa, _ := pessoas.NewPessoa("Bruno", "senha", "8888888")
	db := db.ConectarBanco()
	repo := pessoas.NewPessoaRepository(db.DB)
	saida := repo.Create(pessoa)

	if saida != nil {
		t.Fatal("Problemas para salvar")
	}
}
