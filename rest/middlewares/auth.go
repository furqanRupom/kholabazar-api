package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized! invalid token", http.StatusUnauthorized)
			return
		}


		jwtHeader := tokenParts[0]
		jwtClaim := tokenParts[1]
		jwtSign := tokenParts[2]
		message := jwtHeader + "." + jwtClaim
		byteArrSecret := []byte(m.conf.JWTSecret)
		byteArrMessage := []byte(message)
		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		newSign := Base64UrlEncode(hash)
		if newSign != jwtSign {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
