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
	fmt.Println("Bateu aqui")
}

func HandleSearchByAuthor(w http.ResponseWriter, r *http.Request) {
	authorQuery := r.URL.Query().Get("q")

	if authorQuery == "" {
		http.Error(w, "Parâmetro 'q' (autor) é obrigatório.", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	baseURL := "https://www.googleapis.com/books/v1/volumes"
	fullURL := fmt.Sprintf("%s?q=inauthor:%s&key=%s", baseURL, url.QueryEscape(authorQuery), apiKey)

	resp, err := http.Get(fullURL)
	if err != nil {
		http.Error(w, "Erro ao comunicar com a API externa.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
