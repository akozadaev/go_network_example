package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		// Получение значения из сессии
		value := session.Values["key"]
		if value == nil {
			// Установка значения в сессии
			session.Values["key"] = "value"
			session.Save(r, w)
		}

		fmt.Fprintf(w, "Session Value: %s", value)
	})

	http.ListenAndServe(":8080", nil)
}
