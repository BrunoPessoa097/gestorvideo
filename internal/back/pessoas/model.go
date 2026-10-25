package pessoas

import "uuid"

type Pessoas struct {
	ID    string
	Nome  string
	Senha string
}

func NewPessoa(nome, senha string) *Pessoas {
	return &Pessoas{
		ID:    uuid.New().String(),
		Nome:  nome,
		Senha: senha,
	}
}
