package pessoas

import "gorm.io/gorm"

type PessoaRepository struct {
	DB *gorm.DB
}

func NewPessoaRepository(db *gorm.DB) *PessoaRepository {
	return &PessoaRepository{
		DB: db,
	}
}

func (pr *PessoaRepository) Create(p *Pessoas) error {
	return pr.DB.Create(p).Error
}
