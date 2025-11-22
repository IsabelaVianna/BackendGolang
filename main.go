package main

import (
	"Api-Aula1/router"
	"log"
	"net/http"

	"github.com/jmo/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := router.New()
	const addr = ":10000"
	log.Printf("Servidor ouvindo em %s...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
