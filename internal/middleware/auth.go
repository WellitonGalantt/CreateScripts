package middleware

import (
	"net/http"
	"scriptmake/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authotization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token nao informado!"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		tokenStr := parts[0]

		claims, err := jwtService.ParseToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", sub)
		ctx.Set("role", role)

		ctx.Next()

	}
}
