package pessoas

import (
	"errors"
	"uuid"

	"golang.org/x/crypto/bcrypt"
)

type Pessoas struct {
	ID       string `gorm:"primaryKey"`
	Nome     string `gorm:"not null"`
	Senha    string `gorm:"not null"`
	Telefone string `gorm:"null"`
}

func NewPessoa(nome, senha, telefone string) (*Pessoas, error) {
	if nome == "" || senha == "" {
		return nil, errors.New("Existem campos vazios")
	}

	p := &Pessoas{
		ID:       uuid.New().String(),
		Nome:     nome,
		Senha:    senha,
		Telefone: telefone,
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
