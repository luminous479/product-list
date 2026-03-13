package jwtauthentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/luminous479/product-list/env"
)

func JwtAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		accessToken := tokenParts[1]
		tokenParts = strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Invalid JWT token format", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		conf := env.GetConfig()
		byteSecret := []byte(conf.JwtSecretKey)
		bytemessage := []byte(message)

		h := hmac.New(sha256.New, byteSecret)
		h.Write(bytemessage)
		hash := h.Sum(nil)
		newSignature := base64Encode(hash)

		if newSignature != signature {
			http.Error(w, "Invalid JWT token signature", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
