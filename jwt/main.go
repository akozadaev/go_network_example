package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Создаем токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем утверждения (claims)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "user123"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Время истечения через 24 часа

	// Генерируем секретный ключ
	secretKey := []byte("my-secret-key")

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JWT:", tokenString)
}
