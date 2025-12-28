package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main page")
	response := map[string]string{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/main", mpage)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
