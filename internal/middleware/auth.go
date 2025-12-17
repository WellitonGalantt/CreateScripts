package middleware

import (
	"fmt"
	"net/http"
	"scriptmake/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("Aqui 5")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token nao informado!"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			fmt.Println("Aqui 4")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		tokenStr := parts[1]

		claims, err := jwtService.ParseToken(tokenStr)
		if err != nil {
			fmt.Println("Aqui 3")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			fmt.Println("Aqui 1")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			fmt.Println("Aqui 2")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", sub)
		ctx.Set("role", role)

		ctx.Next()

	}
}
