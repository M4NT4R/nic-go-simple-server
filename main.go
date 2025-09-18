package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "pong"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")

		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {
			http.Error(w, "invalid query params", http.StatusBadRequest)
			return
		}

		result := map[string]int{"result": a + b}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}

