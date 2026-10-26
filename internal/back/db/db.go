package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BrunoPessoa097/gestorvideo/internal/back/pessoas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func ConectarBanco() *Repository {
	env, err := envLoad()

	if err != nil {
		fmt.Print(err.Error())
	}

	db, err := gorm.Open(postgres.Open(env), &gorm.Config{})

	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	err = db.AutoMigrate(&pessoas.Pessoas{})

	if err != nil {
		log.Fatalf("Erro ao rodar migrações: %v", err)
	}

	return &Repository{DB: db}
}

func envLoad() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Arquivo .env não encontrado")
	}

	env := os.Getenv("DATABASE_URL")

	if env != "" {
		return "", errors.New("A variável DATABASE_URL não foi definida no .env")
	}

	return env, nil
}
