package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ItKarma/idocks/utils"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// verificando se o headers authorization existe
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// verificado se o formato do token é valido
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		//vericiando se o jwt é valido
		tokenStr := parts[1]
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// extraindo o id do token
		userID, ok := claims["userID"].(string)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Adicionar o userID ao contexto da requisição
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		// Passar para o próximo handler
		next.ServeHTTP(w, r)

	})
}
