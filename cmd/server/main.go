package main

import (
	"fmt"
	"log"
	config "scriptmake/internal/Config"
	"scriptmake/internal/auth"
	"scriptmake/internal/db"
	"scriptmake/internal/middleware"
	"scriptmake/internal/modules/user"

	"github.com/gin-gonic/gin"
)

// Aquivo principal o coracao do servidor

func main() {

	// Carregando as configuracoes da variaveis de ambiente
	cfg := config.Load()

	// Mando as variaveis carregada para criar a conexao com o banco
	dbConn, err := db.NewPostgresDB(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Falha ao conectar no banco")
	}

	defer dbConn.Close()

	//Parte das instancia, iniciar os modulos

	//JWT
	jwtService := auth.NewService(config.Load().JWTSecret)

	// User
	userRepository := user.NewPostgresRepository(dbConn)
	userUseCase := user.NewService(userRepository, jwtService)
	userController := user.NewUserHandler(userUseCase)

	// Parte de rotas, configurar as rotas

	//Gin
	router := gin.Default()

	router.POST("/user/register", userController.Register)
	router.POST("/user/login", userController.Login)

	// Grupo para rotas protegidas
	authGroup := router.Group("/api")
	authGroup.Use(middleware.AuthMiddleware(jwtService))

	// Adicionando as rotas protegidas
	authGroup.GET("/user/profile", userController.ViewProfile)

	log.Println("🚀 Servidor rodando na porta", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("erro ao subir servidor:", err)
	}
}

//go run .\cmd\server\main.go

//go build -o app
//./app

//go install github.com/air-verse/air@latest
//air
