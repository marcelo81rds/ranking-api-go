package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"ranking-api/db"
	"ranking-api/models"
)

func GetRanking(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query(context.Background(),
		"SELECT id, nome, pontuacao FROM ranking ORDER BY pontuacao DESC")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var rankings []models.Ranking

	for rows.Next() {
		var r models.Ranking
		rows.Scan(&r.ID, &r.Nome, &r.Pontuacao)
		rankings = append(rankings, r)
	}

	json.NewEncoder(w).Encode(rankings)
}

func CreateRanking(w http.ResponseWriter, r *http.Request) {
	var ranking models.Ranking

	json.NewDecoder(r.Body).Decode(&ranking)

	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO ranking (nome, pontuacao) VALUES ($1, $2)",
		ranking.Nome, ranking.Pontuacao)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
