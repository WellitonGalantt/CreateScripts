package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Arquivo para fazer configuracoes, bascar e criar carregar envs.

// Struct é como se fosse um objeto
// Estamos criando um objeto com as coniguracoes do projeto.
type Config struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	JWTSecret  string
	ServerPort string
}

func Load() *Config {
	// Ler o arquivo .env, carregar as variaveis
	_ = godotenv.Load() // Usando o _ para ignorar o erro

	cfg := &Config{
		DBName:     getEnv("DB_NAME", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "0"),
		ServerPort: getEnv("SERVER_PORT", "8000"),
	}

	return cfg
}

// Funcao para pegar um valor da variavel de ambiente com a chave
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("⚠️  %s não encontrado, usando valor padrão: %s\n", key, defaultValue)
		return defaultValue
	}

	return value
}
