package auth

import (
	"scriptmake/internal/apperror"
	"time"

	"github.com/golang-jwt/jwt"
)

// Service encapsula a lógica de geração e validação de JWT.
type Service struct {
	secret []byte
}

// NewService cria um novo serviço de JWT com a secret informada.
func NewService(secret string) *Service {
	// Retornando o struct com a secret convertida de string para byte
	return &Service{secret: []byte(secret)}
}

func (s *Service) GenerateToken(userID int, role string) (string, error) {
	// Definindo os dados/estrutura do token
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}

	// Cria o token com algoritimo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token e retorna a string
	return token.SignedString(s.secret)
}

func (s *Service) ParseToken(tokenStr string) (jwt.MapClaims, error) {

	// jwt.Parse() decodifica o token
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Verificando o algoritimo
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			// Se nao for o mesmo algoritimo usado retorna erro
			return nil, jwt.ErrSignatureInvalid
		}
		// Retorna a secret para validar a assinatura
		return s.secret, nil
	})

	// Verifica se o token é valido
	if !token.Valid {
		return nil, err
	}

	// Exraindo os claims(dados)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apperror.ErrInvalidToken
	}

	return claims, nil
}
