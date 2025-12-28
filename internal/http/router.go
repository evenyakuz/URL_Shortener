package http

import (
	"github.com/claygod/Bxog"
	"io"
	"net/http"
)

// Handlers
func IHandler(w http.ResponseWriter, req *http.Request, r *bxog.Router) {
	io.WriteString(w, "Welcome to Bxog!")
}
func THandler(w http.ResponseWriter, req *http.Request, r *bxog.Router) {
	params := r.Params(req, "/abc/:par")
	io.WriteString(w, "Params:\n")
	io.WriteString(w, " 'par' -> "+params["par"]+"\n")
}
func PHandler(w http.ResponseWriter, req *http.Request, r *bxog.Router) {
	// Получить параметры из урла
	params := r.Params(req, "country")
	io.WriteString(w, "Country:\n")
	io.WriteString(w, " 'name' -> "+params["name"]+"\n")
	io.WriteString(w, " 'capital' -> "+params["city"]+"\n")
	io.WriteString(w, " 'valuta' -> "+params["money"]+"\n")
	// Сгенерировать строку урла
	io.WriteString(w, "Creating a URL from route:\n")
	io.WriteString(w, r.Create("country", map[string]string{"name": "Russia", "capital": "Moscow", "money": "rouble"}))
}

// Main
func main() {
	m := bxog.New()
	m.Add("/", IHandler)
	m.Add("/abc/:par", THandler)
	m.Add("/country/:name/capital/:city/valuta/:money", PHandler).
		Id("country"). // Для удобства короткий ID
		Method("GET")  // Тут GET можно было не указывать, это для примера
	m.Start(":80")
}
