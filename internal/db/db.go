package db

// go install github.com/pressly/goose/v3/cmd/goose@latest
// CLI Goose para migrations
//goose -dir ./migrations postgres "postgres://postgres:123@localhost:5432/scriptcreate?sslmode=disable" up

import (
	"database/sql"
	"fmt"
	"log"
	config "scriptmake/internal/Config"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Configuracao da conexao com o banco de dados
// go get github.com/lib/pq

// retorna um pool de conexao do banco
func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	port, err := strconv.Atoi(cfg.DBPort)
	if err != nil {
		port = 5432
	}

	// Data souce name para se conectar com o banco
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, port, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var dbConn *sql.DB
	for i := 0; i < 5; i++ {
		dbConn, err = sql.Open("postgres", dsn)
		if err == nil {
			if pingErr := dbConn.Ping(); pingErr == nil {
				log.Println("✅ Conectado ao banco:", cfg.DBName)
				return dbConn, nil
			}
		}
		log.Println("⏳ Tentando conectar ao banco, tentativa:", i+1)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("não foi possível conectar ao banco: %w", err)
}
