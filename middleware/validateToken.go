package middleware

import (
	"net/http"

	"github.com/wsaefulloh/go-solid-principle/helpers"
)

func Validate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("token_auth")

		validation, value := helpers.CheckToken(tokenString)

		if validation != true {
			helpers.Respone(w, "Login dulu", 500, true)
			return
		}

		if value != "admin" {
			helpers.Respone(w, "Maaf anda bukan admin", 400, true)
			return
		}

		next.ServeHTTP(w, r)
	}
}
