// handler/books_handler.go
package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição HandleSearch recebida")
}

func HandleSearchByAuthor(w http.ResponseWriter, r *http.Request) {
	authorQuery := r.URL.Query().Get("q")

	if authorQuery == "" {
		http.Error(w, "Faltou dizer quem é o autor (parâmetro 'q')", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	baseURL := "https://www.googleapis.com/books/v1/volumes"
	fullURL := fmt.Sprintf("%s?q=inauthor:%s&key=%s", baseURL, url.QueryEscape(authorQuery), apiKey)

	resp, err := http.Get(fullURL)
	if err != nil {
		http.Error(w, "Erro de comunicação com a API externa", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
