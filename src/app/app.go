package app

import (
	"log"
	"net/http"
	"os"

	"github.com/c1r5/gormigrate"
	_ "github.com/c1r5/url-shortener/src/internal/database"
	"github.com/c1r5/url-shortener/src/internal/shortener"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() error {
	db_url := os.Getenv("DATABASE_URL")
	pos := postgres.Open(db_url)
	db, err := gorm.Open(pos, &gorm.Config{})

	if err != nil {
		return err
	}

	if err := gormigrate.Run(db, 1); err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	mux := http.NewServeMux()

	shortener.New(db).RegisterRoutes(mux)

	log.Printf("Servidor rodando na porta 3001")
	http.ListenAndServe(":3001", mux)

	return nil
}
