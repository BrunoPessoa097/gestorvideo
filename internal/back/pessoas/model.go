package pessoas

import (
	"errors"
	"uuid"
)

type Pessoas struct {
	ID    string
	Nome  string
	Senha string
}

func NewPessoa(nome, senha string) (*Pessoas, error) {
	if nome == "" || senha == "" {
		return nil, errors.New("Existem campos vazios")
	}

	return &Pessoas{
		ID:    uuid.New().String(),
		Nome:  nome,
		Senha: senha,
	}, nil
}
