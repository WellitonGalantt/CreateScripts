package main

import (
	"log"
	config "scriptmake/internal/Config"
	"scriptmake/internal/db"

	"github.com/gin-gonic/gin"
)

// Aquivo principal o coracao do servidor

func main() {

	// Carregando as configuracoes da variaveis de ambiente
	cfg := config.Load()

	// Mando as variaveis carregada para criar a conexao com o banco
	dbConn, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Falha ao conectar no banco")
	}

	defer dbConn.Close()

	//Parte das instancia, iniciar os modulos

	// Parte de rotas, configurar as rotas

	router := gin.Default()

	log.Println("🚀 Servidor rodando na porta", cfg.DBPort)
	if err := router.Run(":" + cfg.DBPort); err != nil {
		log.Fatal("erro ao subir servidor:", err)
	}
}
