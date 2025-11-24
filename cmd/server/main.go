package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func tester(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет\n")
	response := map[string]string{"message": "testetetetet"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/tester", tester)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
