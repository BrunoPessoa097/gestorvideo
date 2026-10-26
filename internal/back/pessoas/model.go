package pessoas

import (
	"errors"
	"uuid"

	"golang.org/x/crypto/bcrypt"
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

	p := &Pessoas{
		ID:    uuid.New().String(),
		Nome:  nome,
		Senha: senha,
	}

	if err := p.toHash(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pessoas) toHash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	p.Senha = string(hashedPassword)

	return nil
}
