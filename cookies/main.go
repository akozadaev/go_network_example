package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем куку
		value := map[string]interface{}{"key": "value"}
		encoded, err := cookieHandler.Encode("cookie-name", value)
		if err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}

		// Чтение куки
		cookie, err := r.Cookie("cookie-name")
		if err == nil {
			value := make(map[string]interface{})
			if err := cookieHandler.Decode("cookie-name", cookie.Value, &value); err == nil {
				fmt.Fprintf(w, "Cookie Value: %s", value["key"])
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
