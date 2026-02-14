package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Удаляем "Bearer " из начала токена
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Парсим токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи и возвращаем секретный ключ
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("my-secret-key"), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Проверяем валидность токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Проверяем срок действия токена
			expTime := time.Unix(int64(claims["exp"].(float64)), 0)
			if expTime.Before(time.Now()) {
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			}

			// Токен валидный, продолжаем обработку
			username := claims["username"].(string)
			fmt.Fprintf(w, "Welcome, %s!", username)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
