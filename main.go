package main

import (
	"log"
	"net/http"

	"ranking-api/db"
	"ranking-api/handlers"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ranking", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetRanking(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateRanking(w, r)
		}
	})

	log.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
