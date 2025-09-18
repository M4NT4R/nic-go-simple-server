package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	// эндпоинт /ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "pong"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// эндпоинт /sum?a=2&b=3
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		// достаём query параметры
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")

		// конвертация строки в число
		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		// проверка на ошибки
		if err1 != nil || err2 != nil {
			http.Error(w, "invalid query params", http.StatusBadRequest)
			return
		}

		result := map[string]int{"result": a + b}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	// запуск сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}
