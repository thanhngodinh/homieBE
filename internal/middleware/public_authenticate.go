package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"hostel-service/internal/user/domain"
)

func PublicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		splitedTokenStr := strings.Split(tokenStr, " ")
		if len(splitedTokenStr) != 2 || splitedTokenStr[0] != "Bearer" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userId", "")))
			return
		}
		token, err := jwt.ParseWithClaims(
			splitedTokenStr[1],
			&jwt.StandardClaims{},
			func(t *jwt.Token) (interface{}, error) {
				return domain.USER_SECRET_KEY, nil
			},
		)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userId", "")))
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)
		// claims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userId", claims.Audience)))
	})
}
